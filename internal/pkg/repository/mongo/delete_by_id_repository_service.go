package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DeleteByIdRepositoryService struct {
	collection *mongo.Collection
	timeout    time.Duration
}

func NewDeleteByIdRepositoryService(collection *mongo.Collection, timeout time.Duration) *DeleteByIdRepositoryService {
	return &DeleteByIdRepositoryService{
		collection: collection,
		timeout:    timeout,
	}
}

func (d *DeleteByIdRepositoryService) Execute(ctx context.Context, id string) error {
	serviceContext, _ := context.WithTimeout(ctx, d.timeout)
	if _, err := d.collection.DeleteOne(serviceContext, bson.M{"_id": id}); err != nil {
		return err
	} else {
		return nil
	}
}
