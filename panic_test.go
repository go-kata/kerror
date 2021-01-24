package kerror

import "testing"

func TestPanic(t *testing.T) {
	const message = "keep calm, this is a test panic"
	defer func() {
		v := recover()
		t.Logf("\n%+v\n", v)
		if v == nil {
			t.Fail()
			return
		}
		err, ok := v.(error)
		if !ok {
			t.Fail()
			return
		}
		if ClassOf(err) != EPanic || MessageOf(err) != message {
			t.Fail()
			return
		}
	}()
	Panic(message)
}

func TestCatch(t *testing.T) {
	const errorMessage = "test error"
	const panicMessage = "keep calm, this is a test panic"
	defer func() {
		if v := recover(); v != nil {
			t.Fail()
			return
		}
	}()
	err := func() (err error) {
		defer Catch(&err)
		err = New(ECustom, errorMessage)
		Panic(panicMessage)
		return
	}()
	t.Logf("\n%+v\n", err)
	if err == nil {
		t.Fail()
		return
	}
	if errs, ok := err.(MultiError); !ok || len(errs) != 2 ||
		ClassOf(errs[0]) != ECustom || MessageOf(errs[0]) != errorMessage ||
		ClassOf(errs[1]) != EPanic || MessageOf(errs[1]) != panicMessage {
		t.Fail()
		return
	}
}

func TestCatchWithNil(t *testing.T) {
	defer func() {
		v := recover()
		t.Logf("\n%+v\n", v)
		if v == nil {
			t.Fail()
			return
		}
	}()
	_ = func() (err error) {
		defer Catch(nil)
		panic("keep calm, this is a test panic")
		return
	}()
}

func TestTryWithError(t *testing.T) {
	const class = ECustom + 1
	const message = "test error"
	err := Try(func() error {
		return New(class, message)
	})
	t.Logf("\n%+v\n", err)
	if err == nil || ClassOf(err) != class || MessageOf(err) != message {
		t.Fail()
		return
	}
}

func TestTryWithNativePanic(t *testing.T) {
	const message = "keep calm, this is a test panic"
	defer func() {
		if v := recover(); v != nil {
			t.Fail()
			return
		}
	}()
	err := Try(func() error {
		panic(message)
		return nil
	})
	t.Logf("\n%+v\n", err)
	if err == nil || ClassOf(err) != EPanic || MessageOf(err) != message {
		t.Fail()
		return
	}
}

func TestTryWithPackagePanic(t *testing.T) {
	const message = "keep calm, this is a test panic"
	defer func() {
		if v := recover(); v != nil {
			t.Fail()
			return
		}
	}()
	err := Try(func() error {
		Panic(message)
		return nil
	})
	t.Logf("\n%+v\n", err)
	if err == nil || ClassOf(err) != EPanic || MessageOf(err) != message {
		t.Fail()
		return
	}
}
