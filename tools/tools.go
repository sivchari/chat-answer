//go:build tools

package tools

import (
	_ "github.com/bufbuild/buf/cmd/buf"
	_ "github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go"
	_ "github.com/ktr0731/evans"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
