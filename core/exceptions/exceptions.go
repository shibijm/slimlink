package exceptions

import "reflect"

type BaseError struct {
	message string
}

func (baseError *BaseError) Error() string {
	return baseError.message
}

func (baseError *BaseError) SetMessage(message string) {
	baseError.message = message
}

type BadRequestError struct {
	BaseError
}

type NotFoundError struct {
	BaseError
}

type UnexpectedError struct {
	BaseError
}

type AppError interface {
	error
	SetMessage(string)
}

type AppErrorType interface {
	AppError
	*BadRequestError | *NotFoundError | *UnexpectedError
}

func NewAppError[T AppErrorType](message string) T {
	err := reflect.New(reflect.TypeOf(new(T)).Elem().Elem()).Interface().(T)
	err.SetMessage(message)
	return err
}

func IsAppError[T AppErrorType](suspect any) bool {
	_, ok := suspect.(T)
	return ok
}
