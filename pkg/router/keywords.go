package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lordscoba/bible_compass_backend/pkg/handler/keywords"
	"github.com/lordscoba/bible_compass_backend/utility"
)

func Keywords(r *gin.Engine, validate *validator.Validate, ApiVersion string, logger *utility.Logger) *gin.Engine {
	keywords := keywords.Controller{Validate: validate, Logger: logger}

	keywordsUrl := r.Group(fmt.Sprintf("/api/%v", ApiVersion))
	{
		keywordsUrl.POST("/admin/createkeywords", keywords.CreateKeywords)
		keywordsUrl.PATCH("/admin/updatekeywords/:id", keywords.UpdateKeywords)
		keywordsUrl.DELETE("/admin/deletekeywords/:id", keywords.DeleteKeywordsById)
		keywordsUrl.GET("/admin/getkeywords", keywords.GetKeywords)
		keywordsUrl.GET("/admin/getkeywordsid/:id", keywords.GetKeywordsById)
		keywordsUrl.GET("/admin/keywordsinfo", keywords.KeywordsInfo)
	}
	return r
}
