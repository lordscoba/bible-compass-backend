package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Subscription struct {
	ID               primitive.ObjectID `bson:"_id, omitempty"  json:"id"`
	Username         string             `bson:"username" json:"username" `
	Email            string             `bson:"email" json:"email" `
	UserID           primitive.ObjectID `bson:"user_id" json:"user_id"`
	Type             string             `bson:"type" json:"type"` //just premium
	Amount           float64            `bson:"amount" json:"amount"`
	Status           bool               `bson:"status" json:"status"`
	Processing       bool               `json:"processing"`
	Failed           bool               `json:"failed"`
	AuthorizationUrl string             `json:"authorization_url"`
	AccessCode       string             `json:"access_code"`
	Duration         time.Duration      `bson:"duration" json:"duration"`
	DateCreated      time.Time          `bson:"date_created" json:"date_created"`
	DateExpiring     time.Time          `bson:"date_expiring" json:"date_expiring"`
	DateUpdated      time.Time          `bson:"date_updated" json:"date_updated"`
	Reference        string             `bson:"reference" json:"reference" `
}

type SubscriptionResponse struct {
	Username     string        `bson:"username" json:"username" validate:"required"`
	Type         string        `bson:"type" json:"type"` //just premium
	Amount       float64       `bson:"amount" json:"amount"`
	Status       bool          `bson:"status" json:"status"`
	Duration     time.Duration `bson:"duration" json:"duration"`
	DateExpiring time.Time     `bson:"date_expiring" json:"date_expiring"`
}

type SubscriptionInfoResponse struct {
	TotalSubscription    int64 `json:"total_subscription"`
	ActiveSubscription   int64 `json:"active_subscription"`
	InActiveSubscription int64 `json:"inactive_subscription"`
}

type InitializePaymentModel struct {
	Amount    int64  `json:"amount"`
	Email     string `json:"email"`
	Reference string `json:"reference"`
}

type InitializeResponse struct {
	Status  bool                `json:"status"`
	Message string              `json:"message"`
	Data    InitializeResponse2 `json:"data"`
	// Add other fields as needed
}

type InitializeResponse2 struct {
	AuthorizationUrl string `json:"authorization_url"`
	AccessCode       string `json:"access_code"`
	Reference        string `json:"reference"`
}

type PayVerificationResponse struct {
	Status  bool             `json:"status"`
	Message string           `json:"message"`
	Data    VerificationData `json:"data"`
}

type VerificationData struct {
	Amount          int    `json:"amount"`
	Currency        string `json:"currency"`
	Channel         string `json:"channel"`
	GatewayResponse string `json:"gateway_response"`
	Reference       string `json:"reference"`
	Status          string `json:"status"`
}
