package errcodes

import (
	"fmt"

	"google.golang.org/grpc/codes"
)

type Code int

const (
	CodeUnknown = iota
	CodeOK
	CodeInvalidArgument
	CodeNotFound
	CodeInternal
)

func (c Code) String() string {
	switch c {
	case CodeUnknown:
		return "Unknown"
	case CodeInvalidArgument:
		return "Invalid argument"
	case CodeNotFound:
		return "Not found"
	case CodeInternal:
		return "Internal"
	case CodeOK:
		return "OK"
	}
	return fmt.Sprintf("Unknown: %d", c)
}

func (c Code) GoString() string {
	return "errcode.Code[" + c.String() + "]"
}

func (c Code) grpcCode() codes.Code {
	switch c {
	case CodeUnknown:
		return codes.Unknown
	case CodeInvalidArgument:
		return codes.InvalidArgument
	case CodeNotFound:
		return codes.NotFound
	case CodeInternal:
		return codes.Internal
	case CodeOK:
		return codes.OK
	}
	return codes.Unknown
}
