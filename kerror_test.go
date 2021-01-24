package kerror

import (
	"fmt"
	"testing"
)

func testNewNativeError(t *testing.T) error {
	err := fmt.Errorf("test error")
	t.Logf("\n%+v\n", err)
	return err
}

func testNewThreeError(t *testing.T) (err1, err2, err3 error) {
	err1 = fmt.Errorf("test error 1")
	err2 = Wrap(err1, ECustom+2, "test error 2")
	err3 = Wrap(err2, ECustom+3, "test error 3")
	t.Logf("\n%+v\n", err3)
	return
}

func TestClassOfNilError(t *testing.T) {
	if ClassOf(nil) != nil {
		t.Fail()
		return
	}
}

func TestClassOfNativeError(t *testing.T) {
	err := testNewNativeError(t)
	if ClassOf(err) != nil {
		t.Fail()
		return
	}
}

func TestClassOfPackageError(t *testing.T) {
	_, _, err3 := testNewThreeError(t)
	if ClassOf(err3) != ECustom+3 {
		t.Fail()
		return
	}
}

func TestMessageOfNilError(t *testing.T) {
	if MessageOf(nil) != "" {
		t.Fail()
		return
	}
}

func TestMessageOfNativeError(t *testing.T) {
	err := testNewNativeError(t)
	if MessageOf(err) != "test error" {
		t.Fail()
		return
	}
}

func TestMessageOfPackageError(t *testing.T) {
	_, _, err3 := testNewThreeError(t)
	if MessageOf(err3) != "test error 3" {
		t.Fail()
		return
	}
}

func TestIs(t *testing.T) {
	_, _, err3 := testNewThreeError(t)
	if Is(nil, nil) || Is(nil, ECustom+3) || Is(err3, nil) || !Is(err3, ECustom+3) || !Is(err3, ECustom+2) {
		t.Fail()
		return
	}
}

func TestBaseOfNilError(t *testing.T) {
	if Base(nil) != nil {
		t.Fail()
		return
	}
}

func TestBaseOfNativeError(t *testing.T) {
	err := testNewNativeError(t)
	if Base(err) != err {
		t.Fail()
		return
	}
}

func TestBaseOfPackageError(t *testing.T) {
	e1, _, e3 := testNewThreeError(t)
	if Base(e3) != e1 {
		t.Fail()
		return
	}
}
