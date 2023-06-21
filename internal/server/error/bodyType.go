package error

import "fmt"

type BodyTypeError struct {
	err   error
	value interface{}
}

func NewBodyTypeError(value interface{}, err error) BodyTypeError {
	return BodyTypeError{value: value, err: err}
}

func (err BodyTypeError) Error() string {
	return fmt.Sprintf("Value '%s' of field 'body_type' is not correct\n %v", err.value, err.err)
}

func (err BodyTypeError) Unwrap() error {
	return err.err
}
