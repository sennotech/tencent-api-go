package cmq

import "github.com/sennotech/tencent-api-go/tencent"

func IsNoMessage(err *tencent.Error) bool {
	return err.Code == 7000
}

func IsExisted(err *tencent.Error) bool {
	return err.ModuleCode() == 4460
}
