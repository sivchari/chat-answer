ARG GO_VERSION
ARG ALPINE_VERSION=3.18

# ===== build go binary =====
FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} as go-builder

WORKDIR /go/src/github.com/sivchari/chat-answer

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o api cmd/api/main.go

# ==== build docker image ====
FROM alpine:${ALPINE_VERSION}

WORKDIR /usr/src/chat-answer

COPY --from=go-builder /go/src/github.com/sivchari/chat-answer/api api
RUN chmod +x "/usr/src/chat-answer/api"

ENTRYPOINT ["/usr/src/chat-answer/api"]
