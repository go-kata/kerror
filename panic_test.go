package kerror

import "testing"

func TestNil(t *testing.T) {
	defer func() {
		v := recover()
		t.Logf("%+v", v)
		if ClassOf(Recovered(v)) != ENil {
			t.Fail()
			return
		}
	}()
	Nil()
}

func TestPanic(t *testing.T) {
	const message = "keep calm, this is a test panic"
	defer func() {
		v := recover()
		t.Logf("%+v", v)
		err := Recovered(v)
		if ClassOf(err) != EPanic || MessageOf(err) != message {
			t.Fail()
			return
		}
	}()
	Panic(message)
}

func TestRecovered__Nil(t *testing.T) {
	defer func() {
		if Recovered(recover()) != nil {
			t.Fail()
			return
		}
	}()
}

func TestRecovered__String(t *testing.T) {
	const message = "keep calm, this is a test panic"
	defer func() {
		v := recover()
		t.Logf("%+v", v)
		err := Recovered(v)
		if ClassOf(err) != nil || MessageOf(err) != message {
			t.Fail()
			return
		}
	}()
	panic(message)
}

func TestRecovered__PackageError(t *testing.T) {
	const class = ECustom
	const message = "keep calm, this is a test panic"
	defer func() {
		v := recover()
		t.Logf("%+v", v)
		err := Recovered(v)
		if ClassOf(err) != class || MessageOf(err) != message {
			t.Fail()
			return
		}
	}()
	panic(New(class, message))
}

func TestTry__Error(t *testing.T) {
	const class = ECustom
	const message = "test error"
	err := Try(func() error {
		return New(class, message)
	})
	t.Logf("%+v", err)
	if err == nil || ClassOf(err) != class || MessageOf(err) != message {
		t.Fail()
		return
	}
}

func TestTry__NativePanic(t *testing.T) {
	const message = "keep calm, this is a test panic"
	err := Try(func() error {
		panic(message)
		return nil
	})
	t.Logf("%+v", err)
	if err == nil || ClassOf(err) != nil || MessageOf(err) != message {
		t.Fail()
		return
	}
}

func TestTry__PackagePanic(t *testing.T) {
	const message = "keep calm, this is a test panic"
	err := Try(func() error {
		Panic(message)
		return nil
	})
	t.Logf("%+v", err)
	if err == nil || ClassOf(err) != EPanic || MessageOf(err) != message {
		t.Fail()
		return
	}
}

func TestCatch(t *testing.T) {
	const errorMessage = "test error"
	const panicMessage = "keep calm, this is a test panic"
	err := func() (err error) {
		defer Catch(&err)
		err = New(ECustom, errorMessage)
		Panic(panicMessage)
		return
	}()
	t.Logf("%+v", err)
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

func TestCatch__NilTarget(t *testing.T) {
	defer func() {
		v := recover()
		t.Logf("%+v", v)
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
