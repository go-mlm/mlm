.PHONY:	all
all:
	@echo ""
	@go fmt ./...
	@go mod tidy
	@go test -v ./...
	@go install -v ./cmd/mlm
	@echo "Compiled."
