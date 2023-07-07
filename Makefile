.PHONY: help
help: ## Display this help screen.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: tools
tools: ## Install required tools.
	echo 'Run go install' && \
	cd ./tools; \
	cat tools.go | grep "_" | awk -F'"' '{print $$2}' | xargs -tI % go install %@latest && \
	cd ../;

.PHONY: buf
buf: ## Generate protobuf codes.
	docker compose run --rm buf-go mod update
	docker compose run --rm buf-go generate --path proto --template buf.gen.go.yaml
	docker compose run --rm buf-go format proto -d -w > /dev/null
	gofmt -s -w proto/proto
	goimports -w -local "github.com/sivchari/chat-answer" proto/proto
	docker compose run --rm buf-ts mod update
	docker compose run --rm buf-ts generate --path proto --template buf.gen.ts.yaml

.PHONY: evans
evans: ## Run evans.
	evans --proto ./proto/healthz.proto --port 8080

.PHONY: run ## Serve api and front.
run: run-api run-front

.PHONY: run-api
run-api: ## Serve api.
	docker compose up api -d --build

.PHONY: run-front
run-front: ## Serve front.
	docker compose up front -d --build

.PHONY: go-generate
go-generate: ## Exec go generate.
	go generate ./pkg/...
