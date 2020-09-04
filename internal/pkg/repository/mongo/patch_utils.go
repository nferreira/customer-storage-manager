package mongo

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"bou.ke/monkey"
	"github.com/gofiber/utils"
	"github.com/nferreira/app/pkg/env"
	"github.com/nferreira/customer-storage-manager/internal/pkg/domain/model"
	"github.com/nferreira/customer-storage-manager/internal/pkg/repository/repository_service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	schema           = env.GetString(Schema, "mongodb")
	uri              = env.GetString(Uri, "localhost:27017")
	user             = env.GetString(Username, "root")
	password         = env.GetString(Password, "abc123")
	databaseName     = env.GetString(DatabaseName, "customerDB")
	collectionName   = env.GetString(CollectionName, "customers")
	clientOptions    = env.GetString(Options, "authSource=admin&readPreference=primary&appname=MongoDB%20Compass&ssl=false")
	connectionString = ""
)

func init() {
	if len(user) == 0 && len(password) == 0 {
		connectionString = fmt.Sprintf("%s://%s/%s?%s",
			schema, uri, databaseName, clientOptions)
	} else {
		connectionString = fmt.Sprintf("%s://%s:%s@%s/%s?%s",
			schema, user, password, uri, databaseName, clientOptions)
	}
}

func getCollectionNew(database, collection string) (*mongo.Collection, error) {
	if client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017")); err != nil {
		return nil, err
	} else {
		if err = client.Connect(context.Background()); err != nil {
			return nil, err
		}

		return client.Database(database).Collection(collection), nil
	}
}

func getCollectionClientNew() (*mongo.Collection, *mongo.Client, error) {
	if client, err := mongo.NewClient(options.Client().ApplyURI(connectionString)); err != nil {
		return nil, nil, err
	} else {
		if err = client.Connect(context.TODO()); err != nil {
			return nil, nil, err
		}
		return client.Database(databaseName).Collection(collectionName), client, nil
	}
}

func getCollection() (*mongo.Collection, error) {
	if client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017")); err != nil {
		return nil, err
	} else {
		if err = client.Connect(context.Background()); err != nil {
			return nil, err
		}

		return client.Database("someDB").Collection("anyCollection"), nil
	}
}

func patchDelete(deleteCount int64, err error) {
	var collection *mongo.Collection
	monkey.PatchInstanceMethod(reflect.TypeOf(collection), "DeleteOne",
		func(collection *mongo.Collection,
			context context.Context,
			filter interface{},
			opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
			if err != nil {
				return nil, err
			}
			return &mongo.DeleteResult{
				DeletedCount: deleteCount,
			}, nil
		})
}

func patchFind() {
	var collection *mongo.Collection
	monkey.PatchInstanceMethod(reflect.TypeOf(collection), "Find", func(_ *mongo.Collection, _ context.Context, _ interface{}, _ ...*options.FindOptions) (*mongo.Cursor, error) {
		return &mongo.Cursor{}, nil
	})
}

func patchCursorNext() {
	var cursor *mongo.Cursor
	monkey.PatchInstanceMethod(reflect.TypeOf(cursor), "Next", func(_ *mongo.Cursor, _ context.Context) bool {
		return false
	})
}

func patchCursorClose() {
	var cursor *mongo.Cursor
	monkey.PatchInstanceMethod(reflect.TypeOf(cursor), "Close", func(_ *mongo.Cursor, _ context.Context) error {
		return nil
	})
}

func patchDisconnect(disconnectOk bool) {
	var c *mongo.Client
	monkey.PatchInstanceMethod(reflect.TypeOf(c), "Disconnect", func(_ *mongo.Client, _ context.Context) error {
		if disconnectOk {
			return nil
		}
		return mongo.ErrClientDisconnected
	})
}

func patchPing(pingOk bool) {
	var c *mongo.Client
	monkey.PatchInstanceMethod(reflect.TypeOf(c), "Ping", func(_ *mongo.Client, _ context.Context, _ *readpref.ReadPref) error {
		if pingOk {
			return nil
		}
		return mongo.ErrNilDocument
	})
}

func patchConnect(connectOk bool) {
	var c *mongo.Client
	monkey.PatchInstanceMethod(reflect.TypeOf(c), "Connect", func(_ *mongo.Client, _ context.Context) error {
		if connectOk {
			return nil
		}
		return mongo.ErrClientDisconnected
	})
}

func patchNewClient(clientOk bool) {
	if !clientOk {
		monkey.Patch(mongo.NewClient, func(...*options.ClientOptions) (*mongo.Client, error) {
			return nil, mongo.ErrClientDisconnected
		})
	}
}

type TmpCustomer struct {
	Id   string
	Name string
	Ssn  string
}

func pathCreateCustomer() (tmpCustomer *TmpCustomer) {
	collectionx, clientx, err := getCollectionClientNew()
	if err != nil {
		log.Println("Error getCollectionClientNew in path:", err)
		return
	}
	updateCustomerRepositoryService := NewUpdateCustomerRepositoryService(collectionx, 15)
	deleteByIdRepositoryService := NewDeleteByIdRepositoryService(collectionx, 15)

	var ctx_ = context.Background()
	Id := utils.UUID()
	var customery = &model.Customer{Id: Id, Name: "Nadilson"}
	tmpCustomer = &TmpCustomer{Id: Id, Name: "Nadilson"}

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
			fmt.Errorf("CustomerRepository.CreateCustomer() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(got, tt.want) {
			fmt.Errorf("CustomerRepository.CreateCustomer() = %v, want %v", got, tt.want)
		}
	}
	return
}
