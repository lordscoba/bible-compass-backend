package admin

import (
	"errors"
	"time"

	"github.com/lordscoba/bible_compass_backend/internal/constants"
	"github.com/lordscoba/bible_compass_backend/internal/model"
	"github.com/lordscoba/bible_compass_backend/pkg/repository/storage/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func AdminCreateUser(user model.User) (model.UserResponse, string, int, error) {

	// check if email already exists
	emailsearch := map[string]any{
		"email": user.Email,
	}
	emailCount, _ := mongodb.MongoCount(constants.UserCollection, emailsearch)
	if emailCount >= 1 {
		return model.UserResponse{}, "email already exist", 403, errors.New("email already exist in database")
	}

	// check if username already exists
	usernamesearch := map[string]any{
		"username": user.Username,
	}
	usernameCount, _ := mongodb.MongoCount(constants.UserCollection, usernamesearch)
	if usernameCount >= 1 {
		return model.UserResponse{}, "username already exist", 403, errors.New("username already exist in database")
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	user.Password = string(hash)
	user.ID = primitive.NewObjectID()
	user.DateCreated = time.Now().Local()
	user.DateUpdated = time.Now().Local()

	// save to DB
	_, err := mongodb.MongoPost(constants.UserCollection, user)
	if err != nil {
		return model.UserResponse{}, "Unable to save user to database", 500, err
	}

	userResponse := model.UserResponse{
		Username: user.Username,
		Name:     user.Name,
		Email:    user.Email,
	}
	return userResponse, "", 0, nil
}
