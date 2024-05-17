BINARY_NAME=httpshare
BUILD_DIR=build

build: ## Build the binary file
	go build -o ./${BUILD_DIR}/${BINARY_NAME} main.go

run: ## Run the binary file
	go build -o ${BINARY_NAME} main.go
	./${BINARY_NAME}

clean: ## Remove previous build
	go clean
	rm -f ${BINARY_NAME}
	rm -Rf ./${BUILD_DIR}