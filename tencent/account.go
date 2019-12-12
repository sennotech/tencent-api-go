package tencent

import "fmt"

type Account struct {
    SecretId  string
    SecretKey string
}

func NewAccount(secretId, secretKey string) (*Account, error) {
    if secretId == "" || secretKey == "" {
        return nil, fmt.Errorf("secretId or secretKey cannot be empty")
    }
    return &Account{secretId, secretKey}, nil
}
