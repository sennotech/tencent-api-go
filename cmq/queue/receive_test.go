package queue

import (
    "os"
    "testing"

    "github.com/sennotech/tencent-api-go/cmq"
    "github.com/sennotech/tencent-api-go/tencent"
)

func TestReceiveMessage(t *testing.T) {
    account, _ := tencent.NewAccount(os.Getenv("CMQ_SECRET_ID"),
        os.Getenv("CMQ_SECRET_KEY"))
    q, _ := New(account, os.Getenv("CMQ_QUEUE_ENDPOINT"), os.Getenv("CMQ_QUEUE_NAME"))

    result, err := q.ReceiveMessage(5)
    if err != nil {
        if err, ok := err.(*tencent.Error); ok && cmq.IsNoMessage(err) {
            t.Logf("there's no message")
        }
    }
    t.Log(result)
}
