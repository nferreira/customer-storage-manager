package create_customer

import (
	"context"
	"time"

	"github.com/nferreira/app/pkg/env"
	"github.com/nferreira/app/pkg/service"
	"github.com/nferreira/customer-storage-manager/internal/pkg/domain/model"
	"github.com/nferreira/customer-storage-manager/internal/pkg/repository"
	"github.com/nferreira/customer-storage-manager/internal/pkg/service/abstract"
	serviceUtils "github.com/nferreira/customer-storage-manager/internal/pkg/service/utils"
	"github.com/nferreira/logging/pkg/logging"
)

type Service struct {
	abstract.Service
}

const (
	Timeout = "CREATE_CUSTOMER_TIMEOUT"
)

func New() service.BusinessService {
	return &Service{
		abstract.Service{
			Timeout: env.GetDuration(Timeout, 25*time.Second),
		},
	}
}

func (s *Service) CreateRequest() interface{} {
	return &model.Customer{}
}

func (s *Service) Name() string {
	return "CreateCustomerService"
}

func (s *Service) Start(ctx context.Context) error {
	return nil
}

func (s *Service) Stop(_ context.Context) error {
	return nil
}

func (s *Service) CheckHealth(ctx context.Context) error {
	return nil
}

func (s *Service) Execute(ctx context.
	Context, params service.Params) *service.Result {
	correlationId := serviceUtils.GetCorrelationId(ctx)
	customer := params[service.Payload].(*model.Customer)
	app := serviceUtils.GetApp(ctx)
	logger := app.GetService(logging.LogService).(logging.Logger)
	logger.Info(correlationId, "Create Service : INVOKED")

	customerRepository := app.GetService(repository.CustomerRepositoryId).(repository.CustomerRepository)

	// TODO: handle timeout
	ctx, _ = context.WithTimeout(ctx, s.Timeout)
	if customer, err := customerRepository.CreateCustomer(ctx, customer); err != nil {
		return service.NewResultBuilder().
			WithError(err).
			Build()
	} else {
		return service.NewResultBuilder().
			WithCode(201).
			WithHeaders(map[string]interface{}{
				"Content-Type": "application/json",
			}).
			WithResponseObject(customer).
			Build()
	}
}
