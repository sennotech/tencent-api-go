package queue

import (
    "strconv"

    "github.com/sennotech/tencent-api-go/cmq"
    "github.com/sennotech/tencent-api-go/tencent"
)

const createQueue = "CreateQueue"

type Message4Create struct {
    RequestId string `json:"requestId"`
    QueueId   string `json:"queueId"`
}

func (q *Queue) Create(maxMsgHeapNum, pollingWaitSeconds,
    visibilityTimeoutSeconds, maxMsgSizeKs, msgRetentionSeconds,
    rewindSeconds int) (*Message4Create, error) {
    msg := &Message4Create{}
    err := tencent.Get(q.Scheme, q.Domain, cmq.Path, createQueue, q.Region,
        q.SecretId, q.SecretKey, map[string]string{
            "queueName":           q.Name,
            "maxMsgHeapNum":       strconv.Itoa(maxMsgHeapNum),
            "pollingWaitSeconds":  strconv.Itoa(pollingWaitSeconds),
            "visibilityTimeout":   strconv.Itoa(visibilityTimeoutSeconds),
            "maxMsgSize":          strconv.Itoa(maxMsgSizeKs * 1024),
            "msgRetentionSeconds": strconv.Itoa(msgRetentionSeconds),
            "rewindSeconds":       strconv.Itoa(rewindSeconds),
        }, msg)
    return msg, err
}
