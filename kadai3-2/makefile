.PHONY: build
build:
	@sh -c "go build -o ./run cmd/downloader/main.go"

.PHONY: gosec
gosec:
	@sh -c "gosec ./..."

.PHONY: golint
golint:
	@sh -c "golint ./..."

.PHONY: test
test:
	@sh -c "go test -cover ./..."
