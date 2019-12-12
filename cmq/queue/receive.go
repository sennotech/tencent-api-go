package queue

import (
    "strconv"

    "github.com/sennotech/tencent-api-go/cmq"
    "github.com/sennotech/tencent-api-go/tencent"
)

type ReceiveMessage struct {
    Id              string `json:"msgId"`
    Body            string `json:"msgBody"`
    ReceiptHandle   string `json:"receiptHandle"`
    EnqueueTime     int    `json:"enqueueTime"`
    NextVisibleTime int    `json:"nextVisibleTime"`
}

func (q *Queue) ReceiveMessage(pollingWaitSeconds int) (
    *ReceiveMessage, error) {
    msg := &ReceiveMessage{}
    err := tencent.Get(q.Scheme, q.Domain, cmq.Path, "ReceiveMessage", q.Region,
        q.SecretId, q.SecretKey, map[string]string{
            "queueName":          q.Name,
            "pollingWaitSeconds": strconv.Itoa(pollingWaitSeconds),
        }, msg)
    return msg, err
}
