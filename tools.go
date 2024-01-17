//go:build tools

package tools

import (
	_ "go.uber.org/mock/mockgen"
	_ "golang.org/x/tools/cmd/deadcode"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
	_ "sourcegraph.com/sourcegraph/prototools/cmd/protoc-gen-json"
)
