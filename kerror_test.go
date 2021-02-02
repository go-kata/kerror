package kerror

import (
	"fmt"
	"testing"
)

func newTestNativeError(t *testing.T) error {
	err := fmt.Errorf("test error")
	t.Logf("\n%+v\n", err)
	return err
}

func newThreeTestErrors(t *testing.T) (err1, err2, err3 error) {
	err1 = fmt.Errorf("test error 1")
	err2 = Wrap(err1, ECustom+2, "test error 2")
	err3 = Wrap(err2, ECustom+3, "test error 3")
	t.Logf("\n%+v\n", err3)
	return
}

func TestClassOfWithNilError(t *testing.T) {
	if ClassOf(nil) != nil {
		t.Fail()
		return
	}
}

func TestClassOfWithNativeError(t *testing.T) {
	err := newTestNativeError(t)
	if ClassOf(err) != nil {
		t.Fail()
		return
	}
}

func TestClassOfWithPackageError(t *testing.T) {
	_, _, err3 := newThreeTestErrors(t)
	if ClassOf(err3) != ECustom+3 {
		t.Fail()
		return
	}
}

func TestMessageOfWithNilError(t *testing.T) {
	if MessageOf(nil) != "" {
		t.Fail()
		return
	}
}

func TestMessageOfWithNativeError(t *testing.T) {
	err := newTestNativeError(t)
	if MessageOf(err) != "test error" {
		t.Fail()
		return
	}
}

func TestMessageOfWithPackageError(t *testing.T) {
	_, _, err3 := newThreeTestErrors(t)
	if MessageOf(err3) != "test error 3" {
		t.Fail()
		return
	}
}

func TestIs(t *testing.T) {
	_, _, err3 := newThreeTestErrors(t)
	if Is(nil, nil) || Is(nil, ECustom+3) || Is(err3, nil) || !Is(err3, ECustom+3) || !Is(err3, ECustom+2) {
		t.Fail()
		return
	}
}

func TestBaseOfWithNilError(t *testing.T) {
	if Base(nil) != nil {
		t.Fail()
		return
	}
}

func TestBaseOfWithNativeError(t *testing.T) {
	err := newTestNativeError(t)
	if Base(err) != err {
		t.Fail()
		return
	}
}

func TestBaseOfWithPackageError(t *testing.T) {
	e1, _, e3 := newThreeTestErrors(t)
	if Base(e3) != e1 {
		t.Fail()
		return
	}
}
