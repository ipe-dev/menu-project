package errors

import (
	"fmt"
	"runtime"
)

type CustomError struct {
	Msg   string
	Value interface{}
}
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

type LoginPasswordError struct {
	Err   error
	Msg   string
	Value interface{}
}
type LoginNotFoundError struct {
	Msg   string
	Value interface{}
}
type ExistError struct {
	Text  string
	Msg   string
	Value interface{}
}

func NewCustomError(msg string, value ...interface{}) *CustomError {
	return &CustomError{msg, value}
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
	pc, file, line, _ := runtime.Caller(1)
	function := runtime.FuncForPC(pc)
	msg := fmt.Sprintf("call:%s file:%s:%d", function, file, line)

	return &ValidateError{e, msg, values}
}
func NewExistError(text string, values ...interface{}) *ExistError {
	pc, file, line, _ := runtime.Caller(1)
	function := runtime.FuncForPC(pc)
	msg := fmt.Sprintf("call:%s file:%s:%d", function, file, line)
	return &ExistError{text, msg, values}
}
func NewLoginPasswordError(e error, values ...interface{}) *LoginPasswordError {
	pc, file, line, _ := runtime.Caller(1)
	function := runtime.FuncForPC(pc)
	msg := fmt.Sprintf("call:%s file:%s:%d", function, file, line)
	return &LoginPasswordError{e, msg, values}
}
func NewLoginNotFoundError(msg string, values ...interface{}) *LoginNotFoundError {
	return &LoginNotFoundError{msg, values}
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("error:%s value:%#v", e.Msg, e.Value)
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
		e.Text,
		e.Msg,
		e.Value,
	)
}

func (e *LoginPasswordError) Error() string {
	return fmt.Sprintf("error:%s %#v value:%#v",
		e.Err.Error(),
		e.Msg,
		e.Value,
	)
}
func (e *LoginNotFoundError) Error() string {
	return fmt.Sprintf("error:%s value:%#v",
		e.Msg,
		e.Value,
	)
}
