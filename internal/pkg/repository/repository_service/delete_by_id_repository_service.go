package repository_service

import (
	"context"
)

type DeleteByIdRepositoryService interface {
	Execute(ctx context.Context, id string) error
}
