default: clean_build

.PHONY: build
build:
	go build -o monkey ./...

clean:
	rm ./monkey

clean_build: clean build

test:
	go test ./...

cover:
	go test -coverprofile=.cover ./...
	go tool cover -html=.cover

run:
	go run ./...

vet:
	go vet ./...
