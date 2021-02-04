package kerror

// Collector represents an error collector.
//
// The usual identifier for variables of this type is coerr.
type Collector struct {
	// errors specifies collected errors.
	errors []error
}

// NewCollector returns a new error collector.
func NewCollector() *Collector {
	return &Collector{}
}

// Collect collects given errors skipping nils and unfolding multiple errors.
func (c *Collector) Collect(errors ...error) {
	if c == nil {
		NPE()
		return
	}
	for _, err := range errors {
		if err == nil {
			continue
		}
		if errs, ok := err.(MultiError); ok {
			c.Collect(errs...)
		} else {
			c.errors = append(c.errors, err)
		}
	}
	return
}

// Error returns nil if no errors was collected, a single error if only one was collected
// or a multiple error which represents all collected errors.
func (c *Collector) Error() error {
	if c == nil || len(c.errors) == 0 {
		return nil
	}
	if len(c.errors) == 1 {
		return c.errors[0]
	}
	errs := make(MultiError, len(c.errors))
	copy(errs, c.errors)
	return errs
}
