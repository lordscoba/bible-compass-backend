package profile

import (
	"errors"
	"time"

	"github.com/lordscoba/bible_compass_backend/internal/constants"
	"github.com/lordscoba/bible_compass_backend/internal/model"
	"github.com/lordscoba/bible_compass_backend/pkg/repository/storage/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func UpdateProfile(user model.User, id string) (model.UserResponse, string, int, error) {

	idHash, _ := primitive.ObjectIDFromHex(id)

	if user.Password != "" {
		hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
		user.Password = string(hash)
	}

	user.DateUpdated = time.Now().Local()

	idsearch := map[string]any{
		"_id": idHash,
	}

	// check if id exists
	idCount, _ := mongodb.MongoCount(constants.UserCollection, idsearch)
	if idCount < 1 {
		return model.UserResponse{}, "ID does not exist", 403, errors.New("ID does not exist in database")
	}

	// save to DB
	_, err := mongodb.MongoUpdate(idsearch, user, constants.UserCollection)
	if err != nil {
		return model.UserResponse{}, "Unable to save profile to database", 500, err
	}

	userResponse := model.UserResponse{
		Username: user.Username,
		Name:     user.Name,
		Email:    user.Email,
	}
	return userResponse, "", 0, nil
}

func GetProfileDetails(id string) (model.User, string, int, error) {
	idHash, _ := primitive.ObjectIDFromHex(id)
	search := map[string]any{
		"_id": idHash,
	}

	// get from db
	var resultOne model.User
	result, err := mongodb.MongoGetOne(constants.UserCollection, search)

	if err != nil {
		return model.User{}, "Unable to get user to database", 500, err
	}

	result.Decode(&resultOne)
	return resultOne, "", 0, nil
}
