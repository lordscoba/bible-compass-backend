package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/lordscoba/bible_compass_backend/internal/config"
	"github.com/lordscoba/bible_compass_backend/internal/constants"
	"github.com/lordscoba/bible_compass_backend/internal/model"
	"github.com/lordscoba/bible_compass_backend/pkg/repository/mailer"
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
	user.Upgrade = false

	// save to DB
	_, err := mongodb.MongoPost(constants.UserCollection, user)
	if err != nil {
		return model.UserResponse{}, "Unable to save user to database", 500, err
	}

	emailData := utility.EmailData{
		Email: user.Email,
	}

	emailContent, err := utility.GenerateRegistrationHTMLEmail(emailData)
	if err != nil {
		fmt.Println("Error generating email content:", err)
		// return model.UserResponse{}, "Unable Generate Email", 403, err
	}

	err = mailer.SendMail(user.Email, emailContent)
	if err != nil {
		fmt.Println("Error generating email content:", err)
		// return model.UserResponse{}, "Unable to send email", 403, err
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
		return model.UserResponse{}, "Enter Email", 403, errors.New("email is missing")
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
	searchText := map[string]string{
		"username": "",
	}
	result, err := mongodb.MongoGet(constants.UserCollection, emailsearch, searchText)
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
			saved.ID = userSaved.ID
			saved.Upgrade = userSaved.Upgrade

		}
	}

	secretkey := config.GetConfig().Server.Secret

	token, err2 := utility.GenerateAllTokens(secretkey, user.Email, saved.Username, saved.Type, saved.ID.Hex())
	if err2 != nil {
		return model.UserResponse{}, "Unable to generate token", 500, err
	}

	user.Token = token
	user.TokenType = "bearer"
	user.DateUpdated = time.Now().Local()
	user.LastLogin = time.Now().Local()
	user.Password = ""

	// save to DB
	_, err = mongodb.MongoUpdate(emailsearch, user, constants.UserCollection)
	if err != nil {
		return model.UserResponse{}, "Unable to save user to database", 500, err
	}

	userResponse := model.UserResponse{
		Upgrade:   saved.Upgrade,
		Name:      saved.Name,
		ID:        saved.ID.Hex(),
		Username:  saved.Username,
		Email:     user.Email,
		Type:      saved.Type,
		TokenType: user.TokenType,
		Token:     token,
		LastLogin: user.LastLogin,
	}

	return userResponse, "", 0, nil
}

func VerificationService(user model.VerifyModel) (model.VerifyModel, string, int, error) {

	// check for input data
	if user.Email == "" {
		return model.VerifyModel{}, "Enter Email", 403, errors.New("email is missing")
	}

	// check if email exists
	emailsearch := map[string]any{
		"email": user.Email,
	}
	emailCount, _ := mongodb.MongoCount(constants.UserCollection, emailsearch)
	if emailCount < 1 {
		return model.VerifyModel{}, "email is not registered", 403, errors.New("email does not exist in database")
	}

	// get user details
	var resultOne model.User
	result, err := mongodb.MongoGetOne(constants.UserCollection, emailsearch)
	if err != nil {
		return model.VerifyModel{}, "Unable to get users details from database", 500, err
	}
	result.Decode(&resultOne)
	// get from db end

	// update db token
	// generatedCode := primitive.NewObjectID().Hex()
	// fmt.Println(emailsearch)
	// fmt.Println(generatedCode)

	resultOne.VerificationCode = primitive.NewObjectID().Hex()
	resultOne.VerificationTime = time.Now().Local()
	fmt.Println(resultOne)

	// save to DB
	_, err = mongodb.MongoUpdate(emailsearch, resultOne, constants.UserCollection)
	if err != nil {
		return model.VerifyModel{}, "Unable to save user to database", 500, err
	}

	host := "https://www.bible-compass.com/verify"
	vid := "?verification_id=" + resultOne.VerificationCode
	vemail := "&email=" + user.Email
	vuserid := "&user_id=" + resultOne.ID.Hex()
	vlink := host + vid + vemail + vuserid

	fmt.Println(vlink)

	emailData := utility.EmailData{
		Name:       resultOne.Name,
		Email:      user.Email,
		Link:       vlink,
		ExpiryTime: "1 hr",
	}

	_, err3 := utility.GenerateHTMLEmail(emailData)
	if err3 != nil {

		// fmt.Println("Error generating email content:", err)
		return model.VerifyModel{}, "Unable Generate Email", 403, err3
	}

	// err = mailer.SendMail("e2scoba2tm@gmail.com", emailContent)

	// if err != nil {

	// 	// fmt.Println("Error generating email content:", err)
	// 	return model.VerifyModel{}, "Unable to send email", 403, err
	// }

	VericationResponse := model.VerifyModel{
		Email: user.Email,
	}

	return VericationResponse, "", 0, nil
}

func ChangePasswordService(user model.ChangePassword) (model.ChangePassword, string, int, error) {

	// check if required data is entered
	if user.Email == "" {
		return model.ChangePassword{}, "Enter Email", 403, errors.New("email is missing")
	}
	if user.Password == "" {
		return model.ChangePassword{}, "Enter Password", 403, errors.New("password is missing")
	}
	if user.ConfirmPassword == "" {
		return model.ChangePassword{}, "Enter Password", 403, errors.New("change password is missing")
	}

	// check if id exists
	idHash, _ := primitive.ObjectIDFromHex(user.UserId)
	idsearch := map[string]any{
		"_id": idHash,
	}
	idCount, _ := mongodb.MongoCount(constants.UserCollection, idsearch)

	if idCount < 1 {
		return model.ChangePassword{}, "id does not exist", 403, errors.New("id does not exist")
	}

	// check if verification exist
	if user.VerificationCode == "" {
		return model.ChangePassword{}, "verification code is missing", 403, errors.New("verification code is missing")
	}

	vcodesearch := map[string]any{
		"verification_code": user.VerificationCode,
	}
	vcodeCount, _ := mongodb.MongoCount(constants.UserCollection, vcodesearch)

	if vcodeCount < 1 {
		return model.ChangePassword{}, "verfication code is not available", 403, errors.New("verfication code is not available")
	}

	// get user details
	var resultOne model.User
	result, err := mongodb.MongoGetOne(constants.UserCollection, vcodesearch)
	if err != nil {
		return model.ChangePassword{}, "Unable to get users details from database", 500, err
	}
	result.Decode(&resultOne)
	// get from db end

	// check if verification code  has expired
	verificationExpiring := resultOne.VerificationTime.Add(time.Hour * 1)
	// fmt.Println(verificationExpiring.Local())
	// fmt.Println(resultOne.VerificationTime.Local())
	// fmt.Println(time.Now().Local())

	if verificationExpiring.Local().Before(time.Now().Local()) {
		return model.ChangePassword{}, "verfication link has expired", 403, errors.New("verfication link has expired")
	}

	// check if passwords match
	if user.Password != user.ConfirmPassword {
		return model.ChangePassword{}, "Passwords does not match", 403, errors.New("passwords does not match")
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	resultOne.Password = string(hash)
	resultOne.ID = idHash
	resultOne.DateUpdated = time.Now().Local()
	// resultOne.VerificationCode = ""

	// save to DB
	_, err = mongodb.MongoUpdate(vcodesearch, resultOne, constants.UserCollection)
	if err != nil {
		return model.ChangePassword{}, "Unable to save user to database", 500, err
	}

	// fmt.Println(user)

	return user, "", 0, nil
}
