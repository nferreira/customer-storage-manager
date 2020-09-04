package update_customer

import (
	"context"
	"github.com/nferreira/app/pkg/env"
	"github.com/nferreira/app/pkg/service"
	"github.com/nferreira/customer-storage-manager/internal/pkg/domain/model"
	"github.com/nferreira/customer-storage-manager/internal/pkg/repository"
	"github.com/nferreira/customer-storage-manager/internal/pkg/service/abstract"
	serviceUtils "github.com/nferreira/customer-storage-manager/internal/pkg/service/utils"
	"time"
)

type Service struct {
	abstract.Service
}

const (
	Timeout = "UPDATE_CUSTOMER_TIMEOUT"
	ParamId = "id"
)

func New() service.BusinessService {
	return &Service{
		abstract.Service{
			Timeout: env.GetDuration(Timeout, 5 * time.Second),
		},
	}
}

func (s *Service) Name() string {
	return "UpdateCustomerService"
}

func (s *Service) CreateRequest() interface{} {
	return &model.Customer{}
}

func (s *Service) Start(_ context.Context) error {
	return nil
}

func (s *Service) Stop(_ context.Context) error {
	return nil
}

func (s *Service) CheckHealth(ctx context.Context) error {
	return nil
}
func (s *Service) Execute(ctx context.Context, params service.Params) *service.Result {
	_ = serviceUtils.GetCorrelationId(ctx)
	id := params[ParamId].(string)
	customer := params[service.Payload].(*model.Customer)
	customer.Id = id
	app := serviceUtils.GetApp(ctx)
	customerRepository := app.GetService(repository.CustomerRepositoryId).(repository.CustomerRepository)

	// TODO: handle timeout
	ctx, _ = context.WithTimeout(ctx, s.Timeout)

	if customer, err := customerRepository.UpdateCustomer(ctx, customer); err != nil {
		return service.NewResultBuilder().
			WithError(err).
			Build()
	} else {
		return service.NewResultBuilder().
			WithCode(200).
			WithHeaders(map[string]interface{}{
				"Content-Type": "application/json",
			}).
			WithResponseObject(customer).
			Build()
	}
}


