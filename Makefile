.PHONY: build

build:
	go build -buildmode=plugin -o store-file.so store.go
