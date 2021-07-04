package errors

import (
	"fmt"
	"runtime"
)

type InfraError struct {
	Err   error
	Msg   string
	Value interface{}
}

type ValidateError struct {
	Err   error
	Msg   string
	Value interface{}
}

type LoginError struct {
	Err   error
	Msg   string
	Value interface{}
}
type ExistError struct {
	Err   error
	Msg   string
	Value interface{}
}

func NewInfraError(e error, values ...interface{}) *InfraError {
	if e == nil {
		return nil
	}
	pc, file, line, _ := runtime.Caller(1)
	function := runtime.FuncForPC(pc)
	msg := fmt.Sprintf("call:%s file:%s:%d", function, file, line)
	return &InfraError{e, msg, values}
}
func NewValidateError(e error, values ...interface{}) *ValidateError {
	if e == nil {
		return nil
	}
	pc, file, line, _ := runtime.Caller(1)
	function := runtime.FuncForPC(pc)
	msg := fmt.Sprintf("call:%s file:%s:%d", function, file, line)

	return &ValidateError{e, msg, values}
}
func NewExistError(e error, msg string, values ...interface{}) *ExistError {
	return &ExistError{e, msg, values}
}
func NewLoginError(e error, msg string, values ...interface{}) *LoginError {
	return &LoginError{e, msg, values}
}

func (e *InfraError) Error() string {
	return fmt.Sprintf("error:%s %#v value:%#v",
		e.Err.Error(),
		e.Msg,
		e.Value,
	)
}

func (e *ValidateError) Error() string {
	return fmt.Sprintf("error:%s %#v value:%#v",
		e.Err.Error(),
		e.Msg,
		e.Value,
	)
}
func (e *ExistError) Error() string {
	return fmt.Sprintf("error:%s %#v value:%#v",
		e.Err.Error(),
		e.Msg,
		e.Value,
	)
}

func (e *LoginError) Error() string {
	return fmt.Sprintf("error:%s %#v value:%#v",
		e.Err.Error(),
		e.Msg,
		e.Value,
	)
}
