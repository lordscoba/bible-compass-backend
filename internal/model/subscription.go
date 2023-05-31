package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Subscription struct {
	ID           primitive.ObjectID `bson:"_id, omitempty"`
	Type         string             `bson:"passage" json:"passage"` //just premium
	Amount       float64            `bson:"amount" json:"amount"`
	Status       bool               `bson:"status" json:"status"`
	Duration     time.Duration      `bson:"duration" json:"duration"`
	DateCreated  time.Time          `bson:"date_created" json:"date_created"`
	DateExpiring time.Time          `bson:"date_expiring" json:"date_expiring"`
	DateUpdated  time.Time          `bson:"date_updated" json:"date_updated"`
}
