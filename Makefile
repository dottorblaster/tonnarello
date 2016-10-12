SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

build:
	go build -o target/tonnarello $(SOURCES)

test:
	go test

clean:
	rm -rf target/

all: build
