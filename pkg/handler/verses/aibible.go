package verses

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lordscoba/bible_compass_backend/service/verses"
	"github.com/lordscoba/bible_compass_backend/utility"
)

func (base *Controller) AiBible(c *gin.Context) {

	var passage string = c.DefaultQuery("passage", "john3:16")

	// var kid string = c.Param("kid")
	// Find := map[string]string{
	// 	"passage": c.DefaultQuery("passage", ""),
	// }
	// bind userdetails to User struct
	// var Find model.AiBible
	// err := c.Bind(&Find)
	// if err != nil {
	// 	rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Unable to bind Verses details", err, nil)
	// 	c.JSON(http.StatusBadRequest, rd)
	// 	return
	// }

	// err = base.Validate.Struct(&Find)
	// if err != nil {
	// 	rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Validation failed", utility.ValidationResponse(err, base.Validate), nil)
	// 	c.JSON(http.StatusBadRequest, rd)
	// 	return
	// }

	VerseResponse, msg, code, err := verses.AiBibleService(passage)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusCreated, "bible verse found successfully", VerseResponse)
	c.JSON(http.StatusOK, rd)

}
