package subscription

import (
	"context"
	"errors"
	"fmt"
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
		return model.SubscriptionResponse{}, "Unable to save sub to database", 500, err
	}

	// update user
	var user model.User
	user.ID = idHash
	user.Upgrade = true

	// save to DB
	_, err = mongodb.MongoUpdate(idsearch, user, constants.UserCollection)
	if err != nil {
		return model.SubscriptionResponse{}, "Unable to update user to database", 500, err
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
	idCount, _ := mongodb.MongoCount(constants.SubscriptionCollection, idsearch)
	if idCount < 1 {
		return model.SubscriptionResponse{}, "subscription does not exist", 403, errors.New("user does not exist in database")
	}

	subscription.ID = idHash
	subscription.DateUpdated = time.Now().Local()
	// save to DB
	_, err := mongodb.MongoUpdate(idsearch, subscription, constants.SubscriptionCollection)
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

func AdminGetSubscription(searchText map[string]string) ([]model.Subscription, string, int, error) {

	// get from db
	result, err := mongodb.MongoGetAll(constants.SubscriptionCollection, searchText)
	if err != nil {
		return []model.Subscription{}, "Unable to save subscription to database", 500, err
	}

	var subscription = make([]model.Subscription, 0)
	result.All(context.TODO(), &subscription)
	return subscription, "", 0, nil
}

func AdminGetSubscriptionbyId(id string) (model.Subscription, string, int, error) {
	idHash, _ := primitive.ObjectIDFromHex(id)
	search := map[string]any{
		"_id": idHash,
	}
	// get from db
	var resultOne model.Subscription
	result, err := mongodb.MongoGetOne(constants.SubscriptionCollection, search)
	if err != nil {
		return model.Subscription{}, "Unable to get subscription from database", 500, err
	}
	result.Decode(&resultOne)
	// get from db end

	fmt.Println(resultOne)
	return resultOne, "", 0, nil
}

func AdminDeleteSubscriptionbyId(id string) (int64, string, int, error) {

	idHash, _ := primitive.ObjectIDFromHex(id)
	search := map[string]any{
		"_id": idHash,
	}

	// check if id exists
	idCount, _ := mongodb.MongoCount(constants.SubscriptionCollection, search)
	if idCount < 1 {
		return 0, "Subscription does not exist", 403, errors.New("subscription does not exist in database")
	}

	// get from db
	result, err := mongodb.MongoDelete(constants.SubscriptionCollection, search)
	if err != nil {
		return 0, "Unable to save subscription to database", 500, err
	}

	fmt.Println(result.DeletedCount)
	return result.DeletedCount, "", 0, nil
}

func AdminGetUserSubService(userid string, searchText map[string]string) ([]model.Subscription, string, int, error) {

	userIdHash, _ := primitive.ObjectIDFromHex(userid)
	search := map[string]any{
		"_id": userIdHash,
	}

	// check if id exists
	idCount, _ := mongodb.MongoCount(constants.UserCollection, search)
	if idCount < 1 {
		return []model.Subscription{}, "user does not exist", 403, errors.New("subscription does not exist in database")
	}

	searchcol := map[string]any{
		"user_id": userIdHash,
	}
	// get from db
	result, err := mongodb.MongoGet(constants.SubscriptionCollection, searchcol, searchText)
	if err != nil {
		return []model.Subscription{}, "Unable to get subscription to database", 500, err
	}

	var subscription = make([]model.Subscription, 0)
	result.All(context.TODO(), &subscription)
	return subscription, "", 0, nil

}

func AdminSubscriptionInfo() (model.SubscriptionInfoResponse, string, int, error) {
	// total users
	search := map[string]any{}
	TotalSubscription, err := mongodb.MongoCount(constants.SubscriptionCollection, search)
	if err != nil {
		return model.SubscriptionInfoResponse{}, "Unable to get count", 500, err
	}

	// total subscribers category
	StatusSearch1 := map[string]any{
		"status": true,
	}
	ActiveSubscription, err := mongodb.MongoCount(constants.SubscriptionCollection, StatusSearch1)
	if err != nil {
		return model.SubscriptionInfoResponse{}, "Unable to get count", 500, err
	}

	// total keywords
	StatusSearch2 := map[string]any{
		"status": false,
	}
	InActiveSubscription, err := mongodb.MongoCount(constants.SubscriptionCollection, StatusSearch2)
	if err != nil {
		return model.SubscriptionInfoResponse{}, "Unable to get count", 500, err
	}

	CategoryInfo := model.SubscriptionInfoResponse{
		TotalSubscription:    TotalSubscription,
		ActiveSubscription:   ActiveSubscription,
		InActiveSubscription: InActiveSubscription,
	}

	return CategoryInfo, "", 0, nil
}

func AdminGetUserSubServiceStats(userid string) (model.SubscriptionInfoResponse, string, int, error) {

	userIdHash, _ := primitive.ObjectIDFromHex(userid)
	search := map[string]any{
		"user_id": userIdHash,
	}

	// total users

	TotalSubscription, err := mongodb.MongoCount(constants.SubscriptionCollection, search)
	if err != nil {
		return model.SubscriptionInfoResponse{}, "Unable to get count", 500, err
	}

	// total subscribers category
	StatusSearch1 := map[string]any{
		"user_id": userIdHash,
		"status":  true,
	}
	ActiveSubscription, err := mongodb.MongoCount(constants.SubscriptionCollection, StatusSearch1)
	if err != nil {
		return model.SubscriptionInfoResponse{}, "Unable to get count", 500, err
	}

	// total keywords
	StatusSearch2 := map[string]any{
		"user_id": userIdHash,
		"status":  false,
	}
	InActiveSubscription, err := mongodb.MongoCount(constants.SubscriptionCollection, StatusSearch2)
	if err != nil {
		return model.SubscriptionInfoResponse{}, "Unable to get count", 500, err
	}

	CategoryInfo := model.SubscriptionInfoResponse{
		TotalSubscription:    TotalSubscription,
		ActiveSubscription:   ActiveSubscription,
		InActiveSubscription: InActiveSubscription,
	}

	return CategoryInfo, "", 0, nil
}
