BINARY_NAME=2048
GO=go
BUILD_DIR=build
DIST_DIR=dist

.PHONY: all run build clean dist

all: run

run:
	$(GO) run -tags=audio .

build:
	mkdir -p $(BUILD_DIR)
	$(GO) build -tags=audio -o $(BUILD_DIR)/$(BINARY_NAME) .

clean:
	rm -rf $(BUILD_DIR)
	rm -rf $(DIST_DIR)

dist: clean
	mkdir -p $(DIST_DIR)
	# Mac (x86_64)
	GOOS=darwin GOARCH=amd64 $(GO) build -tags=audio -o $(DIST_DIR)/$(BINARY_NAME)-mac-x86_64 .
	# Mac (M1/ARM64)
	GOOS=darwin GOARCH=arm64 $(GO) build -tags=audio -o $(DIST_DIR)/$(BINARY_NAME)-mac-arm64 .
	# Linux (x86)
	GOOS=linux GOARCH=386 $(GO) build -tags=audio -o $(DIST_DIR)/$(BINARY_NAME)-linux-x86 .
	# Linux (ARM)
	GOOS=linux GOARCH=arm $(GO) build -tags=audio -o $(DIST_DIR)/$(BINARY_NAME)-linux-arm .
	# Linux (x64)
	GOOS=linux GOARCH=amd64 $(GO) build -tags=audio -o $(DIST_DIR)/$(BINARY_NAME)-linux-x64 .
	# Windows (64-bit)
	GOOS=windows GOARCH=amd64 $(GO) build -tags=audio -o $(DIST_DIR)/$(BINARY_NAME)-windows-64.exe .