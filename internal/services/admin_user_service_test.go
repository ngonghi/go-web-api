package services

import (
	"github.com/ngonghi/vian-backend/config"
	"github.com/ngonghi/vian-backend/internal/repositories"
	"github.com/ngonghi/vian-backend/pkg/database"
	"github.com/ngonghi/vian-backend/pkg/hash"
	"github.com/ngonghi/vian-backend/pkg/logger"
	"github.com/ngonghi/vian-backend/pkg/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAdminUserService(t *testing.T) {
	t.Run("Create NewAdminUserService", func(t *testing.T) {
		service := createAdminUserService()
		assert.NotNil(t, service)
	})
}

func createAdminUserService() AdminUserServiceInterface {
	configInstance, _ := config.LoadConfig()
	db, _, _ := database.GetMockDB()

	service := NewAdminUserService(
		&repositories.MockAdminUserRepository{},
		configInstance,
		db,
		logger.NewTestLogger(),
		&token.ProviderMock{},
		&hash.ProviderMock{},
	)

	return service
}
