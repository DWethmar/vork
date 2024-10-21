.PHONY: all build run clean

BINARY_NAME=vork

all: build

build:
	go build -o bin/$(BINARY_NAME)

run: build
	./bin/$(BINARY_NAME)

clean:
	rm -f bin/$(BINARY_NAME)