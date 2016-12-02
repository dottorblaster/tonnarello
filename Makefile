SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -maxdepth 1 -name '*.go')

build:
	go build -o target/tonnarello $(SOURCES)

run:
	go run $(SOURCES)

test:
	go test

clean:
	rm -rf target/

all: build
