package queue

import (
    "github.com/sennotech/tencent-api-go/cmq"
    "github.com/sennotech/tencent-api-go/tencent"
)

type deleteMessage struct {
    *tencent.Response
}

func (q *Queue) DeleteMessage(receiptHandle string) (
    *deleteMessage, error) {
    msg := &deleteMessage{}
    err := tencent.Get(q.Scheme, q.Domain, cmq.Path, "DeleteMessage", q.Region,
        q.SecretId, q.SecretKey, map[string]string{
            "queueName":     q.Name,
            "receiptHandle": receiptHandle,
        }, msg)
    return msg, err
}
