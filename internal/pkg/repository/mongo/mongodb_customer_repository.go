package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/nferreira/app/pkg/env"
	"github.com/nferreira/customer-storage-manager/internal/pkg/domain/model"
	"github.com/nferreira/customer-storage-manager/internal/pkg/repository"
	"github.com/nferreira/customer-storage-manager/internal/pkg/repository/repository_service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type CustomerRepository struct {
	client       *mongo.Client
	pingTimeout  time.Duration
	readTimeout  time.Duration
	writeTimeout time.Duration
	collection   *mongo.Collection

	findCustomerByIdRepositoryService repository_service.FindCustomerByIdRepositoryService
	updateCustomerRepositoryService   repository_service.UpdateCustomerRepositoryService
	deleteByIdRepositoryService       repository_service.DeleteByIdRepositoryService
}

const (
	Schema            = "MONGODB_SCHEMA"
	Uri               = "MONGODB_URI"
	Username          = "MONGODB_USERNAME"
	Password          = "MONGODB_PASSWORD"
	DatabaseName      = "MONGODB_DATABASE"
	CollectionName    = "MONGODB_COLLECTION"
	PingTimeout       = "MONGODB_PING_TIMEOUT"
	ReadTimeout       = "MONGODB_READ_TIMEOUT"
	WriteTimeout      = "MONGODB_WRITE_TIMEOUT"
	Options           = "MONGODB_OPTIONS"
	ConnectionTimeout = "MONGODB_CONNECTION_TIMEOUT_IN_SECONDS"
)

func NewCustomerRepository() repository.CustomerRepository {
	return &CustomerRepository{}
}

func (c *CustomerRepository) Start(ctx context.Context) (err error) {
	schema := env.GetString(Schema, "mongodb")
	uri := env.GetString(Uri, "localhost:27017")
	user := env.GetString(Username, "root")
	password := env.GetString(Password, "abc123")
	databaseName := env.GetString(DatabaseName, "customerDB")
	collectionName := env.GetString(CollectionName, "customers")

	c.pingTimeout = env.GetDuration(PingTimeout, 3*time.Second)
	c.readTimeout = env.GetDuration(ReadTimeout, 5*time.Second)
	c.writeTimeout = env.GetDuration(WriteTimeout, 5*time.Second)
	clientOptions := env.GetString(Options, "authSource=admin&readPreference=primary&appname=MongoDB%20Compass&ssl=false")

	connectionTimeout := env.GetDuration(ConnectionTimeout, 10*time.Second)
	connectionString := fmt.Sprintf("%s://%s:%s@%s/%s?%s",
		schema, user, password, uri, databaseName, clientOptions)
	if len(user) == 0 && len(password) == 0 {
		connectionString = fmt.Sprintf("%s://%s/%s?%s",
			schema, uri, databaseName, clientOptions)
	}

	// TODO: Add log ONLY with the pubic info *** PLEASE DO NOT ADD THE PASSWORD TO THE LOG
	if c.client, err = mongo.NewClient(options.Client().ApplyURI(connectionString)); err != nil {
		// TODO: add logOne
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, connectionTimeout*time.Second)
	defer cancel()

	if err = c.client.Connect(ctx); err != nil {
		return err
	}

	c.collection = c.client.Database(databaseName).Collection(collectionName)
	c.findCustomerByIdRepositoryService = NewFindCustomerByIdRepositoryService(c.collection, c.readTimeout)
	c.updateCustomerRepositoryService = NewUpdateCustomerRepositoryService(c.collection, c.writeTimeout)
	c.deleteByIdRepositoryService = NewDeleteByIdRepositoryService(c.collection, c.writeTimeout)

	return err
}

func (c *CustomerRepository) Stop(ctx context.Context) error {
	if err := c.client.Disconnect(ctx); err != nil {
		return err
	}
	return nil
}

func (c *CustomerRepository) CheckHealth(ctx context.Context) error {
	ctx, _ = context.WithTimeout(ctx, c.pingTimeout)
	return c.client.Ping(ctx, readpref.Primary())
}

func (c *CustomerRepository) FindById(ctx context.Context, id string) (*model.Customer, error) {
	return c.findCustomerByIdRepositoryService.Execute(ctx, id)
}

func (c *CustomerRepository) UpdateCustomer(ctx context.Context, customer *model.Customer) (_ *model.Customer, err error) {
	err = c.updateCustomerRepositoryService.Execute(ctx, customer)
	return customer, err
}

func (c *CustomerRepository) DeleteById(ctx context.Context, id string) error {
	return c.deleteByIdRepositoryService.Execute(ctx, id)
}
