BINARY_NAME=web-demo
CC=go

debug:
	$(CC) run .

build:
	$(CC) build -o $(BINARY_NAME) -v

run:
	./$(BINARY_NAME)