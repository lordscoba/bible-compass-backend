package verses

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lordscoba/bible_compass_backend/service/verses"
	"github.com/lordscoba/bible_compass_backend/utility"
)

func (base *Controller) AiBible(c *gin.Context) {

	var passage string = c.DefaultQuery("passage", "john3:1-10")

	VerseResponse, msg, code, err := verses.AiBibleService(passage)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusCreated, "bible verse found successfully", VerseResponse)
	c.JSON(http.StatusOK, rd)

}
