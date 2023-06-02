package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lordscoba/bible_compass_backend/pkg/handler/verses"
	"github.com/lordscoba/bible_compass_backend/utility"
)

func Verses(r *gin.Engine, validate *validator.Validate, ApiVersion string, logger *utility.Logger) *gin.Engine {
	verses := verses.Controller{Validate: validate, Logger: logger}

	verseUrl := r.Group(fmt.Sprintf("/api/%v", ApiVersion))
	{
		verseUrl.POST("/admin/createverse/:id", verses.CreateVerses)
		verseUrl.PATCH("/admin/updateverse/:id", verses.UpdateVerses)
		verseUrl.GET("/admin/getverses", verses.GetVerses)
		verseUrl.GET("/admin/getverse/:id", verses.GetVersesById)
		verseUrl.GET("/admin/verseinfo", verses.VersesInfo)
		verseUrl.DELETE("/admin/deleteverse/:id", verses.DeleteVersesById)
	}
	return r
}
