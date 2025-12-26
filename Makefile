BINARY_NAME=gopomo

install:
	go build -o ./${BINARY_NAME} main.go && cp ${BINARY_NAME} $$HOME/.local/bin

build:
	go build

run:
	go run main.go

tidy:
	go mod tidy

clean:
	go clean
