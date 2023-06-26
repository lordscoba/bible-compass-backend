package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lordscoba/bible_compass_backend/internal/model"
	"github.com/lordscoba/bible_compass_backend/service/admin"
	"github.com/lordscoba/bible_compass_backend/utility"
)

type Controller struct {
	Validate *validator.Validate
	Logger   *utility.Logger
}

func (base *Controller) CreateUser(c *gin.Context) {

	// bind userdetails to User struct
	var User model.User
	err := c.Bind(&User)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Unable to bind user signup details", err, nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	err = base.Validate.Struct(&User)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Validation failed", utility.ValidationResponse(err, base.Validate), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	userResponse, msg, code, err := admin.AdminCreateUser(User)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusCreated, "user created successfully", userResponse)
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) UpdateUser(c *gin.Context) {

	var id string = c.Param("id")

	// bind userdetails to User struct
	var User model.User
	err := c.Bind(&User)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Unable to bind user update details", err, nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	err = base.Validate.Struct(&User)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Validation failed", utility.ValidationResponse(err, base.Validate), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	userResponse, msg, code, err := admin.AdminUpdateUser(User, id)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusCreated, "User updated successfully", userResponse)
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) GetUsers(c *gin.Context) {

	searchText := map[string]string{
		"username": c.DefaultQuery("username", ""),
		"email":    c.DefaultQuery("email", ""),
	}

	userResponse, msg, code, err := admin.AdminGetUser(searchText)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "Users Gotten successfully", userResponse)
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) GetUsersById(c *gin.Context) {
	var id string = c.Param("id")
	userResponse, msg, code, err := admin.AdminGetUserbyId(id)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "User Gotten successfully", userResponse)
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) DeleteUsersById(c *gin.Context) {
	var id string = c.Param("id")
	userResponse, msg, code, err := admin.AdminDeleteUserbyId(id)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "Users deleted successfully", userResponse)
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) UsersInfo(c *gin.Context) {

	userResponse, msg, code, err := admin.AdminUsersInfo()
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "Users Info successfully", userResponse)
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) VerifyUser(c *gin.Context) {

	// var id string = c.Param("id")

	// bind userdetails to User struct
	var User model.User
	err := c.Bind(&User)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Unable to bind user update details", err, nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	err = base.Validate.Struct(&User)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Validation failed", utility.ValidationResponse(err, base.Validate), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	userResponse, msg, code, err := admin.AdminVerifyUser(User)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusCreated, "User verified successfully", userResponse)
	c.JSON(http.StatusOK, rd)

}
