


default:
	$(MAKE) build

build:
	go build

test:
	go test ./...

static:
	CGO_ENABLED=0 go build -x -ldflags '-w -extldflags "-static"'

