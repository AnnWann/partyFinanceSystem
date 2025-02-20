# Variables
BUILD_DIR := build
SRC_DIR := ./src/
TST_DIR := ./test/...
ENV_FILE := .env

# Targets
.PHONY: all run-b build-and-run build run clean test check-env

all: build-and-run

build-and-run: build
	@echo 'Running...'
	./$(BUILD_DIR)/main

run-b: 
	@echo 'Running...'
	./$(BUILD_DIR)/main

build: check-env
	
	@if [ ! -d $(BUILD_DIR) ]; then \
		mkdir -p $(BUILD_DIR); \
	fi
	@echo 'Building...'
	go build -o $(BUILD_DIR)/main $(SRC_DIR)

run: check-env
	@echo 'Running...'
	go run $(SRC_DIR)

clean:
	@echo 'Cleaning...'
	rm -rf $(BUILD_DIR)

test:
	@echo 'Testing...'
	go test $(TST_DIR)

check-env:
	@if [ ! -f $(ENV_FILE) ]; then \
		echo "Creating $(ENV_FILE)"; \
		echo -e "PDF_FOLDER=./reports\nDB_PATH=./database" > $(ENV_FILE); \
	fi
