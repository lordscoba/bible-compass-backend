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

	// SubscriptionResponse, msg, code, err := admin.AdminGetSubscription()
	// if err != nil {
	// 	rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
	// 	c.JSON(code, rd)
	// 	return
	// }

	rd := utility.BuildSuccessResponse(http.StatusOK, "Subscription Gotten successfully", "put data here")
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) GetSubscriptionById(c *gin.Context) {
	var _ string = c.Param("id")
	// SubscriptionResponse, msg, code, err := admin.AdminGetSubscriptionbyId(id)
	// if err != nil {
	// 	rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
	// 	c.JSON(code, rd)
	// 	return
	// }

	rd := utility.BuildSuccessResponse(http.StatusOK, "Subscription Gotten successfully", "put data here")
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) DeleteSubscriptionById(c *gin.Context) {
	var _ string = c.Param("id")
	// SubscriptionResponse, msg, code, err := admin.AdminDeleteSubscriptionbyId(id)
	// if err != nil {
	// 	rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
	// 	c.JSON(code, rd)
	// 	return
	// }

	rd := utility.BuildSuccessResponse(http.StatusOK, "Subscription deleted successfully", "Put data here")
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) SubscriptionInfo(c *gin.Context) {

	// SubscriptionResponse, msg, code, err := admin.AdminSubscriptionInfo()
	// if err != nil {
	// 	rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
	// 	c.JSON(code, rd)
	// 	return
	// }

	rd := utility.BuildSuccessResponse(http.StatusOK, "Subscription Info successfully", "put data here")
	c.JSON(http.StatusOK, rd)

}
