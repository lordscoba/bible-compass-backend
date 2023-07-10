package subscription

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/lordscoba/bible_compass_backend/internal/constants"
	"github.com/lordscoba/bible_compass_backend/internal/model"
	"github.com/lordscoba/bible_compass_backend/pkg/repository/paystack"
	"github.com/lordscoba/bible_compass_backend/pkg/repository/storage/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InitializePaymentOne(subscriptionRes model.InitializePaymentModel, id string) (model.InitializeResponse, string, int, error) {

	subscriptionRes.Amount = 100

	idHash, _ := primitive.ObjectIDFromHex(id)

	// check if id  exists
	idsearch := map[string]any{
		"_id": idHash,
	}
	idCount, _ := mongodb.MongoCount(constants.UserCollection, idsearch)
	if idCount < 1 {
		return model.InitializeResponse{}, "user does not exist", 403, errors.New("user does not exist in database")
	}

	// get from db
	var resultTwo model.User
	result, err := mongodb.MongoGetOne(constants.UserCollection, idsearch)
	if err != nil {
		return model.InitializeResponse{}, "Unable to get user from database", 500, err
	}
	result.Decode(&resultTwo)
	// get from db end

	var subscription model.Subscription
	subscription.Amount = float64(subscriptionRes.Amount)
	subscription.UserID = idHash
	subscription.Username = resultTwo.Username
	subscription.ID = primitive.NewObjectID()
	subscription.DateCreated = time.Now().Local()
	subscription.DateUpdated = time.Now().Local()
	subscription.Email = subscriptionRes.Email
	subscription.Reference = primitive.NewObjectID().Hex()

	// save to DB
	_, err = mongodb.MongoPost(constants.SubscriptionCollection, subscription)
	if err != nil {
		return model.InitializeResponse{}, "Unable to save sub to database", 500, err
	}

	payload := map[string]interface{}{
		"amount":    subscriptionRes.Amount * 100,
		"email":     subscriptionRes.Email,
		"reference": subscription.Reference,
	}

	response, err := paystack.PaystackInitPost(payload)
	if err != nil {
		return model.InitializeResponse{}, "Payment initiation failed", response.StatusCode(), err
	}

	// Unmarshal the response body into the struct
	var resultOne model.InitializeResponse
	err = json.Unmarshal(response.Body(), &resultOne)
	if err != nil {
		return model.InitializeResponse{}, "Payment initiation failed", 0, err
	}

	if !response.IsSuccess() {
		return model.InitializeResponse{}, resultOne.Message, response.StatusCode(), errors.New("payment initiation failed")
	}

	return resultOne, "Initiation success", 0, nil
}
