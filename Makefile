all: build

build:
	go build -o ./build/fling ./main.go

clean:
	rm -r ./build

.PHONY: build clean