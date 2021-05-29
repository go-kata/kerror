package kerror

import (
	"fmt"
	"testing"
)

func TestFormat(t *testing.T) {
	err1 := MultiError{
		Wrap(New(Label("test.Error1"), "test error 1"), Label("test.Error2"), "test error 2"),
		New(nil, "test error 3"),
		MultiError{
			New(Label("test.Error4"), "test error 4"),
			New(Label("test.Error5"), "test error 5"),
			MultiError{
				New(Label("test.Error6"), "test error 6"),
			},
			fmt.Errorf("test error 7"),
			New(Label("test.Error8"), "test error 8"),
		},
	}
	err2 := Wrap(err1, Label("test.Error9"), "test error 9")
	err3 := Wrap(err2, Label("test.Error10"), "test error 10")
	t.Logf("\n%+v", err3)
}
