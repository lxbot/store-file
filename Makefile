.PHONY: build debug

build:
	go build -buildmode=plugin -o store-file.so store.go

debug:
	go build -gcflags="all=-N -l" -buildmode=plugin -o store-file.so store.go
