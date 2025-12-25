BINARY_NAME=gopomo

build:
	go build -o bin/${BINARY_NAME} main.go

run:
	go run main.go

tidy:
	go mod tidy

clean:
	go clean
