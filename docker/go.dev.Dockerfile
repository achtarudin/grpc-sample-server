FROM golang:1.24-alpine

WORKDIR /app

COPY . . 

RUN apk add --no-cache make

RUN make install-tools

RUN make install-deps

ENV PATH="/root/go/bin:${PATH}"

ENV GRPC_PORT=9000

EXPOSE ${GRPC_PORT}

CMD ["make", "dev-server"]