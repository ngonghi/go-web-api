package db

import (
	"fmt"
	"github.com/ngonghi/vian-backend/cmd"
	"github.com/spf13/cobra"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
)

func CreateMigration(command *cobra.Command, args []string) error {
	var logger *zap.Logger
	var db *bun.DB

	container := cmd.BuildContainer()
	err := container.Invoke(func(_logger *zap.Logger, _db *bun.DB) {
		logger = _logger
		db = _db
	})
	if err != nil {
		return err
	}
	migrator := initMigrator(db)
	ctx := command.Context()
	if err := migrator.Init(ctx); err != nil {
		return err
	}

	if len(args) == 0 {
		return fmt.Errorf("must input name")
	}

	if _, err := migrator.CreateGoMigration(ctx, args[0]); err != nil {
		return err
	}
	
	logger.Info("successfully create migration")
	return nil
}
