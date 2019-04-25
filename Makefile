.PHONY: build clean


build:
	go build -o build/local-http-server main.go

clean:
	go clean
	rm -rf build/
