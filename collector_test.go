package kerror

import "testing"

func TestEmptyCollector(t *testing.T) {
	coerr := NewCollector()
	if coerr.Error() != nil {
		t.Fail()
		return
	}
}

func TestCollectorWithSingleError(t *testing.T) {
	err1 := New(ECustom, "test error")
	coerr := NewCollector()
	coerr.Collect(err1)
	err := coerr.Error()
	t.Logf("\n%+v\n", err)
	if err != err1 {
		t.Fail()
		return
	}
}

func TestCollectorWithLists(t *testing.T) {
	err1 := New(ECustom+1, "test error 1")
	err2 := New(ECustom+2, "test error 2")
	err3 := New(ECustom+3, "test error 3")
	coerr := NewCollector()
	coerr.Collect(err1)
	coerr.Collect(MultiError{err2, nil, err3})
	err := coerr.Error()
	t.Logf("\n%+v\n", err)
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
