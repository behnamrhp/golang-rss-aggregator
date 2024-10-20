BINARY_NAME=./dist/built

build:
	go mod tidy
	go build -o ${BINARY_NAME} main.go
	${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}