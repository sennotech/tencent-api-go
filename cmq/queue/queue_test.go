package queue

import "testing"

func TestParseHttps1(t *testing.T) {
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

func TestParseHttps2(t *testing.T) {
    scheme, domain, region, err := parse("https://cmq-sh.public.tencenttdmq.com")
    if err != nil {
        t.Error("没有正确解析endpoint")
    }
    if scheme != "https" {
        t.Error("没有正确解析scheme")
    }
    if domain != "cmq-sh.public.tencenttdmq.com" {
        t.Error("没有正确解析domain")
    }
    if region != "sh" {
        t.Error("没有正确解析region")
    }
}

func TestParseHttp1(t *testing.T) {
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

func TestParseHttp2(t *testing.T) {
    scheme, domain, region, err := parse("http://sh.mqadapter.cmq.tencentyun.com")
    if err != nil {
        t.Error("没有正确解析endpoint")
    }
    if scheme != "http" {
        t.Error("没有正确解析scheme")
    }
    if domain != "sh.mqadapter.cmq.tencentyun.com" {
        t.Error("没有正确解析domain")
    }
    if region != "sh" {
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
