package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lordscoba/bible_compass_backend/pkg/handler/fav"
	"github.com/lordscoba/bible_compass_backend/utility"
)

func Fav(r *gin.Engine, validate *validator.Validate, ApiVersion string, logger *utility.Logger) *gin.Engine {
	fav := fav.Controller{Validate: validate, Logger: logger}

	favUrl := r.Group(fmt.Sprintf("/api/%v", ApiVersion))
	{
		favUrl.GET("/like/:keyword/:email", fav.LikeKeyword)
		favUrl.GET("/unlike/:keyword/:email", fav.UnLikeKeyword)
		favUrl.GET("/favstatus/:keyword/:email", fav.GetKeywordStatus)
		favUrl.GET("/getfavs/:email", fav.GetFavKeywords)
	}
	return r
}
