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
	var catid string = c.Param("catid")

	searchText := map[string]string{
		"keyword": c.DefaultQuery("keyword", ""),
	}

	keywordsResponse, msg, code, err := keywords.AdminGetKeywords(catid, searchText)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "Keywords Gotten successfully", keywordsResponse)
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) GetKeywordsById(c *gin.Context) {
	var id string = c.Param("id")
	keywordsResponse, msg, code, err := keywords.AdminGetkeywordsbyId(id)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "Keywords Gotten successfully", keywordsResponse)
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) DeleteKeywordsById(c *gin.Context) {
	var id string = c.Param("id")
	keywordsResponse, msg, code, err := keywords.AdminDeletekeywordsbyId(id)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "Keywords deleted successfully", keywordsResponse)
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) KeywordsInfo(c *gin.Context) {

	KeywordsResponse, msg, code, err := keywords.AdminKeywordsInfo()
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "Keywords Info successfully", KeywordsResponse)
	c.JSON(http.StatusOK, rd)

}
