//go:build tools

package tools

import (
	_ "github.com/pressly/goose/v3/cmd/goose"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
	_ "sourcegraph.com/sourcegraph/prototools/cmd/protoc-gen-json"
)
