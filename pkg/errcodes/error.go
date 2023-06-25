package errcodes

import (
	"errors"
	"fmt"
	"runtime/debug"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Error struct {
	Code    Code
	origin  error
	stack   string
	callers []uintptr
}

func (e *Error) Message() string {
	if e.origin == nil {
		return e.Code.String()
	}
	return e.Code.String() + ": " + e.origin.Error()
}

func (e *Error) Error() string {
	if e.origin == nil {
		return fmt.Sprintf("%s: StackTrace:\n%s", e.Code.String(), e.stack)
	}
	return fmt.Sprintf("%s: %s\nStackTrace:\n%s", e.Code.String(), e.origin.Error(), e.stack)
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

func New(code Code, format string, a ...interface{}) error {
	stack := debug.Stack()
	return &Error{
		Code:   code,
		origin: fmt.Errorf(format, a...),
		stack:  string(stack),
	}
}

func NewGrpcError(err error) error {
	if err == nil {
		return nil
	}

	var e *Error
	if !errors.As(err, &e) {
		return status.Error(codes.Unknown, "unexpected error")
	}
	return status.Error(e.Code.grpcCode(), e.Message())
}

func NewCode(err error) Code {
	if err == nil {
		return CodeOK
	}
	var e *Error
	if errors.As(err, &e) {
		return e.Code
	}
	return CodeUnknown
}
