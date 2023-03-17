package models

import (
	"github.com/bxcodec/faker/v4"
	"github.com/uptrace/bun"
	"time"
)

type AdminUser struct {
	bun.BaseModel `bun:"table:admin_users,alias:admin_users"`
	ID            int64     `bun:",pk,autoincrement"`
	Name          string    `bun:",nullzero,notnull" faker:"name"`
	Email         string    `bun:",nullzero,notnull,unique" faker:"email"`
	Password      string    `bun:",nullzero,notnull"`
	Roles         string    `bun:"type:text,nullzero,notnull"`
	CreatedAt     time.Time `bun:"type:timestamp,nullzero,notnull,default:current_timestamp"`
	UpdatedAt     time.Time `bun:"type:timestamp,nullzero,notnull,default:current_timestamp"`
}

// GetFakeAdminUser ... get fake AdminUser model
func GetFakeAdminUser() (*AdminUser, error) {
	entity := &AdminUser{}
	err := faker.FakeData(entity)

	if err != nil {
		return nil, err
	}

	entity.ID = 1
	entity.CreatedAt = time.Now()
	entity.UpdatedAt = time.Now()
	entity.Roles = ""

	return entity, nil
}
