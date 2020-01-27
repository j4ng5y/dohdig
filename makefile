VERSION = v0.2.3
DOHDIGBIN = dohdig_${VERSION}

all: test build

@PHONY: test
test:
	go test ./...

@PHONY: build
build: build_linux build_darwin build_windows

@PHONY: build_linux
build_linux:
	CGO_ENABLE=0 GOOS=linux GOARCH=386 go build -a -o bin/${DOHDIGBIN}_linux_386
	CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -a -o bin/${DOHDIGBIN}_linux_amd64

@PHONY: build_darwin
build_darwin:
	CGO_ENABLE=0 GOOS=darwin GOARCH=386 go build -a -o bin/${DOHDIGBIN}_darwin_386
	CGO_ENABLE=0 GOOS=darwin GOARCH=amd64 go build -a -o bin/${DOHDIGBIN}_darwin_amd64

@PHONY: build_windows
build_windows:
	CGO_ENABLE=0 GOOS=windows GOARCH=386 go build -a -o bin/${DOHDIGBIN}_win32.exe
	CGO_ENABLE=0 GOOS=windows GOARCH=amd64 go build -a -o bin/${DOHDIGBIN}_win64.exe