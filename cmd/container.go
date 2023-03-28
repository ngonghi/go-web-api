package cmd

import (
	"github.com/ngonghi/vian-backend/config"
	adminHandlers "github.com/ngonghi/vian-backend/internal/http/admin/handlers"
	adminMiddlewares "github.com/ngonghi/vian-backend/internal/http/admin/middlewares"
	"github.com/ngonghi/vian-backend/internal/http/common/handlers"
	"github.com/ngonghi/vian-backend/internal/repositories"
	"github.com/ngonghi/vian-backend/internal/services"
	"github.com/ngonghi/vian-backend/pkg/database"
	"github.com/ngonghi/vian-backend/pkg/hash"
	"github.com/ngonghi/vian-backend/pkg/logger"
	"github.com/ngonghi/vian-backend/pkg/middlewares"
	"github.com/ngonghi/vian-backend/pkg/token"
	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	container := dig.New()
	_ = container.Provide(config.LoadConfig)
	_ = container.Provide(logger.NewLogger)
	_ = container.Provide(database.InitDatabase)

	// Packages
	_ = container.Provide(
		func(config *config.Config) (token.ProviderInterface, error) {
			return token.NewTokenProvider([]byte(config.Authentication.Secret))
		},
	)
	_ = container.Provide(hash.NewHashProvider)

	// Middlewares
	_ = container.Provide(middlewares.Logger, dig.Name("logger"))
	_ = container.Provide(middlewares.RequestID, dig.Name("requestID"))
	_ = container.Provide(middlewares.SecurityHeaders, dig.Name("securityHeaders"))
	_ = container.Provide(adminMiddlewares.Auth, dig.Name("adminAuth"))

	// Repositories
	_ = container.Provide(repositories.NewAdminUserRepository)

	// Services
	_ = container.Provide(services.NewAdminUserService)

	// Handler
	_ = container.Provide(handlers.NewHandler)
	_ = container.Provide(adminHandlers.NewHandler)

	return container
}
