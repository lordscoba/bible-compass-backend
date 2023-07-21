package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lordscoba/bible_compass_backend/pkg/handler/auth"
	"github.com/lordscoba/bible_compass_backend/utility"
)

func Auth(r *gin.Engine, validate *validator.Validate, ApiVersion string, logger *utility.Logger) *gin.Engine {
	auth := auth.Controller{Validate: validate, Logger: logger}

	authUrl := r.Group(fmt.Sprintf("/api/%v", ApiVersion))
	{
		authUrl.POST("/signup", auth.Signup)
		authUrl.POST("/login", auth.Login)
		authUrl.POST("/verify", auth.Verify)
	}
	return r
}
