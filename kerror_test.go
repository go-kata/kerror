package kerror

import (
	"fmt"
	"testing"
)

func newTestNativeError(t *testing.T) error {
	err := fmt.Errorf("test error")
	t.Logf("%+v", err)
	return err
}

func newThreeTestErrors(t *testing.T) (err1, err2, err3 error) {
	err1 = fmt.Errorf("test error 1")
	err2 = Wrap(err1, Label("test.Error2"), "test error 2")
	err3 = Wrap(err2, Label("test.Error3"), "test error 3")
	t.Logf("%+v", err3)
	return
}

func TestClassOfWithNilError(t *testing.T) {
	if ClassOf(nil) != nil {
		t.Fail()
		return
	}
}

func TestClassOf__NativeError(t *testing.T) {
	err := newTestNativeError(t)
	if ClassOf(err) != nil {
		t.Fail()
		return
	}
}

func TestClassOf__PackageError(t *testing.T) {
	_, _, err3 := newThreeTestErrors(t)
	if ClassOf(err3) != Label("test.Error3") {
		t.Fail()
		return
	}
}

func TestMessageOf__NilError(t *testing.T) {
	if MessageOf(nil) != "" {
		t.Fail()
		return
	}
}

func TestMessageOf__NativeError(t *testing.T) {
	err := newTestNativeError(t)
	if MessageOf(err) != "test error" {
		t.Fail()
		return
	}
}

func TestMessageOf__PackageError(t *testing.T) {
	_, _, err3 := newThreeTestErrors(t)
	if MessageOf(err3) != "test error 3" {
		t.Fail()
		return
	}
}

func TestBaseOf__NilError(t *testing.T) {
	if Base(nil) != nil {
		t.Fail()
		return
	}
}

func TestBaseOf__NativeError(t *testing.T) {
	err := newTestNativeError(t)
	if Base(err) != err {
		t.Fail()
		return
	}
}

func TestBaseOf__PackageError(t *testing.T) {
	e1, _, e3 := newThreeTestErrors(t)
	if Base(e3) != e1 {
		t.Fail()
		return
	}
}

func TestIs__NilError(t *testing.T) {
	if Is(nil, Label("test.Error")) {
		t.Fail()
		return
	}
}

func TestIs__NilClass(t *testing.T) {
	if Is(New(Label("test.Error"), "test error"), nil) {
		t.Fail()
		return
	}
}

func TestIs__WrappedError(t *testing.T) {
	_, _, err3 := newThreeTestErrors(t)
	if !Is(err3, Label("test.Error3")) || !Is(err3, Label("test.Error2")) {
		t.Fail()
		return
	}
}

func TestIs__MultiError(t *testing.T) {
	errs := MultiError{
		New(Label("test.Error1"), "test error 1"),
		New(Label("test.Error2"), "test error 2"),
	}
	if !Is(errs, Label("test.Error1")) || !Is(errs, Label("test.Error2")) {
		t.Fail()
		return
	}
}

func TestJoin(t *testing.T) {
	err1 := New(Label("test.Error1"), "test error 1")
	err2 := New(Label("test.Error2"), "test error 2")
	err3 := New(Label("test.Error3"), "test error 3")
	err := Join(err1, err2, err3)
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
