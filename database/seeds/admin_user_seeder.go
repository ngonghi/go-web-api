package seeds

import (
	"context"
	"github.com/ngonghi/vian-backend/internal/models"
	"github.com/ngonghi/vian-backend/pkg/hash"
	"github.com/uptrace/bun"
)

type AdminUserSeeder struct {
	database *bun.DB
}

// Execute ...
func (seeder *AdminUserSeeder) Execute(context context.Context) error {
	hashProvider, _ := hash.NewHashProvider()
	hashedPassword, _ := hashProvider.HashPassword("test")
	model := &models.AdminUser{
		Email:    "test@example.com",
		Name:     "Test User",
		Password: hashedPassword,
		Roles:    "",
	}
	_, err := seeder.database.NewInsert().Model(model).Exec(context)
	return err
}

// NewAdminUserSeeder ... Create New Seed Instance
func NewAdminUserSeeder(db *bun.DB) *AdminUserSeeder {
	return &AdminUserSeeder{database: db}
}
