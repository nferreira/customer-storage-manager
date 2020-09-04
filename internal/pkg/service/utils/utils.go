package utils

import (
	"context"
	"github.com/gofiber/utils"
	"github.com/nferreira/app/pkg/app"
	"github.com/nferreira/app/pkg/service"
)

func GetCorrelationId(ctx context.Context) string {
	executionContext := getExecutionContext(ctx)
	if executionContext != nil {
		return executionContext.CorrelationId
	} else {
		return utils.UUID()
	}
}

func GetApp(ctx context.Context) app.App {
	executionContext := getExecutionContext(ctx)
	if executionContext != nil {
		return executionContext.App.(app.App)
	} else {
		panic("I need an application to live")
	}
}

func getExecutionContext(ctx context.Context) *service.ExecutionContext {
	return ctx.Value(service.ExecutionContextKey).(*service.ExecutionContext)
}
