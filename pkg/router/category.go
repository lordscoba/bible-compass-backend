package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lordscoba/bible_compass_backend/pkg/handler/category"
	"github.com/lordscoba/bible_compass_backend/utility"
)

func Category(r *gin.Engine, validate *validator.Validate, ApiVersion string, logger *utility.Logger) *gin.Engine {
	category := category.Controller{Validate: validate, Logger: logger}

	categoryUrl := r.Group(fmt.Sprintf("/api/%v", ApiVersion))
	{
		categoryUrl.POST("/admin/createcategory", category.CreateCategory)
		categoryUrl.PATCH("/admin/updatecategory/:id", category.UpdateCategory)
		categoryUrl.DELETE("/admin/deletecategory/:id", category.DeleteCategoryById)
		categoryUrl.GET("/admin/getcategory", category.GetCategory)
		categoryUrl.GET("/admin/getcategoryid/:id", category.GetCategoryById)
		categoryUrl.GET("/admin/categoryinfo", category.CategoryInfo)
	}
	return r
}
