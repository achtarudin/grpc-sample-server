
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum Makefile ./

RUN apk add --no-cache make

COPY . .

RUN make build-server



FROM alpine:3.22

WORKDIR /app

# result build name is grpc-sample-server in Makefile
COPY --from=builder /app/bin/grpc-sample-server .

ENV GRPC_PORT=9000

EXPOSE ${GRPC_PORT}

CMD ["./grpc-sample-server"]
