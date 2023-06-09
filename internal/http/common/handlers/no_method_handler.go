package handlers

import (
	"github.com/gin-gonic/gin"
	commonResponses "github.com/ngonghi/vian-backend/internal/http/common/responses"
	"net/http"
)

// NoMethod ... endpoint for 405
func (handler *Handler) NoMethod(ctx *gin.Context) {
	ctx.JSON(http.StatusMethodNotAllowed, commonResponses.NewNotFoundError("This method is not allowed"))
	return
}
