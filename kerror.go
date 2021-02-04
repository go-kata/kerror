// Package kerror provides tools for handling errors.
package kerror

// ClassOf returns a class of the given error.
func ClassOf(err error) Class {
	if err == nil {
		return nil
	}
	if e, ok := err.(interface{ Class() Class }); ok {
		return e.Class()
	}
	return nil
}

// MessageOf returns a message of the given error.
func MessageOf(err error) string {
	if err == nil {
		return ""
	}
	if e, ok := err.(interface{ Message() string }); ok {
		return e.Message()
	}
	return err.Error()
}

// Base returns an error on which the given error is based (the last error in a cause chain).
func Base(err error) error {
	if err == nil {
		return nil
	}
	if e, ok := err.(interface{ Cause() error }); ok {
		if cause := e.Cause(); cause != nil {
			return Base(cause)
		}
	}
	return err
}

// Is returns true if the given error (or any of errors that caused it) has the given class.
func Is(err error, class Class) bool {
	if err == nil || class == nil {
		return false
	}
	if e, ok := err.(interface{ Class() Class }); ok {
		if e.Class() == class {
			return true
		}
	}
	if e, ok := err.(interface{ Cause() error }); ok {
		if cause := e.Cause(); cause != nil {
			return Is(cause, class)
		}
	}
	if errs, ok := err.(MultiError); ok {
		for _, e := range errs {
			if Is(e, class) {
				return true
			}
		}
	}
	return false
}

// Join passes given errors through a collector.
func Join(errors ...error) error {
	coerr := NewCollector()
	coerr.Collect(errors...)
	return coerr.Error()
}

// Collect calls Join.
//
// Deprecated: since 0.2.0, use Join instead.
func Collect(errors ...error) error {
	return Join(errors...)
}
