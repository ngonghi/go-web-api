package services

import (
	"context"
	"errors"
	"github.com/ngonghi/vian-backend/config"
	"github.com/ngonghi/vian-backend/internal/models"
	"github.com/ngonghi/vian-backend/internal/repositories"
	"github.com/ngonghi/vian-backend/pkg/hash"
	"github.com/ngonghi/vian-backend/pkg/token"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
	"strconv"
)

// AdminUserServiceInterface ...
type AdminUserServiceInterface interface {
	GetTokenFromEmailAndPassword(ctx context.Context, email string, password string) (*string, *models.AdminUser, error)
	GetAdminUserFromToken(ctx context.Context, token string) (*models.AdminUser, error)
	CreateAdminUser(ctx context.Context, name string, email string, password string) (*models.AdminUser, error)
	UpdateAdminUser(ctx context.Context, id int64, name *string, email *string, password *string) (*models.AdminUser, error)
	DeleteAdminUser(ctx context.Context, id int64) error
}

// AdminUserService ... provides admin user related features
type AdminUserService struct {
	adminUserRepository repositories.AdminUserRepositoryInterface
	config              *config.Config
	database            *bun.DB
	logger              *zap.Logger
	tokenProvider       token.ProviderInterface
	hashProvider        hash.ProviderInterface
}

type UserInfo struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

// NewAdminUserService ... creates a new instance of AdminUserService
func NewAdminUserService(
	adminUserRepository repositories.AdminUserRepositoryInterface,
	config *config.Config,
	database *bun.DB,
	logger *zap.Logger,
	tokenProvider token.ProviderInterface,
	hashProvider hash.ProviderInterface,
) AdminUserServiceInterface {
	return &AdminUserService{
		adminUserRepository: adminUserRepository,
		config:              config,
		database:            database,
		logger:              logger,
		tokenProvider:       tokenProvider,
		hashProvider:        hashProvider,
	}
}

func (service *AdminUserService) GetTokenFromEmailAndPassword(ctx context.Context, email string, password string) (*string, *models.AdminUser, error) {
	adminUser, err := service.adminUserRepository.FindByEmail(ctx, email)
	if err != nil {
		return nil, nil, err
	}
	if adminUser == nil {
		return nil, nil, errors.New("admin user not found")
	}
	result := service.hashProvider.CheckPasswordHash(password, adminUser.Password)
	if result {
		_token, err := service.tokenProvider.GenerateToken(strconv.FormatInt(adminUser.ID, 10), UserInfo{
			ID:    strconv.FormatInt(adminUser.ID, 10),
			Email: adminUser.Email,
		})
		if err != nil {
			return nil, nil, err
		}
		return &_token, adminUser, nil
	}
	return nil, nil, errors.New("invalid password")
}

func (service *AdminUserService) GetAdminUserFromToken(ctx context.Context, token string) (*models.AdminUser, error) {
	claims, err := service.tokenProvider.ValidateToken(token)
	val, ok := claims["UserInfo"]
	if !ok {
		return nil, errors.New("invalid token")
	}
	id := (val.(map[string]interface{}))["id"]
	_uuid, err := strconv.ParseInt(id.(string), 10, 64)
	if err != nil {
		return nil, err
	}
	adminUser, err := service.adminUserRepository.FindByID(ctx, _uuid)
	if err != nil {
		return nil, err
	}
	if adminUser == nil {
		return nil, errors.New("invalid token")
	}
	return adminUser, nil
}

func (service *AdminUserService) CreateAdminUser(ctx context.Context, name string, email string, password string) (*models.AdminUser, error) {
	adminUser, err := service.adminUserRepository.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if adminUser != nil {
		return nil, errors.New("admin user already exists")
	}
	hashedPassword, err := service.hashProvider.HashPassword(password)
	if err != nil {
		return nil, err
	}
	adminUser = &models.AdminUser{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
		Roles:    "",
	}
	adminUser, err = service.adminUserRepository.Save(ctx, *adminUser)
	if err != nil {
		return nil, err
	}
	return adminUser, nil
}

func (service *AdminUserService) UpdateAdminUser(ctx context.Context, id int64, name *string, email *string, password *string) (*models.AdminUser, error) {
	adminUser, err := service.adminUserRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if adminUser == nil {
		return nil, errors.New("admin user not found")
	}
	if name != nil {
		adminUser.Name = *name
	}
	if email != nil {
		adminUser.Email = *email
	}
	if password != nil {
		hashedPassword, err := service.hashProvider.HashPassword(*password)
		if err != nil {
			return nil, err
		}
		adminUser.Password = hashedPassword
	}
	adminUser, err = service.adminUserRepository.Save(ctx, *adminUser)
	if err != nil {
		return nil, err
	}
	return adminUser, nil
}

func (service *AdminUserService) DeleteAdminUser(ctx context.Context, id int64) error {
	adminUser, err := service.adminUserRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if adminUser == nil {
		return errors.New("admin user not found")
	}
	return service.adminUserRepository.Delete(ctx, adminUser.ID)
}
