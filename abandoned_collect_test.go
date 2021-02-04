package kerror

import "testing"

func TestCollect(t *testing.T) {
	err1 := New(ECustom+1, "test error 1")
	err2 := New(ECustom+2, "test error 2")
	err3 := New(ECustom+3, "test error 3")
	err := Collect(err1, err2, err3)
	t.Logf("%+v", err)
	errs, ok := err.(MultiError)
	if !ok {
		t.Fail()
		return
	}
	if len(errs) != 3 || errs[0] != err1 || errs[1] != err2 || errs[2] != err3 {
		t.Fail()
		return
	}
}
