package delete_customer_by_id

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

const (
	Timeout = "DELETE_CUSTOMER_BY_ID_TIMEOUT"
	ParamId = "id"
)

type Service struct {
	abstract.Service
}

func New() service.BusinessService {

	return &Service{
		abstract.Service{
			Timeout: env.GetDuration(Timeout, 5*time.Second),
		},
	}
}

func (s *Service) CreateRequest() interface{} {
	return &model.Customer{}
}

func (s *Service) Name() string {
	return "DeleteCustomerById"
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
	app := serviceUtils.GetApp(ctx)
	customerRepository := app.GetService(repository.CustomerRepositoryId).(repository.CustomerRepository)

	id := params[ParamId].(string)

	// TODO: handle timeout
	ctx, _ = context.WithTimeout(ctx, s.Timeout)

	if err := customerRepository.DeleteById(ctx, id); err != nil {
		return service.NewResultBuilder().
			WithError(err).
			Build()
	} else {
		return service.NewResultBuilder().
			WithCode(200).
			Build()
	}
}
