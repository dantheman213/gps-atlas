BIN_FILENAME = export
ifeq ($(OS),Windows_NT)
	BIN_FILENAME = export.exe
endif

.PHONY: all build

all: build

build:
	go build -o bin/$(BIN_FILENAME) cmd/app/main.go

clean:
	@echo Cleaning bin/ directory... && \
		rm -rfv bin/

deps:
	@echo Downloading go.mod dependencies && \
		go mod download