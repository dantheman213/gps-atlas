BASE_NAME = gps-atlas
BIN_FILENAME = ""

ifeq ($(OS),Windows_NT)
	BIN_FILENAME = ${BASE_NAME}.exe
else
	BIN_FILENAME = ${BASE_NAME}
endif

.PHONY: all build clean deps

all: build

build:
	CGO_ENABLED=0 GO111MODULE=on go build -installsuffix "static" -o bin/$(BIN_FILENAME) $$(find cmd/app/*.go)

clean:
	@echo Cleaning bin/ directory... && \
		rm -rfv bin/

deps:
	@echo Downloading go.mod dependencies && \
		go mod download