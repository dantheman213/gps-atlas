BIN_FILENAME = gps-plotter
ifeq ($(OS),Windows_NT)
	BIN_FILENAME = gps-plotter.exe
endif

.PHONY: all build

all: build

build:
	CGO_ENABLED=0 GO111MODULE=on go build -installsuffix "static" -o bin/$(BIN_FILENAME) $$(find cmd/app/*.go)

clean:
	@echo Cleaning bin/ directory... && \
		rm -rfv bin/

deps:
	@echo Downloading go.mod dependencies && \
		go mod download