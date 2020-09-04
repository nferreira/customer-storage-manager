package repository_service

import (
	"context"
	"github.com/nferreira/customer-storage-manager/internal/pkg/domain/model"
)

type UpdateCustomerRepositoryService interface {
	Execute(ctx context.Context, customer *model.Customer) error
}
