package subscription

import (
	"errors"
	"time"

	"github.com/lordscoba/bible_compass_backend/internal/constants"
	"github.com/lordscoba/bible_compass_backend/internal/model"
	"github.com/lordscoba/bible_compass_backend/pkg/repository/storage/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AdminCreateSubscription(subscription model.Subscription, id string) (model.SubscriptionResponse, string, int, error) {

	idHash, _ := primitive.ObjectIDFromHex(id)

	// check if id  exists
	idsearch := map[string]any{
		"_id": idHash,
	}
	idCount, _ := mongodb.MongoCount(constants.UserCollection, idsearch)
	if idCount < 1 {
		return model.SubscriptionResponse{}, "user does not exist", 403, errors.New("user does not exist in database")
	}

	// get from db
	var resultOne model.User
	result, err := mongodb.MongoGetOne(constants.UserCollection, idsearch)
	if err != nil {
		return model.SubscriptionResponse{}, "Unable to get user from database", 500, err
	}
	result.Decode(&resultOne)
	// get from db end

	subscription.UserID = idHash
	subscription.Username = resultOne.Username
	subscription.ID = primitive.NewObjectID()
	subscription.DateCreated = time.Now().Local()
	subscription.DateUpdated = time.Now().Local()

	// save to DB
	_, err = mongodb.MongoPost(constants.SubscriptionCollection, subscription)
	if err != nil {
		return model.SubscriptionResponse{}, "Unable to save user to database", 500, err
	}

	SubscriptionResponse := model.SubscriptionResponse{
		Username:     subscription.Username,
		Amount:       subscription.Amount,
		Status:       subscription.Status,
		Duration:     subscription.Duration,
		DateExpiring: subscription.DateExpiring,
	}
	return SubscriptionResponse, "", 0, nil
}

func AdminUpdateSubscription(subscription model.Subscription, id string) (model.SubscriptionResponse, string, int, error) {

	idHash, _ := primitive.ObjectIDFromHex(id)

	idsearch := map[string]any{
		"_id": idHash,
	}

	// check if id exists
	idCount, _ := mongodb.MongoCount(constants.UserCollection, idsearch)
	if idCount < 1 {
		return model.SubscriptionResponse{}, "user does not exist", 403, errors.New("user does not exist in database")
	}

	subscription.UserID = idHash
	subscription.DateUpdated = time.Now().Local()
	idSearchCollection := map[string]any{
		"user_id": idHash,
		// "username": "lordscoba",
	}
	// save to DB
	_, err := mongodb.MongoUpdate(idSearchCollection, subscription, constants.SubscriptionCollection)
	if err != nil {
		return model.SubscriptionResponse{}, "Unable to save user to database", 500, err
	}

	SubscriptionResponse := model.SubscriptionResponse{
		Username:     subscription.Username,
		Amount:       subscription.Amount,
		Status:       subscription.Status,
		Duration:     subscription.Duration,
		DateExpiring: subscription.DateExpiring,
		Type:         subscription.Type,
	}
	return SubscriptionResponse, "", 0, nil
}
