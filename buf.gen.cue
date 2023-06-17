#Plugin: {
	name: "go" | "connect-go"
	out:  "proto"
	opt:  "paths=source_relative"
}

_plugins: [...#Plugin]
_plugins: [{
	name: "go" // protoc-gen-go
	out:  "proto"
	opt:  "paths=source_relative"
}, {
	name: "connect-go" // protoc-gen-connect-go
	out:  "proto"
	opt:  "paths=source_relative"
}]

version: "v1"
managed: {
	enabled: true
	go_package_prefix: {
		default: "github.com/sivchari/chat-answer/proto"
	}
}
plugins: _plugins
