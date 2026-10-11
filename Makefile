gen:
	@ go generate ./...

fmt:
	@ go fmt ./...

test:
	@ go test ./...

build:
	@ cd examples/all-primitives && go build
	@ cd examples/simple-house && go build
	@ cd examples/single-wall && go build
