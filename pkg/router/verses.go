package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lordscoba/bible_compass_backend/pkg/handler/verses"
	"github.com/lordscoba/bible_compass_backend/pkg/middleware"
	"github.com/lordscoba/bible_compass_backend/utility"
)

func Verses(r *gin.Engine, validate *validator.Validate, ApiVersion string, logger *utility.Logger) *gin.Engine {
	verses := verses.Controller{Validate: validate, Logger: logger}

	verseUrl := r.Group(fmt.Sprintf("/api/%v", ApiVersion))
	{

		// for AI bible
		verseUrl.GET("/randombible", verses.RandomBible)
		verseUrl.GET("/aibible", verses.AiBible) // this is "?passage=john3:16-18"

		verseUrl.GET("/admin/getverses/:kid", verses.GetVerses)
		verseUrl.GET("/admin/getverse/:kid/:Bid", verses.GetVersesById)
		verseUrl.GET("/admin/verseinfo/:kid", verses.VersesInfo)

		// this requires middleware
		verseUrl.Use(middleware.AuthMiddleware())
		verseUrl.POST("/admin/createverse/:kid", verses.CreateVerses)
		verseUrl.PATCH("/admin/updateverse/:kid/:Bid", verses.UpdateVerses)
		verseUrl.DELETE("/admin/deleteverse/:kid/:Bid", verses.DeleteVersesById)

	}
	return r
}
