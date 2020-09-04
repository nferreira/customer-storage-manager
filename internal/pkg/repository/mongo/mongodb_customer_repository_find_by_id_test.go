package mongo

// import (
// 	"context"
// 	"reflect"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/assert"

// 	"github.com/nferreira/customer-storage-manager/internal/pkg/domain/model"
// 	"github.com/nferreira/customer-storage-manager/internal/pkg/repository/repository_service"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"gopkg.in/go-playground/assert.v1"
// )

// func TestCustomerRepository_FindById(t *testing.T) {

// 	// aqui ira gerar temporariamente
// 	// os dados para fazer o find
// 	// e os proprio find remote
// 	tmpCustomer := pathCreateCustomer()

// 	collectionx, clientx, err := getCollectionClientNew()
// 	assert.Equal(t, err, nil)

// 	updateCustomerRepositoryService := NewUpdateCustomerRepositoryService(collectionx, 15)
// 	deleteByIdRepositoryService := NewDeleteByIdRepositoryService(collectionx, 15)

// 	var ctx_ = context.Background()
// 	var customery = &model.Customer{
// 		Id:   tmpCustomer.Id,
// 		Name: tmpCustomer.Name, Ssn: tmpCustomer.Ssn,
// 	}

// 	type fields struct {
// 		client                          *mongo.Client
// 		pingTimeout                     time.Duration
// 		readTimeout                     time.Duration
// 		writeTimeout                    time.Duration
// 		collection                      *mongo.Collection
// 		updateCustomerRepositoryService repository_service.UpdateCustomerRepositoryService
// 		deleteByIdRepositoryService     repository_service.DeleteByIdRepositoryService
// 	}
// 	type args struct {
// 		ctx context.Context
// 		id  string
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		want    *model.Customer
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			"test_createcustomer_1",
// 			fields{
// 				client:                          clientx,
// 				pingTimeout:                     10,
// 				readTimeout:                     11,
// 				writeTimeout:                    12,
// 				collection:                      collectionx,
// 				updateCustomerRepositoryService: updateCustomerRepositoryService,
// 				deleteByIdRepositoryService:     deleteByIdRepositoryService,
// 			},
// 			args{ctx: ctx_, id: tmpCustomer.Id},
// 			customery,
// 			false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			c := &CustomerRepository{
// 				client:                          tt.fields.client,
// 				pingTimeout:                     tt.fields.pingTimeout,
// 				readTimeout:                     tt.fields.readTimeout,
// 				writeTimeout:                    tt.fields.writeTimeout,
// 				collection:                      tt.fields.collection,
// 				updateCustomerRepositoryService: tt.fields.updateCustomerRepositoryService,
// 				deleteByIdRepositoryService:     tt.fields.deleteByIdRepositoryService,
// 			}
// 			got, err := c.FindById(tt.args.ctx, tt.args.id)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("CustomerRepository.FindById() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("CustomerRepository.FindById() = %v, want %v", got, tt.want)
// 			}
// 			r := NewDeleteByIdRepositoryService(tt.fields.collection, 5*time.Second)
// 			err = r.Execute(context.Background(), tmpCustomer.Id)
// 			assert.Equal(t, err, nil)
// 		})
// 	}
// }
