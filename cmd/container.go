package cmd

import (
	"github.com/ngonghi/vian-backend/config"
	"github.com/ngonghi/vian-backend/pkg/database"
	"github.com/ngonghi/vian-backend/pkg/hash"
	"github.com/ngonghi/vian-backend/pkg/logger"
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

	return container
}
