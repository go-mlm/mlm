.PHONY:	all
all:
	@echo ""
	@go fmt ./...
	@go mod tidy
	@go test -v
