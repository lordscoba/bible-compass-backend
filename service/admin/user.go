package admin

import (
	"context"
	"errors"
	"fmt"
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

func AdminUpdateUser(user model.User, id string) (model.UserResponse, string, int, error) {

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
		return model.UserResponse{}, "Unable to save user to database", 500, err
	}

	userResponse := model.UserResponse{
		Username: user.Username,
		Name:     user.Name,
		Email:    user.Email,
	}
	return userResponse, "", 0, nil
}

func AdminGetUser() ([]model.User, string, int, error) {

	// get from db
	result, err := mongodb.MongoGetAll(constants.UserCollection)
	if err != nil {
		return []model.User{}, "Unable to save user to database", 500, err
	}

	var users = make([]model.User, 0)
	result.All(context.TODO(), &users)
	return users, "", 0, nil
}

func AdminGetUserbyId(id string) ([]model.User, string, int, error) {
	idHash, _ := primitive.ObjectIDFromHex(id)
	search := map[string]any{
		"_id": idHash,
	}
	// get from db
	result, err := mongodb.MongoGet(constants.UserCollection, search)
	if err != nil {
		return []model.User{}, "Unable to save user to database", 500, err
	}

	var users = make([]model.User, 0)
	result.All(context.TODO(), &users)

	fmt.Println(users)
	return users, "", 0, nil
}

func AdminDeleteUserbyId(id string) (int64, string, int, error) {

	idHash, _ := primitive.ObjectIDFromHex(id)
	search := map[string]any{
		"_id": idHash,
	}

	// check if id exists
	idCount, _ := mongodb.MongoCount(constants.UserCollection, search)
	if idCount < 1 {
		return 0, "ID does not exist", 403, errors.New("ID does not exist in database")
	}

	// get from db
	result, err := mongodb.MongoDelete(constants.UserCollection, search)
	if err != nil {
		return 0, "Unable to save user to database", 500, err
	}

	fmt.Println(result.DeletedCount)
	return result.DeletedCount, "", 0, nil
}
