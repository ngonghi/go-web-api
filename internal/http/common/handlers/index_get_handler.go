package handlers

import (
	commonResponses "github.com/ngonghi/vian-backend/internal/http/common/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

// IndexGet ... endpoint for checking health
func (handler *Handler) IndexGet(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, commonResponses.NewSuccessStatus())
}
