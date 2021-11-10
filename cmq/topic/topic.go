package topic

import (
    "fmt"
    "regexp"

    "github.com/sennotech/tencent-api-go/tencent"
)

type Topic struct {
    *tencent.Account
    Endpoint string
    Scheme   string
    Domain   string
    Region   string
    Name     string
}

func New(account *tencent.Account, endpoint, name string) (*Topic, error) {
    if account == nil {
        return nil, fmt.Errorf("account should not be empty")
    }

    if name == "" {
        return nil, fmt.Errorf("queue name should not be empty")
    }

    scheme, domain, region, err := parse(endpoint)
    if err != nil {
        return nil, err
    }

    return &Topic{
        Account:  account,
        Endpoint: endpoint,
        Scheme:   scheme,
        Domain:   domain,
        Region:   region,
        Name:     name,
    }, nil
}

type InvalidEndpointError struct {
    endpoint string
}

func (e *InvalidEndpointError) Error() string {
    return fmt.Sprintf("invalid endpoint: %s", e.endpoint)
}

var _http1 = regexp.MustCompile("(http)://(cmq-topic-(\\w+).api.tencentyun.com)")
var _http2 = regexp.MustCompile("(http)://((\\w+).mqadapter.cmq.tencentyun.com)")
var _https1 = regexp.MustCompile("(https)://(cmq-topic-(\\w+).api.qcloud.com)")
var _https2 = regexp.MustCompile("(https)://(cmq-(\\w+).public.tencenttdmq.com)")

func parse(endpoint string) (scheme, domain, region string, err error) {
    var reg *regexp.Regexp
    switch {
    case _http1.MatchString(endpoint):
        reg = _http1
    case _http2.MatchString(endpoint):
        reg = _http2
    case _https1.MatchString(endpoint):
        reg = _https1
    case _https2.MatchString(endpoint):
        reg = _https2
    default:
        err = &InvalidEndpointError{endpoint}
        return
    }

    pieces := reg.FindStringSubmatch(endpoint)

    return pieces[1], pieces[2], pieces[3], nil
}
