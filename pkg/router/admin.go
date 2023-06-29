package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lordscoba/bible_compass_backend/pkg/handler/admin"
	"github.com/lordscoba/bible_compass_backend/utility"
)

func Admin(r *gin.Engine, validate *validator.Validate, ApiVersion string, logger *utility.Logger) *gin.Engine {
	admin := admin.Controller{Validate: validate, Logger: logger}

	adminUrl := r.Group(fmt.Sprintf("/api/%v", ApiVersion))

	{
		adminUrl.POST("/admin/createuser", admin.CreateUser)
		adminUrl.PATCH("/admin/updateuser/:id", admin.UpdateUser)
		adminUrl.GET("/admin/getusers", admin.GetUsers)
		adminUrl.GET("/admin/getuser/:id", admin.GetUsersById)
		adminUrl.GET("/admin/usersinfo", admin.UsersInfo)
		adminUrl.DELETE("/admin/deleteuser/:id", admin.DeleteUsersById)
		adminUrl.PATCH("/admin/verifyuser", admin.VerifyUser)
	}
	return r
}
