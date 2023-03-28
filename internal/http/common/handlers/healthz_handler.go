package handlers

import (
	commonResponses "github.com/ngonghi/vian-backend/internal/http/common/responses"
	"github.com/ngonghi/vian-backend/pkg/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Healthz ... endpoint for checking health
func (handler *Handler) Healthz(ctx *gin.Context) {
	if database.PingDB(ctx, handler.db, handler.logger) {
		ctx.JSON(http.StatusOK, commonResponses.NewSuccessStatus())
		return
	}

	ctx.JSON(http.StatusServiceUnavailable, commonResponses.NewServiceUnavailableErrorStatus())
}
