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

	keywordsUrl := r.Group(fmt.Sprintf("/api/%v", ApiVersion))
	{
		keywordsUrl.POST("/admin/createsubscription", subscription.CreateSubscription)
		keywordsUrl.PATCH("/admin/updatesubscription/:id", subscription.UpdateSubscription)
		keywordsUrl.DELETE("/admin/deletesubscription/:id", subscription.DeleteSubscriptionById)
		keywordsUrl.GET("/admin/getsubscription", subscription.GetSubscription)
		keywordsUrl.GET("/admin/getsubscriptionid/:id", subscription.GetSubscriptionById)
		keywordsUrl.GET("/admin/subscriptioninfo", subscription.SubscriptionInfo)
	}
	return r
}
