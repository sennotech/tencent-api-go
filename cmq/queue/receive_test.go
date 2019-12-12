package queue

import (
    "os"
    "testing"

    "github.com/sennotech/tencent-api-go/cmq"
    "github.com/sennotech/tencent-api-go/tencent"
)

func TestReceiveMessage(t *testing.T) {
    account, _ := tencent.NewAccount(os.Getenv("secretId"),
        os.Getenv("secretKey"))
    q, _ := New(account, os.Getenv("endpoint"), os.Getenv("queueName"))

    result, err := q.ReceiveMessage(5)
    if err != nil {
        if err, ok := err.(*tencent.Error); ok && cmq.IsNoMessage(err) {
            t.Logf("there's no message")
        }
    }
    t.Log(result)
}
