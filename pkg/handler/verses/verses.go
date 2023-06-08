package verses

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lordscoba/bible_compass_backend/internal/model"
	"github.com/lordscoba/bible_compass_backend/service/verses"
	"github.com/lordscoba/bible_compass_backend/utility"
)

type Controller struct {
	Validate *validator.Validate
	Logger   *utility.Logger
}

func (base *Controller) CreateVerses(c *gin.Context) {

	var kid string = c.Param("kid")
	// bind userdetails to User struct
	var Verses model.BibleVerse
	err := c.Bind(&Verses)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Unable to bind Verses details", err, nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	err = base.Validate.Struct(&Verses)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Validation failed", utility.ValidationResponse(err, base.Validate), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	VerseResponse, msg, code, err := verses.AdminCreateVerses(Verses, kid)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusCreated, "Verses created successfully", VerseResponse)
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) UpdateVerses(c *gin.Context) {

	var kid string = c.Param("kid")
	var Bid string = c.Param("Bid")

	// bind userdetails to User struct
	var Verses model.BibleVerse
	err := c.Bind(&Verses)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Unable to bind Verse update details", err, nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	err = base.Validate.Struct(&Verses)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Validation failed", utility.ValidationResponse(err, base.Validate), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	VerseResponse, msg, code, err := verses.AdminUpdateVerse(Verses, kid, Bid)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusCreated, "Verses updated successfully", VerseResponse)
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) GetVerses(c *gin.Context) {
	var kid string = c.Param("kid")

	VerseResponse, msg, code, err := verses.AdminGetVerse(kid)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "Verses Gotten successfully", VerseResponse)
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) GetVersesById(c *gin.Context) {
	var kid string = c.Param("kid")
	var Bid string = c.Param("Bid")
	VerseResponse, msg, code, err := verses.AdminGetVersebyId(kid, Bid)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "Verses Gotten successfully", VerseResponse)
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) DeleteVersesById(c *gin.Context) {
	var kid string = c.Param("kid")
	var Bid string = c.Param("Bid")

	VerseResponse, msg, code, err := verses.AdminDeleteVersebyId(kid, Bid)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "Verses deleted successfully", VerseResponse)
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) VersesInfo(c *gin.Context) {

	var kid string = c.Param("kid")
	VerseResponse, msg, code, err := verses.AdminVerseInfo(kid)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "Verses Info successfully", VerseResponse)
	c.JSON(http.StatusOK, rd)

}
