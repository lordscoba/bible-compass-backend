package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Subscription struct {
	ID           primitive.ObjectID `bson:"_id, omitempty"`
	Username     string             `bson:"username" json:"username" validate:"required"`
	UserID       primitive.ObjectID `bson:"user_id" json:"user_id"`
	Type         string             `bson:"type" json:"type"` //just premium
	Amount       float64            `bson:"amount" json:"amount"`
	Status       bool               `bson:"status" json:"status"`
	Duration     time.Duration      `bson:"duration" json:"duration"`
	DateCreated  time.Time          `bson:"date_created" json:"date_created"`
	DateExpiring time.Time          `bson:"date_expiring" json:"date_expiring"`
	DateUpdated  time.Time          `bson:"date_updated" json:"date_updated"`
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
