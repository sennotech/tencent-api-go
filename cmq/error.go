package cmq

import "github.com/sennotech/tencent-api-go/tencent"

func IsNoMessage(err *tencent.Error) bool {
	return err.Code == 7000
}
