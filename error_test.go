package kerror

import "testing"

func TestWrap__NilError(t *testing.T) {
	defer func() {
		v := recover()
		t.Logf("%+v", v)
		if v == nil {
			t.Fail()
			return
		}
	}()
	_ = Wrap(nil, nil, "")
}

func TestNilError_Class(t *testing.T) {
	if (*Error)(nil).Class() != nil {
		t.Fail()
		return
	}
}

func TestNilError_Message(t *testing.T) {
	if (*Error)(nil).Message() != "" {
		t.Fail()
		return
	}
}

func TestNilError_Cause(t *testing.T) {
	if (*Error)(nil).Cause() != nil {
		t.Fail()
		return
	}
}

func TestNilError_String(t *testing.T) {
	if (*Error)(nil).String() != "" {
		t.Fail()
		return
	}
}
