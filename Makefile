PROGRAM:=./build/tinktur
GO_FILES:=$(shell find . -type f -name '*.go' -print)

.PHONY: build
build: $(PROGRAM)

$(PROGRAM): $(GO_FILES)
	go build -o $(PROGRAM)
