package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ngonghi/vian-backend/config"
	"github.com/ngonghi/vian-backend/internal/repositories"
	"github.com/ngonghi/vian-backend/internal/services"
	"github.com/uptrace/bun"
	"go.uber.org/zap"
)

// HandlerInterface ...
type HandlerInterface interface {
	AuthSigninPost(c *gin.Context)
	AdminUsersGet(c *gin.Context)
	AdminUsersPost(c *gin.Context)
	AdminUserGet(c *gin.Context)
	AdminUserPut(c *gin.Context)
	AdminUserDelete(c *gin.Context)
}

// Handler ...
type Handler struct {
	db                  *bun.DB
	config              *config.Config
	logger              *zap.Logger
	adminUserService    services.AdminUserServiceInterface
	adminUserRepository repositories.AdminUserRepositoryInterface
}

// NewHandler ...
func NewHandler(
	db *bun.DB,
	config *config.Config,
	logger *zap.Logger,
	adminUserService services.AdminUserServiceInterface,
	adminUserRepository repositories.AdminUserRepositoryInterface,
) HandlerInterface {
	return &Handler{
		db:                  db,
		config:              config,
		logger:              logger,
		adminUserService:    adminUserService,
		adminUserRepository: adminUserRepository,
	}
}
