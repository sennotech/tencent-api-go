package topic

import (
    "os"
    "testing"

    "github.com/sennotech/tencent-api-go/tencent"
)

func TestPublishMessage(t *testing.T) {
    account, _ := tencent.NewAccount(os.Getenv("CMQ_SECRET_ID"),
        os.Getenv("CMQ_SECRET_KEY"))
    topic, _ := New(account, os.Getenv("CMQ_TOPIC_ENDPOINT"), os.Getenv("CMQ_TOPIC_NAME"))

    result, err := topic.PublishMessage("message", []string{"image"}, "")
    if err != nil {
        t.Error(err)
    }
    t.Log(result)
}
