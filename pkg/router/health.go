package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lordscoba/bible_compass_backend/pkg/handler/health"
	"github.com/lordscoba/bible_compass_backend/utility"
)

func Health(r *gin.Engine, validate *validator.Validate, ApiVersion string, logger *utility.Logger) *gin.Engine {
	health := health.Controller{Validate: validate, Logger: logger}

	authUrl := r.Group(fmt.Sprintf("/api/%v", ApiVersion))
	{
		authUrl.POST("/health", health.Post)
		authUrl.GET("/health", health.Get)
	}
	return r
}
