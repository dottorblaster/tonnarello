SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go' -maxdepth 1)

build:
	go build -o target/tonnarello $(SOURCES)

run:
	go run $(SOURCES)

test:
	go test

clean:
	rm -rf target/

all: build
