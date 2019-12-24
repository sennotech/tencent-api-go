package queue

import (
    "math/rand"
    "os"
    "testing"
    "time"

    "github.com/sennotech/tencent-api-go/tencent"
)

func TestCreate(t *testing.T) {
    account, _ := tencent.NewAccount(os.Getenv("CMQ_SECRET_ID"),
        os.Getenv("CMQ_SECRET_KEY"))
    q, _ := New(account, os.Getenv("CMQ_QUEUE_ENDPOINT"), randomString(10))

    result, err := q.Create(1000000, 30, 3600, 1, 7 * 24 * 3600, 0)
    if err != nil {
        t.Errorf("failed to create queue, %v", err)
        return
    }
    t.Log("queue id: ", result.QueueId)
}

var letterRunes = []rune(
    "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randomString(n int) string {
    rand.Seed(time.Now().UnixNano())

    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}
