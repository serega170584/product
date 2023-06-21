package error

import "fmt"

type ToError struct {
	err   error
	value []string
}

func NewToError(value []string, err error) ToError {
	return ToError{value: value, err: err}
}

func (err ToError) Error() string {
	return fmt.Sprintf("Value '%s' of field 'to' is not correct\n %v", err.value, err.err)
}

func (err ToError) Unwrap() error {
	return err.err
}
