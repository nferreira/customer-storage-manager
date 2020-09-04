# Makefile
.EXPORT_ALL_VARIABLES:	

GO111MODULE=on
GOPROXY=direct
GOSUMDB=off

run:
	@echo "########## Run customers-storage-manager"
	@go run cmd/main.go

build:
	@echo "########## Build customers-storage-manager "
	@go build -trimpath -ldflags="-s -w" -o cmd/customers-storage-manager cmd/main.go
	@echo "buid completo..."

test:
	@echo "########## Executando Tests"
	@sleep 1
	@go test -gcflags=-l github.com/nferreira/customer-storage-manager/internal/pkg/repository/mongo 

test-v:
	@echo "########## Executando Tests"
	@sleep 1
	@go test -gcflags=-l github.com/nferreira/customer-storage-manager/internal/pkg/repository/mongo -v
