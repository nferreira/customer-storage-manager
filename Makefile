# Makefile
.EXPORT_ALL_VARIABLES:	

GO111MODULE=on
GOPROXY=direct
GOSUMDB=off

run:
	@echo "########## Running..."
	@go run cmd/main.go

build:
	@echo "########## Building..."
	@go build -trimpath -ldflags="-s -w" -o cmd/customers-storage-manager cmd/main.go
	@echo "buid completo..."

test:
	@echo "########## Running Tests"
	@sleep 1
	@go test -gcflags=-l github.com/nferreira/customer-storage-manager/internal/pkg/repository/mongo 

test-v:
	@echo "########## Runing Tests"
	@sleep 1
	@go test -gcflags=-l github.com/nferreira/customer-storage-manager/internal/pkg/repository/mongo -v
