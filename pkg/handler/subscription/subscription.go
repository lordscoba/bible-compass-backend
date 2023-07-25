package subscription

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lordscoba/bible_compass_backend/internal/model"
	"github.com/lordscoba/bible_compass_backend/service/subscription"
	"github.com/lordscoba/bible_compass_backend/utility"
)

type Controller struct {
	Validate *validator.Validate
	Logger   *utility.Logger
}

func (base *Controller) InitializePayment(c *gin.Context) {

	var id string = c.Param("id")
	// bind userdetails to User struct
	var Subscription model.InitializePaymentModel
	err := c.Bind(&Subscription)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Unable to bind category details", err, nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	err = base.Validate.Struct(&Subscription)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Validation failed", utility.ValidationResponse(err, base.Validate), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	InitResponse, msg, code, err := subscription.InitializePaymentOne(Subscription, id)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusCreated, "payment initialized successfully", InitResponse)
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) VerifyPaymentById(c *gin.Context) {
	var rid string = c.Param("rid")
	subscriptionResponse, msg, code, err := subscription.VerifyPaymentByIdService(rid)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "Subscription Gotten successfully", subscriptionResponse)
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) CreateSubscription(c *gin.Context) {

	var id string = c.Param("id")
	// bind userdetails to User struct
	var Subscription model.Subscription
	err := c.Bind(&Subscription)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Unable to bind category details", err, nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	err = base.Validate.Struct(&Subscription)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Validation failed", utility.ValidationResponse(err, base.Validate), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	SubscriptionResponse, msg, code, err := subscription.AdminCreateSubscription(Subscription, id)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusCreated, "Subscription created successfully", SubscriptionResponse)
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) UpdateSubscription(c *gin.Context) {

	var id string = c.Param("id")

	// bind userdetails to User struct
	var Subscription model.Subscription
	err := c.Bind(&Subscription)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Unable to bind category update details", err, nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	err = base.Validate.Struct(&Subscription)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Validation failed", utility.ValidationResponse(err, base.Validate), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	SubscriptionResponse, msg, code, err := subscription.AdminUpdateSubscription(Subscription, id)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusCreated, "subscription updated successfully", SubscriptionResponse)
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) GetSubscription(c *gin.Context) {

	searchText := map[string]string{
		"username": c.DefaultQuery("username", ""),
		"type":     c.DefaultQuery("type", ""),
	}

	subscriptionResponse, msg, code, err := subscription.AdminGetSubscription(searchText)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "Subscription Gotten successfully", subscriptionResponse)
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) GetSubscriptionById(c *gin.Context) {
	var id string = c.Param("id")
	subscriptionResponse, msg, code, err := subscription.AdminGetSubscriptionbyId(id)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "Subscription Gotten successfully", subscriptionResponse)
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) GetUserSub(c *gin.Context) {
	subscription.UpdateSubStatus()
	var userId string = c.Param("userId")

	searchText := map[string]string{
		"type": c.DefaultQuery("type", ""),
	}
	subscriptionResponse, msg, code, err := subscription.AdminGetUserSubService(userId, searchText)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "Subscription Gotten successfully", subscriptionResponse)
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) DeleteSubscriptionById(c *gin.Context) {
	var id string = c.Param("id")
	subscriptionResponse, msg, code, err := subscription.AdminDeleteSubscriptionbyId(id)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "Subscription deleted successfully", subscriptionResponse)
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) SubscriptionInfo(c *gin.Context) {

	SubscriptionResponse, msg, code, err := subscription.AdminSubscriptionInfo()
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "Subscription Info successfully", SubscriptionResponse)
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) GetUserSubStats(c *gin.Context) {
	var userId string = c.Param("userId")

	subscriptionResponse, msg, code, err := subscription.AdminGetUserSubServiceStats(userId)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "Subscription info gotten successfully", subscriptionResponse)
	c.JSON(http.StatusOK, rd)

}
