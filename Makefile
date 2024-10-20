BINARY_NAME=./dist/built

build:
	go mod tidy
	go mod vendor
	go build -o ${BINARY_NAME}
	${BINARY_NAME}

refresh:
	go mod tidy
	go mod vendor

clean:
	go clean
	rm ${BINARY_NAME}