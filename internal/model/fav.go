package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Fav struct {
	ID          primitive.ObjectID `bson:"_id, omitempty"  json:"id"`
	Email       string             `bson:"email" json:"email" validate:"required"`
	Keyword     string             `bson:"keyword" json:"keyword" validate:"required"`
	Fav         []FavData          `bson:"fav,truncate" json:"fav"`
	TotalFavs   int                `bson:"total_favs" json:"total_favs"`
	DateCreated time.Time          `bson:"date_created" json:"date_created"`
	DateUpdated time.Time          `bson:"date_updated" json:"date_updated"`
}

type FavData struct {
	ID      string `bson:"id" json:"id"`
	Keyword string `bson:"keyword" json:"keyword" `
}

type FavStatus struct {
	Email   string `json:"email"`
	Keyword string `json:"keyword"`
	Status  bool   `json:"status"`
}
