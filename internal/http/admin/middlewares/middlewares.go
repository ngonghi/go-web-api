package middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// Middlewares ... middleware structure for DI
type Middlewares struct {
	dig.In
	Auth gin.HandlerFunc `name:"adminAuth"`
}
