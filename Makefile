fmt:
	@ go fmt ./...

test:
	@ go test ./...

build:
	@ cd examples/single_wall && go build
