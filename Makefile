BINARY_NAME=gopomo

build:
	go build -o ./${BINARY_NAME} main.go && cp ${BINARY_NAME} $HOME/.local/bin

run:
	go run main.go

tidy:
	go mod tidy

clean:
	go clean
