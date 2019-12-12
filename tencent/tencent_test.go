package tencent

import (
	"net/http"
	"testing"
)

func TestSignature(t *testing.T) {
	expected := "C16WEtEXsD5v5tnaUMLAbZewXhI%3D"
	p := &parameter{
		httpMethod:      http.MethodPost,
		domain:          "cmq-queue-gz.api.tencentyun.com",
		path:            "/v2/index.php",
		secretId:        "AKIDPcYDclDJCn8D0Xypa4f3pKYUCVYLn3zT",
		secretKey:       "pPgfLipfEXZ7VcRzhAMIyPaU7UbQyFFx",
		signatureMethod: HmacSHA1,
		nonce:           2889712707386595659,
		timestamp:       1534154812,
	}
	actual := p.signature(map[string]interface{}{
		"Action":          "SendMessage",
		"Nonce":           "2889712707386595659",
		"RequestClient":   "SDK_Python_1.3",
		"SecretId":        "AKIDPcYDclDJCn8D0Xypa4f3pKYUCVYLn3zT",
		"SignatureMethod": "HmacSHA1",
		"Timestamp":       "1534154812",
		"clientRequestId": "1231231231",
		"delaySeconds":    "0",
		"msgBody":         "msg",
		"queueName":       "test1",
	})

	if actual != expected {
		t.Errorf("签名失败，期望：%s，实际：%s", expected, actual)
	}
}
