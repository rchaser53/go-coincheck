.PHONY: build clean

default: build

OUTPUT_PATH=build/coincheck

clean:
	rm -rf build && mkdir build && touch build/.gitkeep

build: clean
	go build -o $(OUTPUT_PATH)

run:
	go run index.go

goRun:
	$(OUTPUT_PATH)

nyan:
	@echo 234