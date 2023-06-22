package fav

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lordscoba/bible_compass_backend/service/favorite"
	"github.com/lordscoba/bible_compass_backend/utility"
)

type Controller struct {
	Validate *validator.Validate
	Logger   *utility.Logger
}

func (base *Controller) LikeKeyword(c *gin.Context) {

	var keyword string = c.Param("keyword")
	var email string = c.Param("email")

	favResponse, msg, code, err := favorite.LikeKeywordsService(email, keyword)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusCreated, "keyword liked successfully", favResponse)
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) UnLikeKeyword(c *gin.Context) {

	var keyword string = c.Param("keyword")
	var email string = c.Param("email")

	favResponse, msg, code, err := favorite.UnLikeKeywordServive(email, keyword)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusCreated, "keyword unliked successfully", favResponse)
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) GetKeywordStatus(c *gin.Context) {

	var keyword string = c.Param("keyword")
	var email string = c.Param("email")
	favResponse, msg, code, err := favorite.GetFavStatus(email, keyword)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusCreated, "status gotten successfully", favResponse)
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) GetFavKeywords(c *gin.Context) {
	var email string = c.Param("email")

	favResponse, msg, code, err := favorite.GetFavKeywordsService(email)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusCreated, "favourites gotten successfully", favResponse)
	c.JSON(http.StatusOK, rd)

}
