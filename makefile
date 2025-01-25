# Variables
BUILD_DIR := build
SRC_DIR := ./src/...
TST_DIR := ./test/...

# Targets
.PHONY: all build run clean test

all: build

build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/main $(SRC_DIR)

run:
	go run $(SRC_DIR)

clean:
	rm -rf $(BUILD_DIR)

test:
	go test $(TST_DIR)
