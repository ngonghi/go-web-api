package app

import (
	"github.com/ngonghi/vian-backend/cmd"
	"github.com/ngonghi/vian-backend/config"
	"github.com/ngonghi/vian-backend/internal/http/bootstrap"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"strconv"
)

// Serve ... run the app server
func Serve(command *cobra.Command, args []string) error {
	var configInstance *config.Config
	var logger *zap.Logger

	container := cmd.BuildContainer()

	app, err := bootstrap.Bootstrap(container)
	if err != nil {
		return err
	}

	err = container.Invoke(func(_config *config.Config, _logger *zap.Logger) {
		configInstance = _config
		logger = _logger
	})
	if err != nil {
		return err
	}

	logger.Info("starting APP server ( APIs for Users )",
		zap.String("url", configInstance.App.APP.URL),
		zap.Uint("port", configInstance.App.APP.Port),
	)

	err = app.Run(":" + strconv.FormatUint(uint64(configInstance.App.APP.Port), 10))

	return err
}
