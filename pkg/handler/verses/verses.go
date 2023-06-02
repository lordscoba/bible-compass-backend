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

	var id string = c.Param("id")
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

	VerseResponse, msg, code, err := verses.AdminCreateVerses(Verses, id)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusCreated, "Verses created successfully", VerseResponse)
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) UpdateVerses(c *gin.Context) {

	var _ string = c.Param("id")

	// bind userdetails to User struct
	var verses model.BibleVerse
	err := c.Bind(&verses)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Unable to bind Verse update details", err, nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	err = base.Validate.Struct(&verses)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Validation failed", utility.ValidationResponse(err, base.Validate), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	// VerseResponse, msg, code, err := verses.AdminUpdateVerse(verses, id)
	// if err != nil {
	// 	rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
	// 	c.JSON(code, rd)
	// 	return
	// }

	rd := utility.BuildSuccessResponse(http.StatusCreated, "Verses updated successfully", "")
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) GetVerses(c *gin.Context) {

	// VerseResponse, msg, code, err := verses.AdminGetVerse()
	// if err != nil {
	// 	rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
	// 	c.JSON(code, rd)
	// 	return
	// }

	rd := utility.BuildSuccessResponse(http.StatusOK, "Verses Gotten successfully", "")
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) GetVersesById(c *gin.Context) {
	var _ string = c.Param("id")
	// VerseResponse, msg, code, err := verses.AdminGetVersebyId(id)
	// if err != nil {
	// 	rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
	// 	c.JSON(code, rd)
	// 	return
	// }

	rd := utility.BuildSuccessResponse(http.StatusOK, "Verses Gotten successfully", "")
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) DeleteVersesById(c *gin.Context) {
	var _ string = c.Param("id")
	// VerseResponse, msg, code, err := verses.AdminDeleteVersebyId(id)
	// if err != nil {
	// 	rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
	// 	c.JSON(code, rd)
	// 	return
	// }

	rd := utility.BuildSuccessResponse(http.StatusOK, "Verses deleted successfully", "")
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) VersesInfo(c *gin.Context) {

	// VerseResponse, msg, code, err := verses.AdminVerseInfo()
	// if err != nil {
	// 	rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
	// 	c.JSON(code, rd)
	// 	return
	// }

	rd := utility.BuildSuccessResponse(http.StatusOK, "Verses Info successfully", "")
	c.JSON(http.StatusOK, rd)

}
