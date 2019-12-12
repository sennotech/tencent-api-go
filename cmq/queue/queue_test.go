package queue

import "testing"

func TestParseHttpsAndQCloud(t *testing.T) {
    scheme, domain, region, err := parse("https://cmq-queue-sh.api.qcloud.com")
    if err != nil {
        t.Error("没有正确解析endpoint")
    }
    if scheme != "https" {
        t.Error("没有正确解析scheme")
    }
    if domain != "cmq-queue-sh.api.qcloud.com" {
        t.Error("没有正确解析domain")
    }
    if region != "sh" {
        t.Error("没有正确解析region")
    }
}

func TestParseHttpAndTencentyun(t *testing.T) {
    scheme, domain, region, err := parse("http://cmq-queue-gz.api.tencentyun.com")
    if err != nil {
        t.Error("没有正确解析endpoint")
    }
    if scheme != "http" {
        t.Error("没有正确解析scheme")
    }
    if domain != "cmq-queue-gz.api.tencentyun.com" {
        t.Error("没有正确解析domain")
    }
    if region != "gz" {
        t.Error("没有正确解析region")
    }
}

func TestParseHttpAndQCloud(t *testing.T) {
    _, _, _, err := parse("http://cmq-queue-sh.api.qcloud.com")
    if err == nil {
        t.Error("没有正确解析endpoint")
    }
}

func TestParseEmptyEndpoint(t *testing.T) {
    _, _, _, err := parse("")
    if err == nil {
        t.Error("没有正确解析endpoint")
    }
}
