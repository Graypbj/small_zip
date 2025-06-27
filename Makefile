BINARY_NAME=smallzip

all: build

build:
	go build -o $(BINARY_NAME)

clean:
	rm -f $((BINARY_NAME)

run:
	./$(BINARY_NAME)

.PHONY: all build clean run

