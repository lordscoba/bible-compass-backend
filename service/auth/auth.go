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

	// check if required data is entered
	if user.Username == "" {
		return model.UserResponse{}, "Enter Username", 403, errors.New("username is missing")
	}
	if user.Email == "" {
		return model.UserResponse{}, "Enter Email", 403, errors.New("email is missing")
	}
	if user.Password == "" {
		return model.UserResponse{}, "Enter password", 403, errors.New("password is missing")
	}
	if user.ConfirmPassword == "" {
		return model.UserResponse{}, "Enter Confirm password", 403, errors.New("confirm password is missing")
	}

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

	// check if passwords match
	if user.Password != user.ConfirmPassword {
		return model.UserResponse{}, "Passwords does not match", 403, errors.New("passwords does not match")
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	user.Password = string(hash)
	user.ID = primitive.NewObjectID()
	user.DateCreated = time.Now().Local()
	user.DateUpdated = time.Now().Local()
	user.Type = "user"
	user.ConfirmPassword = ""

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
	if user.Email == "" {
		return model.UserResponse{}, "Enter Email", 403, errors.New("Email is missing")
	}
	if user.Password == "" {
		return model.UserResponse{}, "Enter Password", 403, errors.New("password is missing")
	}

	// check if user exists
	emailsearch := map[string]any{
		"email": user.Email,
	}
	emailCount, _ := mongodb.MongoCount(constants.UserCollection, emailsearch)

	if emailCount < 1 {
		return model.UserResponse{}, "email does not exist", 403, errors.New("email does not exist")
	}

	// get from db
	result, err := mongodb.MongoGet(constants.UserCollection, emailsearch)
	if err != nil {
		return model.UserResponse{}, "Unable to get user to database", 500, err
	}

	var users = make([]model.User, 0)
	result.All(context.TODO(), &users)

	var saved model.User
	for i, userSaved := range users {
		if i == 0 {
			if !utility.IsValidPassword(userSaved.Password, user.Password) {
				return model.UserResponse{}, "password is wrong", 403, errors.New("password is wrong")

			}
			saved.Username = userSaved.Username
			saved.Name = userSaved.Name
			saved.Type = userSaved.Type

		}
	}

	user.DateUpdated = time.Now().Local()
	user.LastLogin = time.Now().Local()
	user.Password = ""

	// save to DB
	_, err = mongodb.MongoUpdate(emailsearch, user, constants.UserCollection)
	if err != nil {
		return model.UserResponse{}, "Unable to save user to database", 500, err
	}

	userResponse := model.UserResponse{
		Username:  saved.Username,
		Email:     user.Email,
		Type:      saved.Type,
		TokenType: "bearer",
		Token:     "",
		LastLogin: user.LastLogin,
	}

	return userResponse, "", 0, nil
}
