SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

build:
	go build $(SOURCES) -o target/tonnarello

test:
	go test

clean:
	rm -rf target/

all: build
