package mongo

import (
	"context"
	"testing"

	"bou.ke/monkey"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestMongoDbCustomerRepository(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("Delete scenarios", func(t *testing.T) {
		t.Run("Connection failed scenarios", func(t *testing.T) {
			t.Run("Client creation failed", func(t *testing.T) {
				patchNewClient(false)
				r := NewCustomerRepository()
				err := r.Start(context.Background())
				assert.Equal(t, err, mongo.ErrClientDisconnected)
			})
			monkey.UnpatchAll()
			t.Run("Client connect failed", func(t *testing.T) {
				patchNewClient(true)
				patchConnect(false)
				r := NewCustomerRepository()
				err := r.Start(context.Background())
				assert.Equal(t, err, mongo.ErrClientDisconnected)
			})
		})

		t.Run("Actual deletion", func(t *testing.T) {
			t.Run("Delete with invalid Mongo ObjectID format", func(t *testing.T) {
				patchNewClient(true)
				patchConnect(true)
				patchPing(true)

				r := NewCustomerRepository()
				err := r.Start(context.Background())
				assert.Equal(t, err, nil)

				err = r.DeleteById(context.Background(), "123")
				//assert.Equal(t, err, hex.ErrLength)
				assert.Equal(t, err, mongo.ErrClientDisconnected)
			})

			monkey.UnpatchAll()

			t.Run("Delete with with non existent id", func(t *testing.T) {
				patchNewClient(true)
				patchConnect(true)
				patchPing(true)
				patchDelete(0, nil)

				r := NewCustomerRepository()
				err := r.Start(context.Background())
				assert.Equal(t, err, nil)

				err = r.DeleteById(context.Background(), "5f04a55da86cd1cb1b01278c")
				assert.Equal(t, err, nil)
			})

			monkey.UnpatchAll()

			t.Run("Delete with with existent id", func(t *testing.T) {
				patchNewClient(true)
				patchConnect(true)
				patchPing(true)
				patchDelete(1, nil)

				r := NewCustomerRepository()
				err := r.Start(context.Background())
				assert.Equal(t, err, nil)

				err = r.DeleteById(context.Background(), "e7c91765-004f-4cc4-95fb-88956c4e11ff")
				assert.Equal(t, err, nil)
			})
		})
	})
}
