package kerror

import (
	"fmt"
	"testing"
)

func TestFormat(t *testing.T) {
	err1 := MultiError{
		Wrap(New(ECustom+1, "test error 1"), ECustom+2, "test error 2"),
		New(nil, "test error 3"),
		MultiError{
			New(ECustom+4, "test error 4"),
			New(ECustom+5, "test error 5"),
			MultiError{
				New(ECustom+6, "test error 6"),
			},
			fmt.Errorf("test error 7"),
			New(ECustom+8, "test error 8"),
		},
	}
	err2 := Wrap(err1, ECustom+9, "test error 9")
	err3 := Wrap(err2, ECustom+10, "test error 10")
	t.Logf("\n%+v\n", err3)
}
