package exceptions

import "reflect"

type baseError struct {
	message string
}

func (err *baseError) Error() string {
	return err.message
}

func (err *baseError) setMessage(message string) {
	err.message = message
}

type BadRequestError struct {
	baseError
}

type NotFoundError struct {
	baseError
}

type UnexpectedError struct {
	baseError
}

type appErrorType interface {
	error
	setMessage(string)
	*BadRequestError | *NotFoundError | *UnexpectedError
}

func NewAppError[T appErrorType](message string) error {
	err := reflect.New(reflect.TypeOf(new(T)).Elem().Elem()).Interface().(T)
	err.setMessage(message)
	return err
}

func IsAppError[T appErrorType](err error) bool {
	_, ok := err.(T)
	return ok
}
