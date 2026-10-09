gen:
	@ go generate ./...

fmt:
	@ go fmt ./...

test:
	@ go test ./...

build:
	@ cd examples/single-wall && go build
