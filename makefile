# Variables
BUILD_DIR := build
SRC_DIR := ./src/...
TST_DIR := ./test/...
ENV_FILE := .env

# Targets
.PHONY: all run-b build run clean test

all: run-b

run-b: 
	build
	./$(BUILD_DIR)/main

build: 
	check-env
	@if [ ! -d $(BUILD_DIR) ]; then 
	mkdir -p $(BUILD_DIR);
	go build -o $(BUILD_DIR)/main $(SRC_DIR);
	fi

run: 
	check-env
	go run $(SRC_DIR)

clean:
	rm -rf $(BUILD_DIR)

test:
	go test $(TST_DIR)

check-env:
	@if [ ! -f $(ENV_FILE) ]; then echo "Creating $(ENV_FILE)"; echo -e "PDF_FOLDER=./reports\nDB_PATH=./database" > $(ENV_FILE); fi
