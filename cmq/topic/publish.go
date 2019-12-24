package topic

import (
    "fmt"

    "github.com/sennotech/tencent-api-go/cmq"
    "github.com/sennotech/tencent-api-go/tencent"
)

const PublishMessage = "PublishMessage"

type PublishResponse struct {
    Code int `json:"code"`
    Message string `json:"message"`
    RequestId string `json:"requestId"`
    MessageId string `json:"msgId"`
}

func (t *Topic) PublishMessage(message string, messageTags []string,
    routingKey string) (*PublishResponse, error) {
    resp := &PublishResponse{}

    parameters := map[string]string{
        "topicName": t.Name,
        "msgBody": message,
    }
    for i, tag := range messageTags {
        parameters[fmt.Sprintf("msgTag.%d", i)] = tag
    }
    if routingKey != "" {
        parameters["routingKey"] = routingKey
    }

    err := tencent.Get(t.Scheme, t.Domain, cmq.Path, PublishMessage, t.Region,
        t.SecretId, t.SecretKey, parameters, resp)

    return resp, err
}
