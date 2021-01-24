package kerror

import "fmt"

// NPE triggers a null pointer exception (panic) with a stack trace.
func NPE() {
	panic(newErrorWithTrace(2, EPanic, "nil pointer dereference", nil))
}

// Panic triggers a panic with a stack trace and the given message.
func Panic(message string) {
	panic(newErrorWithTrace(2, EPanic, message, nil))
}

// Panicf is a variant of the Panic with message formatting.
func Panicf(format string, a ...interface{}) {
	panic(newErrorWithTrace(2, EPanic, fmt.Sprintf(format, a...), nil))
}

// PanicWrap triggers a panic with providing an original error that caused it.
//
// A stack trace will be provided only if an original error is of an external type.
func PanicWrap(err error, message string) {
	panic(wrapError(2, err, EPanic, message))
}

// PanicWrapf is a variant of the PanicWrap with message formatting.
func PanicWrapf(err error, format string, a ...interface{}) {
	panic(wrapError(2, err, EPanic, fmt.Sprintf(format, a...)))
}

// Try calls f, recovers a panic if occurred and returns it as error.
func Try(f func() error) (err error) {
	defer Catch(&err)
	return f()
}

// Catch recovers a panic if occurred and saves it to err.
func Catch(err *error) {
	if err == nil {
		return
	}
	if v := recover(); v != nil {
		coerr := NewCollector()
		coerr.Collect(*err)
		if e, ok := v.(error); ok {
			coerr.Collect(e)
		} else {
			coerr.Collect(newErrorWithTrace(3, EPanic, fmt.Sprintf("%v", v), nil))
		}
		*err = coerr.Error()
	}
}
