package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lordscoba/bible_compass_backend/internal/model"
	"github.com/lordscoba/bible_compass_backend/utility"
)

type Controller struct {
	Validate *validator.Validate
	Logger   *utility.Logger
}

func (base *Controller) CreateCategory(c *gin.Context) {

	// bind userdetails to User struct
	var Category model.Category
	err := c.Bind(&Category)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Unable to bind category details", err, nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	err = base.Validate.Struct(&Category)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Validation failed", utility.ValidationResponse(err, base.Validate), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	// CategoryResponse, msg, code, err := admin.AdminCreateUser(Category)
	// if err != nil {
	// 	rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
	// 	c.JSON(code, rd)
	// 	return
	// }

	rd := utility.BuildSuccessResponse(http.StatusCreated, "user created successfully", Category)
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) UpdateCategory(c *gin.Context) {

	var _ string = c.Param("id")

	// bind userdetails to User struct
	var Category model.Category
	err := c.Bind(&Category)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Unable to bind category update details", err, nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	err = base.Validate.Struct(&Category)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Validation failed", utility.ValidationResponse(err, base.Validate), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	// CategoryResponse, msg, code, err := admin.AdminUpdateUser(Category, id)
	// if err != nil {
	// 	rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
	// 	c.JSON(code, rd)
	// 	return
	// }

	rd := utility.BuildSuccessResponse(http.StatusCreated, "Category updated successfully", Category)
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) GetCategory(c *gin.Context) {

	// userResponse, msg, code, err := admin.AdminGetUser()
	// if err != nil {
	// 	rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
	// 	c.JSON(code, rd)
	// 	return
	// }

	rd := utility.BuildSuccessResponse(http.StatusOK, "Users Gotten successfully", "put data here")
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) GetCategoryById(c *gin.Context) {
	var _ string = c.Param("id")
	// userResponse, msg, code, err := admin.AdminGetUserbyId(id)
	// if err != nil {
	// 	rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
	// 	c.JSON(code, rd)
	// 	return
	// }

	rd := utility.BuildSuccessResponse(http.StatusOK, "User Gotten successfully", "put data here")
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) DeleteCategoryById(c *gin.Context) {
	var _ string = c.Param("id")
	// CategoryResponse, msg, code, err := admin.AdminDeleteCategorybyId(id)
	// if err != nil {
	// 	rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
	// 	c.JSON(code, rd)
	// 	return
	// }

	rd := utility.BuildSuccessResponse(http.StatusOK, "Category deleted successfully", "Put data here")
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) CategoryInfo(c *gin.Context) {

	// CategoryResponse, msg, code, err := admin.AdminCategoryInfo()
	// if err != nil {
	// 	rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
	// 	c.JSON(code, rd)
	// 	return
	// }

	rd := utility.BuildSuccessResponse(http.StatusOK, "Category Info successfully", "put data here")
	c.JSON(http.StatusOK, rd)

}
