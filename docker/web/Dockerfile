ARG BUN_VERSION=0.6.9

FROM oven/bun:${BUN_VERSION}

WORKDIR /usr/src/chat-answer/front

COPY . .
RUN bun install

CMD ["bun", "run", "dev"]
