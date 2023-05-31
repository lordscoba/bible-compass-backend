package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lordscoba/bible_compass_backend/pkg/handler/profile"
	"github.com/lordscoba/bible_compass_backend/utility"
)

func Profile(r *gin.Engine, validate *validator.Validate, ApiVersion string, logger *utility.Logger) *gin.Engine {
	profile := profile.Controller{Validate: validate, Logger: logger}

	profileUrl := r.Group(fmt.Sprintf("/api/%v", ApiVersion))
	{
		profileUrl.PATCH("/profile/:id", profile.UpdateProfile)
		profileUrl.GET("/profiledetails/:id", profile.ProfileDetails)
	}
	return r
}
