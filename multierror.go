package kerror

import (
	"strconv"
	"strings"
)

// MultiError represents a multiple error.
//
// The usual identifier for variables of this type is errs.
type MultiError []error

// String implements the fmt.Stringer interface.
func (e MultiError) String() string {
	return e.stringImpl(0)
}

// stringImpl returns a string representation of this error
// with list items padded by the given number of tab characters.
func (e MultiError) stringImpl(padding int) string {
	if len(e) == 0 {
		return ""
	}
	var p string
	if padding > 0 {
		p = strings.Repeat("\t", padding)
	}
	var s string
	if n := len(e); n == 1 {
		s = "1 error occurred:"
	} else {
		s = strconv.Itoa(n) + " errors occurred:"
	}
	for i, err := range e {
		s += "\n" + p + "\t#" + strconv.Itoa(i+1) + " 🠖 "
		if err != nil {
			if errs, ok := err.(MultiError); ok {
				s += errs.stringImpl(padding + 1)
			} else {
				s += err.Error()
			}
		} else {
			s += "<nil>"
		}
	}
	return s
}

// Error implements the error interface.
func (e MultiError) Error() string {
	return e.String()
}

// Traverse performs the error graph traversal in depth
// and calls f for each non-nil and non-multiple error.
//
// The traversal will be broken if f will return false.
func Traverse(err error, f func(err error) (next bool)) {
	traverseImpl(err, f)
}

// traverseImpl is the internal implementation of the Traverse function.
func traverseImpl(err error, f func(err error) (next bool)) (next bool) {
	if err == nil {
		return true
	}
	if errs, ok := err.(MultiError); ok {
		for _, e := range errs {
			if !traverseImpl(e, f) {
				return false
			}
		}
		return true
	}
	return f(err)
}
