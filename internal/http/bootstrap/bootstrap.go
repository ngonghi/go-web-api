package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/ngonghi/vian-backend/internal/http/admin"
	"github.com/ngonghi/vian-backend/internal/http/common"
	"go.uber.org/dig"
	"os"
)

func Bootstrap(container *dig.Container) (*gin.Engine, error) {
	var err error

	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	gin.DefaultWriter = os.Stderr

	_app := gin.New()

	err = setupMiddlewares(_app, container)
	if err != nil {
		panic(err)
	}
	err = common.SetupRoutes(_app, container)
	if err != nil {
		panic(err)
	}
	err = admin.SetupRoutes(_app, container)

	return _app, err
}
