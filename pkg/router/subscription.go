package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/lordscoba/bible_compass_backend/pkg/handler/subscription"
	"github.com/lordscoba/bible_compass_backend/utility"
)

func Subscription(r *gin.Engine, validate *validator.Validate, ApiVersion string, logger *utility.Logger) *gin.Engine {
	subscription := subscription.Controller{Validate: validate, Logger: logger}

	subUrl := r.Group(fmt.Sprintf("/api/%v", ApiVersion))
	{
		subUrl.POST("/admin/createsubscription/:id", subscription.CreateSubscription)
		subUrl.PATCH("/admin/updatesubscription/:id", subscription.UpdateSubscription)
		subUrl.DELETE("/admin/deletesubscription/:id", subscription.DeleteSubscriptionById)
		subUrl.GET("/admin/getsubscription", subscription.GetSubscription)
		subUrl.GET("/admin/getsubscriptionid/:id", subscription.GetSubscriptionById)
		subUrl.GET("/admin/subscriptioninfo", subscription.SubscriptionInfo)
		subUrl.GET("/admin/getusersub/:userId", subscription.GetUserSub)
		subUrl.GET("/admin/getusersubstats/:userId", subscription.GetUserSubStats)

		// paystack
		subUrl.POST("/user/initialize/:id", subscription.InitializePayment)      // user id
		subUrl.GET("/user/paystack/verify/:rid", subscription.VerifyPaymentById) // reference id

	}
	return r
}
