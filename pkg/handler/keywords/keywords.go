package keywords

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lordscoba/bible_compass_backend/internal/model"
	"github.com/lordscoba/bible_compass_backend/service/keywords"
	"github.com/lordscoba/bible_compass_backend/utility"
)

type Controller struct {
	Validate *validator.Validate
	Logger   *utility.Logger
}

func (base *Controller) CreateKeywords(c *gin.Context) {

	var id string = c.Param("id")

	// bind userdetails to User struct
	var Keywords model.Keywords
	err := c.Bind(&Keywords)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Unable to bind category details", err, nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	err = base.Validate.Struct(&Keywords)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Validation failed", utility.ValidationResponse(err, base.Validate), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	KeywordsResponse, msg, code, err := keywords.AdminCreateKeywords(Keywords, id)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusCreated, "keywords created successfully", KeywordsResponse)
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) UpdateKeywords(c *gin.Context) {

	var id string = c.Param("id")

	// bind userdetails to User struct
	var Keywords model.Keywords
	err := c.Bind(&Keywords)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Unable to bind category update details", err, nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	err = base.Validate.Struct(&Keywords)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Validation failed", utility.ValidationResponse(err, base.Validate), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	KeywordsResponse, msg, code, err := keywords.AdminUpdateKeywords(Keywords, id)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusCreated, "Keywords updated successfully", KeywordsResponse)
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) GetKeywords(c *gin.Context) {

	// KeywordsResponse, msg, code, err := admin.AdminGetKeywords()
	// if err != nil {
	// 	rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
	// 	c.JSON(code, rd)
	// 	return
	// }

	rd := utility.BuildSuccessResponse(http.StatusOK, "Keywords Gotten successfully", "put data here")
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) GetKeywordsById(c *gin.Context) {
	var _ string = c.Param("id")
	// keywordsResponse, msg, code, err := admin.AdminGetKeywordsbyId(id)
	// if err != nil {
	// 	rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
	// 	c.JSON(code, rd)
	// 	return
	// }

	rd := utility.BuildSuccessResponse(http.StatusOK, "Keywords Gotten successfully", "put data here")
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) DeleteKeywordsById(c *gin.Context) {
	var _ string = c.Param("id")
	// KeywordsResponse, msg, code, err := admin.AdminDeleteKeywordsbyId(id)
	// if err != nil {
	// 	rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
	// 	c.JSON(code, rd)
	// 	return
	// }

	rd := utility.BuildSuccessResponse(http.StatusOK, "Keywords deleted successfully", "Put data here")
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) KeywordsInfo(c *gin.Context) {

	// KeywordsResponse, msg, code, err := admin.AdminKeywordsInfo()
	// if err != nil {
	// 	rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
	// 	c.JSON(code, rd)
	// 	return
	// }

	rd := utility.BuildSuccessResponse(http.StatusOK, "Keywords Info successfully", "put data here")
	c.JSON(http.StatusOK, rd)

}
