package errors

import (
	"errors"
	"fmt"
	"runtime/debug"

	"github.com/sivchari/chat-answer/pkg/codes"
)

type Error struct {
	code    codes.Code
	origin  error
	stack   string
	callers []uintptr
}

func (e *Error) Message() string {
	if e.origin == nil {
		return e.code.String()
	}
	return e.code.String() + ": " + e.origin.Error()
}

func (e *Error) Error() string {
	if e.origin == nil {
		return fmt.Sprintf("%s: StackTrace:\n%s", e.code.String(), e.stack)
	}
	return fmt.Sprintf("%s: %s\nStackTrace:\n%s", e.code.String(), e.origin.Error(), e.stack)
}

func (e *Error) Stack() string {
	return e.stack
}

func (e *Error) Unwrap() error {
	return e.origin
}

func (e *Error) Callers() []uintptr {
	return e.callers
}

func New(code codes.Code, format string, a ...interface{}) error {
	stack := debug.Stack()
	return &Error{
		code:   code,
		origin: fmt.Errorf(format, a...),
		stack:  string(stack),
	}
}

func Code(err error) codes.Code {
	if err == nil {
		return codes.CodeOK
	}
	var e *Error
	if errors.As(err, &e) {
		return e.code
	}
	return codes.CodeUnknown
}

func Message(err error) string {
	if err == nil {
		return ""
	}

	var e *Error
	if errors.As(err, &e) {
		if e.origin == nil {
			return ""
		}
		return e.origin.Error()
	}

	return err.Error()
}
