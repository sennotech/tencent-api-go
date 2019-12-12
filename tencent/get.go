package tencent

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Get(scheme, domain, path, action, region, secretId, secretKey string,
	others map[string]string, result interface{}) error {
	p := newParameter(scheme, domain, path, action, region, secretId, secretKey,
		others)

	_url := p.url()
	resp, err := http.Get(_url)
	if err != nil {
		return &ConnectionError{err.Error()}
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return &HttpRequestError{resp.StatusCode}
	}

	body, err := ioutil.ReadAll(resp.Body)

	output := &Response{}
	if err := json.Unmarshal(body, output); err != nil {
		return err
	}
	if output.Code != 0 {
		return &Error{output}
	}

	return json.Unmarshal(body, result)
}
