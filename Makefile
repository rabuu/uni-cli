BINARY=~/.local/bin/uni

all: install

install:
	go build -o ${BINARY} main.go
