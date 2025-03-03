GO := go

all: generate build test

build:
	go build -v ./...

test:
	go test -v ./...

generate:
	cd ./strimzi-go-generator; mvn -e -V -B compile exec:java -Dexec.mainClass="cz.scholz.strimzi.golang.generator.Main"
	./hack/codegen.sh
	go fmt ./...

.PHONY: all build test format generate