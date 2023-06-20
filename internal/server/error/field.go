package error

import "fmt"

type FieldError struct {
	err   error
	field string
	value interface{}
}

func New(field string, value interface{}, err error) FieldError {
	return FieldError{field: field, value: value, err: err}
}

func (err FieldError) Error() string {
	return fmt.Sprintf("Value '%s' of field '%s' is not correct\n %v", err.value, err.field, err.err)
}

func (err FieldError) Unwrap() error {
	return err.err
}
