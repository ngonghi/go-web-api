package migrations

import (
	"context"
	"github.com/uptrace/bun"
	"time"
)

type AdminUser struct {
	ID        int64     `bun:",pk,autoincrement"`
	Name      string    `bun:",nullzero,notnull"`
	Email     string    `bun:",nullzero,notnull,unique"`
	Password  string    `bun:",nullzero,notnull"`
	Roles     string    `bun:"type:text,nullzero,notnull"`
	CreatedAt time.Time `bun:"type:timestamp,nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"type:timestamp,nullzero,notnull,default:current_timestamp"`
}

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		if _, err := db.
			NewCreateTable().
			Model(&AdminUser{}).
			IfNotExists().
			WithForeignKeys().
			Exec(ctx); err != nil {
			return err
		}

		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		if _, err := db.
			NewDropTable().
			Model(&AdminUser{}).
			IfExists().
			Exec(ctx); err != nil {
			return err
		}

		return nil
	})
}
