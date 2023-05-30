package router

import (
	"net/http"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.atatus.com/agent/module/atgin"

	"github.com/lordscoba/bible_compass_backend/pkg/middleware"
	"github.com/lordscoba/bible_compass_backend/utility"
)

func Setup(validate *validator.Validate, logger *utility.Logger) *gin.Engine {

	r := gin.New()
	r.Use(atgin.Middleware(r))
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CORS())
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	r.MaxMultipartMemory = 1 << 20 // 1MB

	ApiVersion := "v1"
	Health(r, validate, ApiVersion, logger)
	Admin(r, validate, ApiVersion, logger)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"name":    "Not Found",
			"message": "Page not found.",
			"code":    404,
			"status":  http.StatusNotFound,
		})
	})

	return r
}
