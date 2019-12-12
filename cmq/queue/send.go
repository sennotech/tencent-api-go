package queue

import (
    "strconv"

    "github.com/sennotech/tencent-api-go/cmq"
    "github.com/sennotech/tencent-api-go/tencent"
)

type sendMessage struct {
    *tencent.Response
    Id string `json:"msgId"`
}

func (q *Queue) SendMessage(msg string, delay int) (
    *sendMessage, error) {
    _msg := &sendMessage{}
    err := tencent.Get(q.Scheme, q.Domain, cmq.Path, "SendMessage", q.Region,
        q.SecretId, q.SecretKey, map[string]string{
            "queueName":    q.Name,
            "msgBody":      msg,
            "delaySeconds": strconv.Itoa(delay),
        }, _msg)
    return _msg, err
}
