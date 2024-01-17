.PHONY: build protoc

build:
	go build -o bin/whatever cmd/main.go

PROTO_DIR=./api/proto
PROTO_OUT_DIR=./internal/pb

protoc:
	protoc --go_out=$(PROTO_OUT_DIR) --go_opt=paths=source_relative \
		--proto_path=$(PROTO_DIR) $(PROTO_DIR)/auth/*.proto
