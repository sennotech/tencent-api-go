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

var _http = regexp.MustCompile("(http)://(cmq-topic-(\\w+).api.tencentyun.com)")
var _https = regexp.MustCompile("(https)://(cmq-topic-(\\w+).api.qcloud.com)")

func parse(endpoint string) (scheme, domain, region string, err error) {
    var reg *regexp.Regexp
    switch {
    case _http.MatchString(endpoint):
        reg = _http
    case _https.MatchString(endpoint):
        reg = _https
    default:
        err = &InvalidEndpointError{endpoint}
        return
    }

    pieces := reg.FindStringSubmatch(endpoint)

    return pieces[1], pieces[2], pieces[3], nil
}
