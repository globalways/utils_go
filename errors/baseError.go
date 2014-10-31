// Package: errors
// File: baseError.go
// Created by mint
// Useage: baseError
// DATE: 14-7-9 10:14
package errors

import "fmt"

// Standard struct for general types of errors.
//
// For an example of custom error type, look at databaseError/newDatabaseError
// in errors_test.go.
type GlobalWaysBaseError struct {
	Msg     string
	Stack   string
	Context string
	inner   error
	Code    int
}

// This returns a string with all available error information, including inner
// errors that are wrapped by this errors.
func (e *GlobalWaysBaseError) Error() string {
	return DefaultError(e)
}

// This returns the error message without the stack trace.
func (e *GlobalWaysBaseError) GetMessage() string {
	return e.Msg
}

// This returns the stack trace without the error message.
func (e *GlobalWaysBaseError) GetStack() string {
	return e.Stack
}

// This returns the stack trace's context.
func (e *GlobalWaysBaseError) GetContext() string {
	return e.Context
}

// This returns the wrapped error, if there is one.
func (e *GlobalWaysBaseError) GetInner() error {
	return e.inner
}

// This returns the error code
func (e *GlobalWaysBaseError) GetCode() int {
	return e.Code
}

func (e *GlobalWaysBaseError) IsError() bool {
	if e.GetCode() == CODE_SUCCESS {
		return false
	}

	return true
}

func (e *GlobalWaysBaseError) ErrorMessage() string {
	return GetMessage(e)
}

// This returns a new GlobalWaysBaseError initialized with the given message and
// the current stack trace.
func New(code int, msg string) GlobalWaysError {
	stack, context := StackTrace()
	return &GlobalWaysBaseError{
		Msg:     msg,
		Stack:   stack,
		Context: context,
		Code:    code,
	}
}

// Same as New, but with fmt.Printf-style parameters.
func Newf(code int, format string, args ...interface{}) GlobalWaysError {
	stack, context := StackTrace()
	return &GlobalWaysBaseError{
		Msg:     fmt.Sprintf(format, args...),
		Stack:   stack,
		Context: context,
		Code:    code,
	}
}

// Wraps another error in a new GlobalWaysBaseError.
func Wrap(code int, err error, msg string) GlobalWaysError {
	stack, context := StackTrace()
	return &GlobalWaysBaseError{
		Msg:     msg,
		Stack:   stack,
		Context: context,
		inner:   err,
		Code:    code,
	}
}

// Same as Wrap, but with fmt.Printf-style parameters.
func Wrapf(code int, err error, format string, args ...interface{}) GlobalWaysError {
	stack, context := StackTrace()
	return &GlobalWaysBaseError{
		Msg:     fmt.Sprintf(format, args...),
		Stack:   stack,
		Context: context,
		inner:   err,
		Code:    code,
	}
}

// equeal error = nil
func ErrorOK() GlobalWaysError {
	return &GlobalWaysBaseError{
		Msg:     "",
		Stack:   "",
		Context: "",
		inner:   nil,
		Code:    CODE_SUCCESS,
	}
}
