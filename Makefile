build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/requeueip cmd/requeueip.go

test:
	@echo "test"

image:
	@echo "image"

clean:
	@echo "clean"