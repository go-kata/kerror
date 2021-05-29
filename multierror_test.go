package kerror

import "testing"

func newTestMultiError(t *testing.T) (errors []error, errs MultiError) {
	errors = []error{
		New(Label("test.Error1"), "test error 1"),
		New(Label("test.Error2"), "test error 2"),
		New(Label("test.Error3"), "test error 3"),
		New(Label("test.Error4"), "test error 4"),
		New(Label("test.Error5"), "test error 5"),
		New(Label("test.Error6"), "test error 6"),
	}
	errs = MultiError{errors[0], errors[1], MultiError{errors[2], MultiError{nil, errors[3]}, errors[4]}, errors[5]}
	t.Logf("%+v", errs)
	return
}

func TestMultiError__Nil(t *testing.T) {
	errs := MultiError{New(Label("test.Error"), "test error"), nil}
	t.Logf("%+v", errs)
}

func TestNilMultiError_Error(t *testing.T) {
	if s := (MultiError)(nil).Error(); s != "" {
		t.Logf("%s", s)
		t.Fail()
		return
	}
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

func TestTraverse__Break(t *testing.T) {
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
