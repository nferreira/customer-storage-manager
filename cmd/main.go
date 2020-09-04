package main

import (
	"context"
	"github.com/nferreira/adapter-fiber/pkg/adapter/fiber"
	"github.com/nferreira/app/pkg/bootstrap"
	"github.com/nferreira/customer-storage-manager/internal/pkg/repository"
	"github.com/nferreira/customer-storage-manager/internal/pkg/repository/mongo"
	"github.com/nferreira/customer-storage-manager/internal/pkg/service/create_customer"
	"github.com/nferreira/customer-storage-manager/internal/pkg/service/delete_customer_by_id"
	"github.com/nferreira/customer-storage-manager/internal/pkg/service/find_customer_by_id"
	"github.com/nferreira/customer-storage-manager/internal/pkg/service/find_customer_by_social_security_number"
	"github.com/nferreira/customer-storage-manager/internal/pkg/service/update_customer"
	"github.com/nferreira/logging/pkg/logging/tab_separated"
	"github.com/nferreira/logging/pkg/logging/zap"
	"golang.org/x/crypto/openpgp/errors"
)

func main() {
	ctx := context.Background()

	if app, err := bootstrap.
		NewBootstrap(ctx).
		WithAdapter(fiber.AdapterId, fiber.NewFiberAdapter()).
		WithLogger(zap.New(tab_separated.New())).
		WithService(repository.CustomerRepositoryId, mongo.NewCustomerRepository()).
		ConnectAdapterWithService(
			fiber.AdapterId,
			fiber.NewBindingRule(
				fiber.Get,
				"/api/v1/customers",
				[]string{find_customer_by_social_security_number.ParamSocialSecurityNumber},
				map[error]int{
					errors.ErrKeyIncorrect: 404,
				}),
			find_customer_by_social_security_number.New()).
		ConnectAdapterWithService(
			fiber.AdapterId,
			fiber.NewBindingRule(
				fiber.Get,
				"/api/v1/customers/:id",
				[]string{find_customer_by_id.ParamId},
				map[error]int{
					errors.ErrKeyIncorrect: 404,
				}),
			find_customer_by_id.New()).
		ConnectAdapterWithService(fiber.AdapterId,
			fiber.NewBindingRule(
				fiber.Post,
				"/api/v1/customers",
				[]string{},
				map[error]int{
					errors.ErrKeyIncorrect: 404,
					fiber.ErrBadPayload: 400,
				}),
			create_customer.New()).
		ConnectAdapterWithService(fiber.AdapterId,
			fiber.NewBindingRule(
				fiber.Put,
				"/api/v1/customers/:id",
				[]string{update_customer.ParamId},
				map[error]int{
					errors.ErrKeyIncorrect: 404,
					fiber.ErrBadPayload: 400,
				}),
			update_customer.New()).
		ConnectAdapterWithService(fiber.AdapterId,
			fiber.NewBindingRule(
				fiber.Patch,
				"/api/v1/customers/:id",
				[]string{update_customer.ParamId},
				map[error]int{
					errors.ErrKeyIncorrect: 404,
					fiber.ErrBadPayload: 400,
				}),
			update_customer.New()).
		ConnectAdapterWithService(
			fiber.AdapterId,
			fiber.NewBindingRule(
				fiber.Delete,
				"/api/v1/customers/:id",
				[]string{delete_customer_by_id.ParamId},
				map[error]int{
					errors.ErrKeyIncorrect: 404,
				}),
			delete_customer_by_id.New()).
		Boot(); err == nil {

		app.WaitForShutdown()
	}
}
