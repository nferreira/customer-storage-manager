package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/nferreira/customer-storage-manager/internal/pkg/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UpdateCustomerRepositoryService struct {
	collection *mongo.Collection
	timeout    time.Duration
}

func NewUpdateCustomerRepositoryService(collection *mongo.Collection, timeout time.Duration) *UpdateCustomerRepositoryService {
	return &UpdateCustomerRepositoryService{
		collection: collection,
		timeout:    timeout,
	}
}

func (u *UpdateCustomerRepositoryService) Execute(ctx context.Context, customer *model.Customer) (err error) {
	//serviceContext, _ := context.WithTimeout(ctx, u.timeout)
	var res *mongo.UpdateResult

	filter := bson.M{"_id": bson.M{"$eq": customer.Id}}
	update := bson.M{"$set": customer}

	if res, err = u.collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true)); err != nil {
		return err
	}

	if res.UpsertedID != nil {
		// TODO: Just for debug
		fmt.Printf("Update info: matched_count: %d modified_count: %d upserted_count: %d\n",
			res.MatchedCount,
			res.ModifiedCount,
			res.UpsertedCount)
	}

	return nil
}
