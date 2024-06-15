
test:
	go test -v ./...

deps:
	go mod tidy
	go mod vendor
