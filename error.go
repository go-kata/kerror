package kerror

import (
	"fmt"
	"io"
	"runtime"
)

// Error represents an error.
//
// The usual identifier for variables of this type is err.
type Error struct {
	// class specifies the error class for a programmatic handling (may be nil).
	class Class
	// message specifies the human readable error message.
	message string
	// cause specifies the original error that caused this error (may be nil).
	cause error
	// trace specifies the array of program counters of function invocations
	// on the erroneous goroutine's stack (may be empty).
	trace []uintptr
}

// newError returns a new error without a stack trace.
func newError(class Class, message string, cause error) *Error {
	return &Error{
		class:   class,
		message: message,
		cause:   cause,
		trace:   nil,
	}
}

// newErrorWithTrace returns a new error with a stack trace.
//
// The argument skip specifies the number of stack frames to skip before recording,
// with 0 identifying the frame for newErrorWithTrace itself and 1 identifying
// the caller of newErrorWithTrace.
func newErrorWithTrace(skip int, class Class, message string, cause error) *Error {
	e := newError(class, message, cause)
	// There we limit recorded frames number up to 32
	// that is probably a sufficient stack depth in most cases.
	pc := make([]uintptr, 32)
	if n := runtime.Callers(skip+1, pc); n > 0 {
		e.trace = pc[:n]
	}
	return e
}

// wrapError returns a new error with providing an original error that caused it.
//
// A stack trace will be provided only if an original error is of an external type.
//
// The argument skip specifies the number of stack frames to skip before recording,
// with 0 identifying the frame for wrapError itself and 1 identifying
// the caller of wrapError.
func wrapError(skip int, err error, class Class, message string) *Error {
	if err == nil {
		Panic("nil error cannot be wrapped")
		return nil
	}
	if _, ok := err.(*Error); ok {
		return newError(class, message, err)
	}
	return newErrorWithTrace(skip+1, class, message, err)
}

// New returns a new error with a stack trace and given class and message.
func New(class Class, message string) *Error {
	return newErrorWithTrace(2, class, message, nil)
}

// Newf is a variant of the New with message formatting.
func Newf(class Class, format string, a ...interface{}) *Error {
	return newErrorWithTrace(2, class, fmt.Sprintf(format, a...), nil)
}

// Wrap is like the New but provides an original error that caused this.
//
// A stack trace will be provided only if an original error is of external type.
func Wrap(err error, class Class, message string) *Error {
	return wrapError(2, err, class, message)
}

// Wrapf is a variant of the Wrap with message formatting.
func Wrapf(err error, class Class, format string, a ...interface{}) *Error {
	return wrapError(2, err, class, fmt.Sprintf(format, a...))
}

// Class returns the error class for a programmatic handling (may be nil).
func (e *Error) Class() Class {
	if e == nil {
		return nil
	}
	return e.class
}

// Message returns the human readable error message.
func (e *Error) Message() string {
	if e == nil {
		return ""
	}
	return e.message
}

// Cause returns the original error that caused this error (may be nil).
func (e *Error) Cause() error {
	if e == nil {
		return nil
	}
	return e.cause
}

// String implements the fmt.Stringer interface.
func (e *Error) String() string {
	if e == nil {
		return ""
	}
	if e.class != nil {
		return e.class.ErrorClass() + ": " + e.message
	}
	return e.message
}

// Error implements the error interface.
func (e *Error) Error() string {
	return e.String()
}

// Format implements the fmt.Formatter interface.
//
// Following formats are supported:
//
//     %s (or %v) prints error message;
//
//     %q prints error message enclosed in double quotes;
//
//     %+v prints non-nil error class, error message and
//     (on next lines) an original error using %+v format and
//     a stack trace if provided.
//
func (e *Error) Format(f fmt.State, c rune) {
	if e == nil {
		_, _ = io.WriteString(f, "<nil>")
		return
	}
	switch c {
	case 'v':
		if f.Flag('+') {
			if e.cause != nil {
				_, _ = fmt.Fprintf(f, "%s\n⤷ %+v", e.String(), e.cause)
			} else {
				_, _ = io.WriteString(f, e.String())
			}
			if e.trace != nil && len(e.trace) > 0 {
				_, _ = io.WriteString(f, "\nStack trace:")
				frames := runtime.CallersFrames(e.trace)
				for {
					frame, more := frames.Next()
					_, _ = fmt.Fprintf(f, "\n\tat %s (%s:%d)", frame.Function, frame.File, frame.Line)
					if !more {
						break
					}
				}
			}
			return
		}
		fallthrough
	case 's':
		_, _ = io.WriteString(f, e.String())
	case 'q':
		_, _ = fmt.Fprintf(f, "%q", e.String())
	}
}
