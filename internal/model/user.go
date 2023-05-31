package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id, omitempty"`
	Username    string             `bson:"username" json:"username" validate:"required"`
	Name        string             `bson:"name" json:"name"`
	Email       string             `bson:"email" json:"email" validate:"required,email"`
	Type        string             `bson:"type" json:"type" `
	Password    string             `bson:"password" json:"password"`
	DateCreated time.Time          `bson:"date_created" json:"date_created"`
	DateUpdated time.Time          `bson:"date_updated" json:"date_updated"`
	LastLogin   time.Time          `bson:"last_login" json:"last_login"`
	IsVerified  bool               `bson:"is_verified" json:"is_verified"`
}

type UserResponse struct {
	Username  string
	Name      string
	Email     string
	Token     string
	TokenType string
	LastLogin time.Time
}

type UserInfoResponse struct {
	TotalUsers      int64
	SubscribedUsers int64
	VerifiedUsers   int64
}
