package mongo

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/gofiber/utils"
	"github.com/nferreira/customer-storage-manager/internal/pkg/domain/model"
	"github.com/nferreira/customer-storage-manager/internal/pkg/repository/repository_service"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/go-playground/assert.v1"
)

func TestCustomerRepository_CreateCustomer(t *testing.T) {
	collectionx, clientx, err := getCollectionClientNew()
	assert.Equal(t, err, nil)

	updateCustomerRepositoryService := NewUpdateCustomerRepositoryService(collectionx, 15)
	deleteByIdRepositoryService := NewDeleteByIdRepositoryService(collectionx, 15)

	var ctx_ = context.Background()
	Id := utils.UUID()
	var socialSecurityNumber string = "234.567.666-77"
	var customery = &model.Customer{Id: Id, Name: "Nadilson", Ssn: socialSecurityNumber}

	type fields struct {
		client                          *mongo.Client
		pingTimeout                     time.Duration
		readTimeout                     time.Duration
		writeTimeout                    time.Duration
		collection                      *mongo.Collection
		updateCustomerRepositoryService repository_service.UpdateCustomerRepositoryService
		deleteByIdRepositoryService     repository_service.DeleteByIdRepositoryService
	}
	type args struct {
		ctx      context.Context
		customer *model.Customer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Customer
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"test_createcustomer_1",
			fields{
				client:                          clientx,
				pingTimeout:                     10,
				readTimeout:                     11,
				writeTimeout:                    12,
				collection:                      collectionx,
				updateCustomerRepositoryService: updateCustomerRepositoryService,
				deleteByIdRepositoryService:     deleteByIdRepositoryService,
			},
			args{ctx: ctx_, customer: customery},
			customery,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CustomerRepository{
				client:                          tt.fields.client,
				pingTimeout:                     tt.fields.pingTimeout,
				readTimeout:                     tt.fields.readTimeout,
				writeTimeout:                    tt.fields.writeTimeout,
				collection:                      tt.fields.collection,
				updateCustomerRepositoryService: tt.fields.updateCustomerRepositoryService,
				deleteByIdRepositoryService:     tt.fields.deleteByIdRepositoryService,
			}
			got, err := c.CreateCustomer(tt.args.ctx, tt.args.customer)
			if (err != nil) != tt.wantErr {
				t.Errorf("CustomerRepository.CreateCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CustomerRepository.CreateCustomer() = %v, want %v", got, tt.want)
			}

			r := NewDeleteByIdRepositoryService(tt.fields.collection, 5*time.Second)
			//assert.NotNil(t, r)
			err = r.Execute(context.Background(), Id)
			assert.Equal(t, err, nil)

		})
	}
}
