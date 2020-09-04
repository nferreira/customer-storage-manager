package repository_service

import (
	"context"
	"github.com/nferreira/customer-storage-manager/internal/pkg/domain/model"
)

type FindCustomerByIdRepositoryService interface {
	Execute(ctx context.Context, id string) (*model.Customer, error)
}
