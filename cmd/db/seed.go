package db

import (
	"github.com/ngonghi/vian-backend/cmd"
	"github.com/ngonghi/vian-backend/database/seeds"
	"github.com/spf13/cobra"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
)

// Seed ... seed initial data to database
func Seed(command *cobra.Command, args []string) error {
	var logger *zap.Logger
	var db *bun.DB
	container := cmd.BuildContainer()
	err := container.Invoke(func(_logger *zap.Logger, _db *bun.DB) {
		logger = _logger
		db = _db
	})
	ctx := command.Context()
	err = seeds.NewSeed(db).Execute(ctx)
	if err != nil {
		return err
	}

	logger.Info("successfully seeded")
	return nil
}
