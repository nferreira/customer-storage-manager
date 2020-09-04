package mongo

import (
	"context"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestDeleteByIdRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	t.Run("Delete scenarios", func(t *testing.T) {
		t.Run("Delete with invalid Mongo ObjectID format", func(t *testing.T) {
			patchNewClient(true)
			patchConnect(true)
			collection, err := getCollection()
			assert.Equal(t, err, nil)

			r := NewDeleteByIdRepositoryService(collection, 5*time.Second)
			assert.NotNil(t, r)

			err = r.Execute(context.Background(), "123")
			//assert.Equal(t, err, hex.ErrLength)
			assert.Equal(t, err, mongo.ErrClientDisconnected)
		})

		monkey.UnpatchAll()

		t.Run("Delete with with non existent id", func(t *testing.T) {
			patchNewClient(true)
			patchConnect(true)
			patchDelete(0, nil)

			collection, err := getCollection()
			assert.Equal(t, err, nil)

			r := NewDeleteByIdRepositoryService(collection, 5*time.Second)
			assert.NotNil(t, r)

			err = r.Execute(context.Background(), "5f04a55da86cd1cb1b01278c")
			assert.Equal(t, err, nil)
		})

		monkey.UnpatchAll()

		t.Run("Delete with with existent id", func(t *testing.T) {
			patchNewClient(true)
			patchConnect(true)
			patchDelete(1, nil)

			collection, err := getCollection()

			assert.Equal(t, err, nil)

			r := NewDeleteByIdRepositoryService(collection, 5*time.Second)
			assert.NotNil(t, r)

			err = r.Execute(context.Background(), "e7c91765-004f-4cc4-95fb-88956c4e11ff")
			assert.Equal(t, err, nil)
		})

		monkey.UnpatchAll()
	})
}
