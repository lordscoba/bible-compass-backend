package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	Username        string             `bson:"username" json:"username"`
	Name            string             `bson:"name" json:"name"`
	Email           string             `bson:"email" json:"email" validate:"required,email"`
	Type            string             `bson:"type" json:"type" `
	Password        string             `bson:"password" json:"password"`
	ConfirmPassword string             `bson:"confirm_password" json:"confirm_password"`
	DateCreated     time.Time          `bson:"date_created" json:"date_created"`
	DateUpdated     time.Time          `bson:"date_updated" json:"date_updated"`
	LastLogin       time.Time          `bson:"last_login" json:"last_login"`
	IsVerified      bool               `bson:"is_verified" json:"is_verified"`
}

type UserResponse struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Type      string    `json:"type"`
	Token     string    `json:"token"`
	TokenType string    `json:"token_type"`
	LastLogin time.Time `json:"last_login"`
}

type UserInfoResponse struct {
	TotalUsers      int64 `json:"total_users"`
	SubscribedUsers int64 `json:"subscribed_users"`
	VerifiedUsers   int64 `json:"verified_users"`
}
