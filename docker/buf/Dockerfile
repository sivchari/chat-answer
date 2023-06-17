ARG GO_VERSION
FROM golang:${GO_VERSION}
ARG PROTOBUF_VERSION=23.3
ARG BUF_VERSION=1.21.0
ARG PROTOC_GEN_GO_VERSION=1.30.0
ARG PROTOC_GEN_CONNECT_GO_VERSION=1.8.0

RUN apt-get update && apt-get install unzip

WORKDIR /tmp/buf
RUN curl -sSL "https://github.com/bufbuild/buf/releases/download/v${BUF_VERSION}/buf-$(uname -s)-$(uname -m)" -o "/usr/local/bin/buf"
RUN chmod +x "/usr/local/bin/buf"

WORKDIR /usr/src/chat-answer
RUN go install "google.golang.org/protobuf/cmd/protoc-gen-go@v${PROTOC_GEN_GO_VERSION}"
RUN go install "github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@v${PROTOC_GEN_CONNECT_GO_VERSION}"

WORKDIR ${GOPATH}

ENTRYPOINT ["/usr/local/bin/buf"]
