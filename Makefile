.PHONY: help
help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: tools
tools: ## Install required tools.
	echo 'Run go install' && \
	cd ./tools; \
	cat tools.go | grep "_" | awk -F'"' '{print $$2}' | xargs -tI % go install %@latest && \
	cd ../;

.PHONY: buf
buf: ## Generate protobuf codes.
	docker compose run --rm buf mod update
	docker compose run --rm buf generate --path proto --template buf.gen.yaml
	gofmt -s -w proto/proto
	go install golang.org/x/tools/cmd/goimports@latest
	goimports -w -local "github.com/sivchari/chat-answer" proto/proto
	docker compose run --rm buf format proto -d -w

.PHONY: run-api
run-api: ## Serve api.
	docker compose up api -d --build

.PHONY: run-front
run-front: ## Serve front.
	docker compose up front -d --build

.PHONY: go-generate
go-generate: ## Exec go generate.
	go generate ./pkg/...
