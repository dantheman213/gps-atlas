GPSATLAS_BASE_NAME = gps-atlas

.PHONY: all build-linux build-linux-arm build-macos build-windows clean deps

build-linux:
	export GOOS=linux && \
	export GOARCH=amd64 && \
	export GPSATLAS_BIN_PATH=bin/$(GPSATLAS_BASE_NAME)-linux-amd64-$$(cat version) && \
	$(MAKE) build-target

build-linux-arm:
	export GOOS=linux && \
	export GOARCH=arm && \
	export GPSATLAS_BIN_PATH=bin/$(GPSATLAS_BASE_NAME)-linux-arm-$$(cat version) && \
	$(MAKE) build-target

build-macos:
	export GOOS=darwin && \
	export GOARCH=amd64 && \
	export GPSATLAS_BIN_PATH=bin/$(GPSATLAS_BASE_NAME)-macos-amd64-$$(cat version) && \
	$(MAKE) build-target

build-windows:
	export GOOS=windows && \
	export GOARCH=amd64 && \
	export GPSATLAS_BIN_PATH=bin/$(GPSATLAS_BASE_NAME)-windows-amd64-$$(cat version).exe && \
	$(MAKE) build-target

build-target:
	CGO_ENABLED=1 \
	GO111MODULE=on \
	go build \
	-installsuffix "static" \
	-ldflags="-X 'main.Version=$$(cat version)'" \
	-o $(GPSATLAS_BIN_PATH) \
	$$(find cmd/app/*.go)

clean:
	@echo Cleaning bin/ directory... && \
		rm -rfv bin/

deps:
	@echo Downloading go.mod dependencies && \
		go mod download
