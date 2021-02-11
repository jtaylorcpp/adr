.PHONY: packr build release

packr:
	go get -u github.com/gobuffalo/packr/packr

build: packr
	$(shell go env GOPATH)/bin/packr
	mkdir -p releases/local
	go build -o releases/local/adr main.go
	$(shell go env GOPATH)/bin/packr clean

release: packr
	$(shell go env GOPATH)/bin/packr
	mkdir -p releases
	GOOS=windows go build -o releases/adr.exe main.go
	GOOS=linux go build -o releases/adr_linux main.go
	GOOS=darwin go build -o releases/adr/darwin main.go
	$(shell go env GOPATH)/bin/packr clean