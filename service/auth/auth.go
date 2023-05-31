package auth

import (
	"context"
	"errors"
	"time"

	"github.com/lordscoba/bible_compass_backend/internal/constants"
	"github.com/lordscoba/bible_compass_backend/internal/model"
	"github.com/lordscoba/bible_compass_backend/pkg/repository/storage/mongodb"
	"github.com/lordscoba/bible_compass_backend/utility"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func AuthSignUp(user model.User) (model.UserResponse, string, int, error) {

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

func AuthLogin(user model.User) (model.UserResponse, string, int, error) {

	// check if required data is entered
	if user.Username == "" {
		return model.UserResponse{}, "Enter Username", 403, errors.New("username is missing")
	}
	if user.Password == "" {
		return model.UserResponse{}, "Enter Password", 403, errors.New("password is missing")
	}

	// check if user exists
	usernamesearch := map[string]any{
		"username": user.Username,
	}
	usernameCount, _ := mongodb.MongoCount(constants.UserCollection, usernamesearch)

	if usernameCount < 1 {
		return model.UserResponse{}, "username does not exist", 403, errors.New("username does not exist")
	}

	// get from db
	result, err := mongodb.MongoGet(constants.UserCollection, usernamesearch)
	if err != nil {
		return model.UserResponse{}, "Unable to get user to database", 500, err
	}

	var users = make([]model.User, 0)
	result.All(context.TODO(), &users)

	for i, userSaved := range users {
		if i == 0 {
			if !utility.IsValidPassword(userSaved.Password, user.Password) {
				return model.UserResponse{}, "password does not match", 403, errors.New("password does not match")

			}
		}
	}

	user.DateUpdated = time.Now().Local()
	user.LastLogin = time.Now().Local()
	user.Password = ""

	// save to DB
	_, err = mongodb.MongoUpdate(usernamesearch, user, constants.UserCollection)
	if err != nil {
		return model.UserResponse{}, "Unable to save user to database", 500, err
	}

	userResponse := model.UserResponse{
		Username: user.Username,
		Email:    user.Email,
		// TokenType:    "bearer",
		// Token:        token,
		LastLogin: user.LastLogin,
	}

	return userResponse, "", 0, nil
}
