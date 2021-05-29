package kerror

// Class represents an error class.
//
// Implement this interface in your package to provide guaranteed to be unique error classes.
type Class interface {
	// ErrorClass returns a string representation of this class.
	//
	// Builtin error classes all have the "common.ClassName" format. Come up with your own format
	// which will not mislead readers about the class source.
	ErrorClass() string
}

// Label represents a string error class (error label).
//
// Use this type in your package but be sure that your own labels don't collide
// with labels from other packages uses this way of the error class definition.
type Label string

// ErrorClass implements the Class interface.
func (l Label) ErrorClass() string {
	return string(l)
}

// EAmbiguous specifies the error label indicating an ambiguous entity.
//
// For example, error of this class may be returned on a try to overwrite existing map key.
const EAmbiguous Label = "common.Ambiguous"

// EIllegal specifies the error label indicating value or operation that are illegal in the current context.
//
// For example, error of this class may be returned from the Close method when it called again.
const EIllegal Label = "common.Illegal"

// EInvalid specifies the error label indicating an invalid value.
const EInvalid Label = "common.Invalid"

// ENil specifies the error label indicating an unacceptable operation on nil.
const ENil Label = "common.Nil"

// ENotFound specifies the error label indicating that a required entity is not found.
const ENotFound Label = "common.NotFound"

// EPanic specifies the panic label.
const EPanic Label = "common.Panic"

// ERuntime specifies the error label indicating an unspecific runtime error.
const ERuntime Label = "common.Runtime"

// ESystem specifies the error label indicating a system error.
//
// For example, error of this class may be used to wrap a file system error.
const ESystem Label = "common.System"

// EViolation specifies the error label indicating a contract violation
// which may not be detected at the compile time for some reason.
const EViolation Label = "common.Violation"
