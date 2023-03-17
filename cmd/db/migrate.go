package db

import (
	"github.com/ngonghi/vian-backend/cmd"
	"github.com/ngonghi/vian-backend/database/migrations"
	"github.com/spf13/cobra"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
	"go.uber.org/zap"
)

func initMigrator(bundb *bun.DB) *migrate.Migrator {
	return migrate.NewMigrator(bundb, migrations.Migrations, migrate.WithMarkAppliedOnSuccess(true))
}

// Migrate ... migrate database
func Migrate(command *cobra.Command, args []string) error {
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

	group, err := migrator.Migrate(ctx)
	if err != nil {
		return err
	}

	if group.IsZero() {
		logger.Info("no new migrations (database is up-to-date)")
		return nil
	}

	logger.Info("successfully migrated")
	return nil
}
