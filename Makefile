GPSATLAS_BASE_NAME = gps-atlas

.PHONY: all build build-linux build-linux-arm build-macos build-windows clean deps

all: build-linux build-linux-arm build-macos build-windows

build:
	CGO_ENABLED=0 \
	GO111MODULE=on \
	go build \
	-installsuffix "static" \
	-ldflags="-X 'main.Version=$$(cat version)'" \
	-o $(GPSATLAS_BIN_PATH) \
	$$(find cmd/app/*.go)

build-linux:
	export GOOS=linux && \
	export GOARCH=amd64 && \
	export GPSATLAS_BIN_PATH=bin/linux-x86/$(GPSATLAS_BASE_NAME) && \
	$(MAKE) build

build-linux-arm:
	export GOOS=linux && \
	export GOARCH=arm && \
	export GPSATLAS_BIN_PATH=bin/linux-arm/$(GPSATLAS_BASE_NAME) && \
	$(MAKE) build

build-macos:
	export GOOS=darwin && \
	export GOARCH=amd64 && \
	export GPSATLAS_BIN_PATH=bin/macos/$(GPSATLAS_BASE_NAME) && \
	$(MAKE) build

build-windows:
	export GOOS=windows && \
	export GOARCH=amd64 && \
	export GPSATLAS_BIN_PATH=bin/windows/$(GPSATLAS_BASE_NAME).exe && \
	$(MAKE) build

clean:
	@echo Cleaning bin/ directory... && \
		rm -rfv bin/

deps:
	@echo Downloading go.mod dependencies && \
		go mod download
