SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -name '*.go')

test:
	go test

all: # TODO
