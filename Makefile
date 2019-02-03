


default:
	$(MAKE) build

build:
	go build

test:
	go test ./... -v

static:
	CGO_ENABLED=0 go build -x -ldflags '-w -extldflags "-static"'

