BUILDDIR := builds
PROJECT_ROOT := $(shell pwd)

all: server util

server: prepare
	go build -o $(BUILDDIR)/server $(PROJECT_ROOT)/cmd/server

util: prepare
	go build -o $(BUILDDIR)/util $(PROJECT_ROOT)/cmd/util

prepare:
	mkdir -p $(BUILDDIR)

clean:
	rm -rf $(BUILDDIR)

.PHONY: all server util prepare clean
