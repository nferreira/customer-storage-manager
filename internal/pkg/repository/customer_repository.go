package repository

import (
	"context"
	"github.com/nferreira/app/pkg/service"
	"github.com/nferreira/customer-storage-manager/internal/pkg/domain/model"
)

const (
	CustomerRepositoryId = "CustomerRepository"
)

type CustomerRepository interface {
	service.Service
	FindById(ctx context.Context, id string) (*model.Customer, error)
	FindBySocialSecurityNumber(ctx context.Context, socialSecurityNumber string) (*model.Customer, error)
	CreateCustomer(ctx context.Context, customer *model.Customer) (*model.Customer, error)
	UpdateCustomer(ctx context.Context, customer *model.Customer) (_ *model.Customer, err error)
	DeleteById(ctx context.Context, id string) error
}
