package kerror

// Collect calls Join.
//
// Deprecated: since 0.2.0, use Join instead.
func Collect(errors ...error) error {
	return Join(errors...)
}
