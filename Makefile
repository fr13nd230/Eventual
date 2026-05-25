
run:
	@go run cmd/*.go

build:
	@go build -o build/eventual cmd/*.go

test:
	@go test -cover -v ./...

install:
	@go install cmd/*go
