package kerror

import "fmt"

// NPE triggers the nil pointer exception (panic) with a stack trace.
//
// Deprecated: since 0.2.1, use Nil instead.
func NPE() {
	panic(newErrorWithTrace(2, EPanic, "nil pointer dereference", nil))
}

// Nil triggers a panic by reason of the nil pointer dereference.
func Nil() {
	panic(newErrorWithTrace(2, ENil, "nil pointer dereference", nil))
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

// Recovered casts the given untyped result of the panic recovering to error.
// If the argument v is nil then nil will be returned. Otherwise error will be
// returned (is v doesn't implement the error interface it will be transformed
// to the native error).
func Recovered(v interface{}) error {
	if v == nil {
		return nil
	}
	if err, ok := v.(error); ok {
		return err
	}
	return fmt.Errorf("%+v", v)
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
			coerr.Collect(fmt.Errorf("%+v", v))
		}
		*err = coerr.Error()
	}
}
