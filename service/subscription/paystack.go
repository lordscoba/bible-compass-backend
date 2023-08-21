package subscription

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/lordscoba/bible_compass_backend/internal/constants"
	"github.com/lordscoba/bible_compass_backend/internal/model"
	raterep "github.com/lordscoba/bible_compass_backend/pkg/repository/exchangerate"
	"github.com/lordscoba/bible_compass_backend/pkg/repository/paystack"
	"github.com/lordscoba/bible_compass_backend/pkg/repository/storage/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	currency string        = "NGN"
	amount   int64         = 2         //USD
	days     time.Duration = 30        // days
	hours    time.Duration = 24 * days // hours
)

func InitializePaymentOne(subscriptionRes model.InitializePaymentModel, id string) (model.InitializeResponse, string, int, error) {

	var exchangeRateData map[string]interface{}
	rateing, _ := raterep.RateRep()
	err2 := json.Unmarshal(rateing.Body(), &exchangeRateData)
	if err2 != nil {
		return model.InitializeResponse{}, "rate getting  failed", 0, err2
	}
	ngnExchangeRate := exchangeRateData["rates"].(map[string]interface{})["NGN"].(float64)

	subscriptionRes.Amount = amount * int64(ngnExchangeRate)
	// fmt.Println(subscriptionRes.Amount)

	idHash, _ := primitive.ObjectIDFromHex(id)

	// check if id  exists
	idsearch := map[string]any{
		"_id": idHash,
	}
	idCount, _ := mongodb.MongoCount(constants.UserCollection, idsearch)
	if idCount < 1 {
		return model.InitializeResponse{}, "user does not exist", 403, errors.New("user does not exist in database")
	}

	// check if user has an active sub
	idSubsearch := map[string]any{
		"user_id": idHash,
		"status":  true,
	}

	idSubCount, _ := mongodb.MongoCount(constants.SubscriptionCollection, idSubsearch)
	if idSubCount >= 1 {
		return model.InitializeResponse{}, "You have an active subscription", 403, errors.New("you have an active subscription")
	}

	// get from db
	var resultTwo model.User
	result, err := mongodb.MongoGetOne(constants.UserCollection, idsearch)
	if err != nil {
		return model.InitializeResponse{}, "Unable to get user from database", 500, err
	}
	result.Decode(&resultTwo)
	// get from db end

	id = primitive.NewObjectID().Hex()
	payload := map[string]interface{}{
		"amount":    subscriptionRes.Amount * 100,
		"email":     subscriptionRes.Email,
		"reference": id,
		"currency":  currency,
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

	var subscription model.Subscription
	subscription.Amount = float64(amount) // used USD
	subscription.UserID = idHash
	subscription.Username = resultTwo.Username
	subscription.ID = primitive.NewObjectID()
	subscription.DateCreated = time.Now().Local()
	subscription.DateUpdated = time.Now().Local()
	subscription.Email = subscriptionRes.Email
	subscription.Reference = id
	subscription.AuthorizationUrl = resultOne.Data.AuthorizationUrl
	subscription.AccessCode = resultOne.Data.AccessCode

	// time
	subscription.Type = "premium"
	subscription.Processing = true
	subscription.Status = false
	subscription.Failed = false
	subscription.Duration = days
	subscription.DateExpiring = subscription.DateCreated.Add(time.Hour * hours)

	// save to DB
	_, err = mongodb.MongoPost(constants.SubscriptionCollection, subscription)
	if err != nil {
		return model.InitializeResponse{}, "Unable to save sub to database", 500, err
	}

	return resultOne, "Initiation success", 0, nil
}

func VerifyPaymentByIdService(rid string) (model.PayVerificationResponse, string, int, error) {

	response, err := paystack.PaystackVerifyGet(rid)
	if err != nil {
		return model.PayVerificationResponse{}, "Could not get payment status", response.StatusCode(), err
	}
	// fmt.Println(response)

	// Unmarshal the response body into the struct
	var resultOne model.PayVerificationResponse
	err = json.Unmarshal(response.Body(), &resultOne)
	if err != nil {
		return model.PayVerificationResponse{}, "Could not get payment status", 0, err
	}

	if !response.IsSuccess() {
		return model.PayVerificationResponse{}, resultOne.Message, response.StatusCode(), errors.New("could not get payment verification try again later")
	}

	message, code, err1 := UpdateSubPaystack(rid, resultOne.Data.Status)

	if err1 != nil {
		return model.PayVerificationResponse{}, message, code, err1
	}
	// fmt.Println(resultOne.Status)

	return resultOne, "veification data received", 0, nil
}

func UpdateSubPaystack(rid string, status string) (string, int, error) {

	// update user and sub status
	// check if user has an active sub
	idVersearch := map[string]any{
		"reference": rid,
	}

	// update user
	// get from db
	var resultTwo model.Subscription
	result, err := mongodb.MongoGetOne(constants.SubscriptionCollection, idVersearch)
	if err != nil {
		return "Unable to get subscription from database", 500, err
	}
	result.Decode(&resultTwo)

	if resultTwo.DateExpiring.After(time.Now().Local()) && status == "success" {
		// updating status in sub collection

		resultTwo.Status = true
		resultTwo.Processing = false
		resultTwo.Failed = false

		_, err2 := mongodb.MongoUpdate(idVersearch, resultTwo, constants.SubscriptionCollection)
		if err2 != nil {
			return "Unable to update subscription to database", 500, err
		}

		// update status ends
		// update user
		var user model.User
		user.ID = resultTwo.UserID
		user.Upgrade = true

		idUsersearch := map[string]primitive.ObjectID{
			"_id": user.ID,
		}

		// update user
		// save to DB
		_, err = mongodb.MongoUpdate(idUsersearch, user, constants.UserCollection)
		if err != nil {
			return "Unable to update user to database", 500, err
		}
		// update user ends

	} else if resultTwo.DateExpiring.After(time.Now().Local()) && status == "failed" {
		// updating status in sub collection
		resultTwo.Status = false
		resultTwo.Processing = false
		resultTwo.Failed = true

		_, err2 := mongodb.MongoUpdate(idVersearch, resultTwo, constants.SubscriptionCollection)
		if err2 != nil {
			return "Unable to update subscription to database", 500, err
		}
		// update user
		var user model.User
		user.ID = resultTwo.UserID
		user.Upgrade = false

		idUsersearch := map[string]primitive.ObjectID{
			"_id": user.ID,
		}

		// update user
		// save to DB
		_, err = mongodb.MongoUpdate(idUsersearch, user, constants.UserCollection)
		if err != nil {
			return "Unable to update user to database", 500, err
		}
		// update user ends
	}

	return "Update Success", 200, nil
}

func UpdateSubStatus() (int64, error) {
	searchText := map[string]string{}
	// get from db
	result, err := mongodb.MongoGetAll(constants.SubscriptionCollection, searchText)
	if err != nil {
		return 0, err
	}
	var subscription = make([]model.Subscription, 0)
	result.All(context.TODO(), &subscription)

	var p int64 = 0
	for _, v := range subscription {
		if v.DateExpiring.Before(time.Now().Local()) && v.Status {
			v.Status = false
			idSubsearch := map[string]primitive.ObjectID{
				"_id": v.ID,
			}
			/// update db
			_, err := mongodb.MongoUpdate(idSubsearch, v, constants.SubscriptionCollection)
			if err != nil {
				return 0, err
			}

			idUsersearch := map[string]primitive.ObjectID{
				"_id": v.UserID,
			}
			var m model.User
			m.ID = v.UserID
			m.Upgrade = false
			/// update db
			_, errs := mongodb.MongoUpdate(idUsersearch, m, constants.UserCollection)
			if errs != nil {
				return 0, errs
			}

			p += 1
		} else if v.DateExpiring.After(time.Now().Local()) && !v.Processing && !v.Failed {
			v.Status = true
			idSubsearch := map[string]primitive.ObjectID{
				"_id": v.ID,
			}
			/// update db
			_, err := mongodb.MongoUpdate(idSubsearch, v, constants.SubscriptionCollection)
			if err != nil {
				return 0, err
			}

			idUsersearch := map[string]primitive.ObjectID{
				"_id": v.UserID,
			}
			var m model.User
			m.ID = v.UserID
			m.Upgrade = true
			/// update db
			_, errs := mongodb.MongoUpdate(idUsersearch, m, constants.UserCollection)
			if errs != nil {
				return 0, errs
			}

			p += 1

		}
	}
	return p, nil
}
