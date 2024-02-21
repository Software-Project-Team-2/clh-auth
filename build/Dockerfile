FROM golang:1.21.5 as builder

WORKDIR /app

RUN apt-get update && apt-get install -y protobuf-compiler && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

ENV PATH="/root/go/bin:${PATH}"

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN make protoc

RUN make build

FROM debian:bullseye-slim

WORKDIR /

COPY --from=builder /app/bin/clh-auth .

CMD ["./clh-auth"]