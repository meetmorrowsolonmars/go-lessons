
LOCAL_BIN:=$(CURDIR)/bin
BUF_VERSION:=v1.45.0

.PHONY: install-bin-deps
install-bin-deps:
	$(info Installing binary dependencies...)

	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.15.2 && \
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.15.2 && \
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.30.0 && \
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0 && \
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@v1.1.0

	curl -sSL \
		"https://github.com/bufbuild/buf/releases/download/$(BUF_VERSION)/buf-$$(uname -s)-$$(uname -m)" \
		-o "$(LOCAL_BIN)/buf" && \
	chmod +x "$(LOCAL_BIN)/buf"

.PHONY: generate
generate:
	PATH="$(LOCAL_BIN):$(PATH)" buf dep update
	PATH="$(LOCAL_BIN):$(PATH)" buf generate
	go mod tidy
