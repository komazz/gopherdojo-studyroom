.PHONY: build
build:
	@sh -c "go build -o ./run"

.PHONY: gosec
gosec:
	@sh -c "gosec ./..."

.PHONY: golint
golint:
	@sh -c "golint ./..."

.PHONY: test
test:
	@sh -c "go test -cover ./..."
