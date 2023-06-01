package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	ID             primitive.ObjectID `bson:"_id, omitempty"`
	CategoryName   string             `bson:"category_name" json:"category_name" validate:"required"`
	ForSubscribers bool               `bson:"for_subscribers" json:"for_subscribers"`
	Keywords       []string           `bson:"keywords" json:"keywords"`
	DateCreated    time.Time          `bson:"date_created" json:"date_created"`
	DateUpdated    time.Time          `bson:"date_updated" json:"date_updated"`
}

type CategoryResponse struct {
	CategoryName   string    `bson:"category_name" json:"category_name"`
	ForSubscribers bool      `bson:"for_subscribers" json:"for_subscribers"`
	Keywords       []string  `bson:"keywords" json:"keywords"`
	DateCreated    time.Time `bson:"date_created" json:"date_created"`
}

type Keywords struct {
	ID             primitive.ObjectID `bson:"_id, omitempty"`
	CategoryID     primitive.ObjectID `bson:"category_id" json:"category_id"`
	Keyword        string             `bson:"keyword" json:"keyword" validate:"required"`
	ForSubscribers bool               `bson:"for_subscribers" json:"for_subscribers"`
	BibleVerse     []BibleVerse       `bson:"bible_verse" json:"bible_verse"`
	DateCreated    time.Time          `bson:"date_created" json:"date_created"`
	DateUpdated    time.Time          `bson:"date_updated" json:"date_updated"`
	Favorite       bool
}

type KeywordsResponse struct {
	Keyword        string       `bson:"keyword" json:"keyword" validate:"required"`
	ForSubscribers bool         `bson:"for_subscribers" json:"for_subscribers"`
	BibleVerse     []BibleVerse `bson:"bible_verse" json:"bible_verse"`
	DateCreated    time.Time    `bson:"date_created" json:"date_created"`
	Favorite       bool         `bson:"favorite" json:"favorite"`
}

type BibleVerse struct {
	ID            primitive.ObjectID `bson:"_id, omitempty"`
	BibleVerse    string             `bson:"bible_verse" json:"bible_verse" validate:"required"`
	RelatedVerses []string           `bson:"related_verses" json:"related_verses"`
	Passage       string             `bson:"passage" json:"passage"`
	Explanation   string             `bson:"explanation" json:"explanation"`
	Like          bool               `bson:"like" json:"like"`
}

type KeywordsInfoResponse struct {
	TotalKeywords       int64
	SubscribersKeywords int64
	TotalBibleVerse     int64
}

type CategoryInfoResponse struct {
	TotalCategory       int64
	SubscribersCategory int64
	TotalKeyWords       int64
}
