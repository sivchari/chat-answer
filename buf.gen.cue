#Plugin: {
	name:  "go" | "connect-go" | "es" | "connect-web"
	path?: string
	out:   string
	opt:   "paths=source_relative" | "target=ts"
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
}, {
	name: "es" // protoc-gen-es
	path: "./front/node_modules/.bin/protoc-gen-es"
	out:  "./front/src/proto"
	opt:  "target=ts"
}, {
	name: "connect-web" // protoc-gen-connect-web
	path: "./front/node_modules/.bin/protoc-gen-connect-web"
	out:  "./front/src/proto"
	opt:  "target=ts"
}]

version: "v1"
managed: {
	enabled: true
	go_package_prefix: {
		default: "github.com/sivchari/chat-answer/proto"
	}
}
plugins: _plugins
