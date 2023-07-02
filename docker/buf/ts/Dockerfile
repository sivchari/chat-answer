ARG NODE_VERSION=18
FROM node:${NODE_VERSION}

ARG BUF_VERSION=1.21.0

WORKDIR /tmp/buf
RUN curl -sSL "https://github.com/bufbuild/buf/releases/download/v${BUF_VERSION}/buf-$(uname -s)-$(uname -m)" -o "/usr/local/bin/buf"
RUN chmod +x "/usr/local/bin/buf"

WORKDIR /usr/src/chat-answer

RUN npm install -g @bufbuild/protoc-gen-es

ENTRYPOINT ["/usr/local/bin/buf"]
