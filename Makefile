.PHONY: test build cgo-build

BINARY := matrix
SRC := src/*.go
LDFLAGS := --ldflags '-w -s -extldflags=-static'
CGO_LDFLAGS := --ldflags '-w -s -linkmode external -extldflags=-static'

test: 
	echo "what"

build: 
	CGO_ENABLED=0 go build -o $(BINARY) $(LDFLAGS) $(SRC)

cgo-build: 
	CGO_ENABLED=1 CC=musl-gcc go build -o $(BINARY) $(CGO_LDFLAGS) $(SRC)

