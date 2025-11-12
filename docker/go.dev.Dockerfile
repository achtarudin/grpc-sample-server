FROM golang:1.24-alpine

WORKDIR /app

RUN apk add --no-cache make

ENV PATH="/root/go/bin:${PATH}"


COPY Makefile go.mod go.sum ./

RUN make install-tools
RUN make install-deps

COPY . .

ENV GRPC_PORT=9000

EXPOSE ${GRPC_PORT}

CMD ["make", "dev-server"]
