package kerror

import "testing"

func newTestMultiError(t *testing.T) (errors []error, errs MultiError) {
	errors = []error{
		New(ECustom+1, "test error 1"),
		New(ECustom+2, "test error 2"),
		New(ECustom+3, "test error 3"),
		New(ECustom+4, "test error 4"),
		New(ECustom+5, "test error 5"),
		New(ECustom+6, "test error 6"),
	}
	errs = MultiError{errors[0], errors[1], MultiError{errors[2], MultiError{nil, errors[3]}, errors[4]}, errors[5]}
	t.Logf("%+v", errs)
	return
}

func TestMultiErrorWithNil(t *testing.T) {
	errs := MultiError{New(ECustom, "test error"), nil}
	t.Logf("%+v", errs)
}

func TestTraverse(t *testing.T) {
	errors, errs := newTestMultiError(t)
	i := 0
	Traverse(errs, func(err error) (next bool) {
		t.Logf("%+v", err)
		if err != errors[i] {
			t.Fail()
			return false
		}
		i++
		return true
	})
	if i != len(errors) {
		t.Fail()
		return
	}
}

func TestTraverseWithBreak(t *testing.T) {
	errors, errs := newTestMultiError(t)
	i := 0
	Traverse(errs, func(err error) (next bool) {
		t.Logf("%+v", err)
		if err != errors[i] {
			t.Fail()
			return false
		}
		i++
		return i < 4
	})
	if i != 4 {
		t.Fail()
		return
	}
}

func TestNilMultiError_Error(t *testing.T) {
	if s := (MultiError)(nil).Error(); s != "" {
		t.Logf("%s", s)
		t.Fail()
		return
	}
}
