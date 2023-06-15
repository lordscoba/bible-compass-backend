package category

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lordscoba/bible_compass_backend/internal/model"
	"github.com/lordscoba/bible_compass_backend/service/category"
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

	CategoryResponse, msg, code, err := category.AdminCreateCategory(Category)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusCreated, "category created successfully", CategoryResponse)
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) UpdateCategory(c *gin.Context) {

	var id string = c.Param("id")

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

	CategoryResponse, msg, code, err := category.AdminUpdatecategory(Category, id)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusCreated, "Category updated successfully", CategoryResponse)
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) GetCategory(c *gin.Context) {

	categoryResponse, msg, code, err := category.AdminGetCategory()
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "Category Gotten successfully", categoryResponse)
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) GetCategoryById(c *gin.Context) {
	var id string = c.Param("id")
	categoryResponse, msg, code, err := category.AdminGetCategorybyId(id)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "category Gotten successfully", categoryResponse)
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) DeleteCategoryById(c *gin.Context) {
	var id string = c.Param("id")
	categoryResponse, msg, code, err := category.AdminDeleteCategorybyId(id)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "Category deleted successfully", categoryResponse)
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) CategoryInfo(c *gin.Context) {

	CategoryResponse, msg, code, err := category.AdminCategoryInfo()
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "Category Info successfully", CategoryResponse)
	c.JSON(http.StatusOK, rd)

}
