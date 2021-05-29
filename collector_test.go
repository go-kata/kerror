package kerror

import "testing"

func TestCollector__Empty(t *testing.T) {
	coerr := NewCollector()
	if coerr.Error() != nil {
		t.Fail()
		return
	}
}

func TestCollector__SingleError(t *testing.T) {
	err1 := New(Label("test.Error"), "test error")
	coerr := NewCollector()
	coerr.Collect(err1)
	err := coerr.Error()
	t.Logf("%+v", err)
	if err != err1 {
		t.Fail()
		return
	}
}

func TestCollector__MultipleErrors(t *testing.T) {
	err1 := New(Label("test.Error1"), "test error 1")
	err2 := New(Label("test.Error2"), "test error 2")
	err3 := New(Label("test.Error3"), "test error 3")
	coerr := NewCollector()
	coerr.Collect(err1)
	coerr.Collect(MultiError{err2, nil, err3})
	err := coerr.Error()
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

func TestNilCollector_Collect(t *testing.T) {
	err := Try(func() error {
		(*Collector)(nil).Collect()
		return nil
	})
	t.Logf("%+v", err)
	if ClassOf(err) != ENil {
		t.Fail()
		return
	}
}

func TestNilCollector_Error(t *testing.T) {
	if err := (*Collector)(nil).Error(); err != nil {
		t.Logf("%+v", err)
		t.Fail()
		return
	}
}
