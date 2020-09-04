package mongo

import (
	"context"

	"github.com/nferreira/customer-storage-manager/internal/pkg/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *CustomerRepository) FindBySocialSecurityNumber(ctx context.Context, socialSecurityNumber string) (*model.Customer, error) {
	serviceContext, _ := context.WithTimeout(ctx, c.readTimeout)

	findOptions := options.Find()
	findOptions.SetLimit(1)
	filter := make(bson.D, 0)
	filter = append(filter, bson.E{Key: "ssn", Value: socialSecurityNumber})

	customer := model.Customer{}

	if cur, err := c.collection.Find(serviceContext, filter, findOptions); err != nil {
		return nil, err
	} else {
		defer cur.Close(ctx)
		if cur.Next(ctx) {
			err = cur.Decode(&customer)
			return &customer, nil
		} else {
			return nil, nil
		}
	}
}
