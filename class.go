package kerror

import "fmt"

// Class represents an error class.
//
// Implement this interface in your package to provide guaranteed to be unique error classes.
type Class interface {
	// ErrorClass returns a string representation of this class.
	//
	// Builtin error classes all have the EXXXXXXXXX format. Come up with your own format
	// which will not mislead readers about the class source.
	ErrorClass() string
}

// Number represents a numeric error class (error number).
//
// Use this type in your package but be sure that your own numbers don't collide
// with numbers from other packages uses this way of the error class definition.
type Number uint32

// ErrorClass implements the Class interface.
func (n Number) ErrorClass() string {
	return fmt.Sprintf("E%09d", n)
}

// EPanic specifies the panic number.
const EPanic Number = 0

// ENil specifies the error number indicating an unacceptable operation on nil.
const ENil Number = 1

// EViolation specifies the error number indicating a contract violation
// which may not be detected at the compile time for some reason.
const EViolation Number = 2

// ERuntime specifies the error number indicating an unspecific runtime error.
const ERuntime Number = 100

// EInvalid specifies the error number indicating an invalid value.
const EInvalid Number = 101

// EIllegal specifies the error number indicating value or operation that are illegal in the current context.
//
// For example, error of this class may be returned from the Close method when it called again.
const EIllegal Number = 102

// ENotFound specifies the error number indicating that a required entity is not found.
const ENotFound Number = 103

// EAmbiguous specifies the error number indicating an ambiguous entity.
//
// For example, error of this class may be returned on a try to overwrite existing map key.
const EAmbiguous Number = 104

// ECustom specifies the base custom error number (65536).
//
// The range 0-65535 is reserved for builtin numbers and for errors in framework components.
//
// Use something like this to define numbers in your package:
//
//     const (
//         EMyFirstError = ECustom + iota
//         EMySecondError
//         EMyThirdError
//     )
//
const ECustom Number = 0x00010000
