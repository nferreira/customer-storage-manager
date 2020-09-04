package mongo

import (
	"context"
	"strings"

	"github.com/gofiber/utils"
	"github.com/nferreira/customer-storage-manager/internal/pkg/domain/model"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *CustomerRepository) CreateCustomer(ctx context.Context, customer *model.Customer) (_ *model.Customer, err error) {
	serviceContext, _ := context.WithTimeout(ctx, c.writeTimeout)
	if len(strings.TrimSpace(customer.Id)) == 0 {
		customer.Id = utils.UUID()
	}
	if _, err = c.collection.InsertOne(serviceContext, &customer, options.InsertOne()); err != nil {
		return nil, err
	}

	return customer, nil
}
