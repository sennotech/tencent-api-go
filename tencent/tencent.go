package tencent

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"hash"
	"math/rand"
	"net/http"
	"net/url"
	"path"
	"sort"
	"strconv"
	"time"
)

const (
	HmacSHA1   = "hmacSHA1"
	HmacSHA256 = "HmacSHA256"
)

type parameter struct {
	scheme          string
	httpMethod      string
	domain          string
	path            string
	action          string
	region          string
	others          map[string]string
	secretId        string
	secretKey       string
	signatureMethod string
	timestamp       string
	nonce           string
}

func newParameter(scheme, domain, path, action, region, secretId,
	secretKey string, others map[string]string) *parameter {
	now := time.Now().Unix()

	p := &parameter{
		scheme:          scheme,
		httpMethod:      http.MethodGet,
		domain:          domain,
		path:            path,
		action:          action,
		region:          region,
		others:          others,
		secretId:        secretId,
		secretKey:       secretKey,
		signatureMethod: HmacSHA256,
		timestamp:       strconv.FormatInt(now, 10),
	}

	rand.Seed(now)
	rand.Seed(rand.Int63())
	p.nonce = strconv.Itoa(rand.Intn(100000))

	return p
}

func (p *parameter) url() string {
	return fmt.Sprintf("%s://%s?%s", p.scheme, path.Join(p.domain, p.path), p.query())
}

func (p *parameter) query() string {
	params := map[string]string{
		"Action":          p.action,
		"Region":          p.region,
		"Timestamp":       p.timestamp,
		"Nonce":           p.nonce,
		"SecretId":        p.secretId,
		"SignatureMethod": p.signatureMethod,
	}
	for k, v := range p.others {
		params[k] = v
	}
	params["Signature"] = p.signature(params)

	return mapToParams(params)
}

func (p *parameter) signature(params map[string]string) string {
	keys := keys(params)
	sort.Strings(keys)

	origin := p.httpMethod + path.Join(p.domain, p.path) + "?"
	for i, key := range keys {
		if i != 0 {
			origin += "&"
		}
		origin += fmt.Sprintf("%v=%v", key, params[key])
	}

	h := hmac.New(p.hash, []byte(p.secretKey))
	h.Write([]byte(origin))

	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return url.QueryEscape(signature)
}

func (p *parameter) hash() hash.Hash {
	switch p.signatureMethod {
	case HmacSHA1:
		return sha1.New()
	case HmacSHA256:
		return sha256.New()
	default:
		return sha256.New()
	}
}

func keys(m map[string]string) []string {
	var keys []string

	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

func mapToParams(m map[string]string) string {
	params := url.Values{}

	for k, v := range m {
		params.Add(k, v)
	}

	return params.Encode()
}
