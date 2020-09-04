package find_customer_by_social_security_number

import (
	"context"
	"time"

	"github.com/nferreira/app/pkg/env"
	"github.com/nferreira/app/pkg/service"
	"github.com/nferreira/customer-storage-manager/internal/pkg/domain/model"
	"github.com/nferreira/customer-storage-manager/internal/pkg/repository"
	"github.com/nferreira/customer-storage-manager/internal/pkg/service/abstract"
	serviceUtils "github.com/nferreira/customer-storage-manager/internal/pkg/service/utils"
)

const (
	Timeout                   = "FIND_BY_SOCIAL_SECURITY_NUMBER_TIMEOUT"
	ParamSocialSecurityNumber = "social_security_number"
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
	return "FindCustomerBySocialSecurityNumber"
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

	socialSecurityNumber := params[ParamSocialSecurityNumber].(string)

	// TODO: handle timeout
	ctx, _ = context.WithTimeout(ctx, s.Timeout)

	if customer, err := customerRepository.FindBySocialSecurityNumber(ctx, socialSecurityNumber); err != nil {
		return service.NewResultBuilder().
			WithError(err).
			Build()
	} else {
		if customer != nil {
			return service.NewResultBuilder().
				WithCode(200).
				WithHeaders(map[string]interface{}{
					"Content-Type": "application/json",
				}).
				WithResponseObject(customer).
				Build()
		} else {
			return service.NewResultBuilder().
				WithCode(404).
				Build()
		}
	}
}
