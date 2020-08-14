.PHONY: build
build:
	@sh -c "go build -o ./exec"

.PHONY: gosec
gosec:
	@sh -c "gosec ./..."

.PHONY: golint
golint:
	@sh -c "golint ./..."

.PHONY: test
test:
	@sh -c "./testdata.zsh"
	@sh -c "go test -cover ./..."
