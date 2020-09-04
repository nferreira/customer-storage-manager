package mongo

import (
	"context"
	"time"

	"github.com/nferreira/customer-storage-manager/internal/pkg/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FindCustomerByIdRepositoryService struct {
	collection *mongo.Collection
	timeout    time.Duration
}

func NewFindCustomerByIdRepositoryService(collection *mongo.Collection, timeout time.Duration) *FindCustomerByIdRepositoryService {
	return &FindCustomerByIdRepositoryService{
		collection: collection,
		timeout:    timeout,
	}
}

func (c *FindCustomerByIdRepositoryService) Execute(ctx context.Context, id string) (*model.Customer, error) {
	serviceContext, _ := context.WithTimeout(ctx, c.timeout)
	customer := model.Customer{}

	result := c.collection.FindOne(serviceContext, bson.M{"_id": id})

	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return nil, nil
		} else {
			return nil, result.Err()
		}
	} else {
		if err := result.Decode(&customer); err == nil {
			return &customer, nil
		} else {
			return nil, err
		}
	}
}
