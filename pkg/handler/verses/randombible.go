package verses

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lordscoba/bible_compass_backend/service/verses"
	"github.com/lordscoba/bible_compass_backend/utility"
)

func (base *Controller) RandomBible(c *gin.Context) {

	VerseResponse, msg, code, err := verses.RandomBibleService()
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "random verse found successfully", VerseResponse)
	c.JSON(http.StatusOK, rd)

}
