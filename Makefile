VERSION=0.1.0

build: fmt
	@go build -o bin/kubenorm -gcflags "all=-trimpath=$$PWD" -ldflags "-s -w -X main.Version=${VERSION}" main.go
fmt:
	@go fmt ./... > /dev/null
