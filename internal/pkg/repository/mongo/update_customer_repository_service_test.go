package mongo

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/nferreira/customer-storage-manager/internal/pkg/domain/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestNewUpdateCustomerRepositoryService(t *testing.T) {
	collectionx, err := getCollection()
	assert.Equal(t, err, nil)
	type args struct {
		collection *mongo.Collection
		timeout    time.Duration
	}
	tests := []struct {
		name string
		args args
		want *UpdateCustomerRepositoryService
	}{
		// TODO: Add test cases.
		{"test_update_service_1", args{collection: collectionx, timeout: 3}, &UpdateCustomerRepositoryService{collection: collectionx, timeout: 3}},
		{"test_update_service_2", args{collection: collectionx, timeout: 2}, &UpdateCustomerRepositoryService{collection: collectionx, timeout: 2}},
		{"test_update_service_3", args{collection: collectionx, timeout: 1}, &UpdateCustomerRepositoryService{collection: collectionx, timeout: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUpdateCustomerRepositoryService(tt.args.collection, tt.args.timeout); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUpdateCustomerRepositoryService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateCustomerRepositoryService_Execute(t *testing.T) {
	collectionx, err := getCollectionNew("customerDB", "customers")
	assert.Equal(t, err, nil)

	collectiony, err := getCollectionNew("customerDB_err", "customers_err")
	assert.Equal(t, err, nil)

	var customerx = &model.Customer{Id: "x39393848x838aa8383838", Name: "jeffotoni", Ssn: "12345678901232"}
	var customery = &model.Customer{Id: "12345xxxxxxx3443433433", Name: "Nadilson", Ssn: "176543211212121"}
	var ctx_ = context.Background()

	type fields struct {
		collection *mongo.Collection
		timeout    time.Duration
	}
	type args struct {
		ctx      context.Context
		customer *model.Customer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"test_update_execute_1", fields{collection: collectionx, timeout: 5}, args{ctx: ctx_, customer: customerx}, true},
		{"test_update_execute_2", fields{collection: collectiony, timeout: 6}, args{ctx: ctx_, customer: customery}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UpdateCustomerRepositoryService{
				collection: tt.fields.collection,
				timeout:    tt.fields.timeout,
			}
			if err := u.Execute(tt.args.ctx, tt.args.customer); (err != nil) != tt.wantErr {
				t.Errorf("UpdateCustomerRepositoryService.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
