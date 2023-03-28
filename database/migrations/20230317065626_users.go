package migrations

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	ID        int64     `bun:",pk,autoincrement"`
	Name      string    `bun:",nullzero,notnull"`
	Email     string    `bun:",nullzero,notnull,unique"`
	Password  string    `bun:",nullzero,notnull"`
	CreatedAt time.Time `bun:"type:timestamp,nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:"type:timestamp,nullzero,notnull,default:current_timestamp"`
}

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		if _, err := db.
			NewCreateTable().
			Model(&User{}).
			IfNotExists().
			WithForeignKeys().
			Exec(ctx); err != nil {
			return err
		}

		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		if _, err := db.
			NewDropTable().
			Model(&User{}).
			IfExists().
			Exec(ctx); err != nil {
			return err
		}

		return nil
	})
}
