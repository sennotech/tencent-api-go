package tencent

import "testing"

func TestModuleCode(t *testing.T) {
    e := &Error{&Response{
        Code:      0,
        Message:   "（4460）queue is already existed,case insensitive",
        RequestId: "",
    }}
    if e.ModuleCode() != 4460 {
        t.Errorf("expected 4460 but is %d", e.ModuleCode())
    }
}
