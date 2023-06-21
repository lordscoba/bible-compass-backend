package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	ID             primitive.ObjectID `bson:"_id, omitempty"  json:"id"`
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
	ID             primitive.ObjectID `bson:"_id, omitempty"  json:"id"`
	CategoryID     primitive.ObjectID `bson:"category_id" json:"category_id"`
	Keyword        string             `bson:"keyword" json:"keyword" validate:"required"`
	ForSubscribers bool               `bson:"for_subscribers" json:"for_subscribers"`
	BibleVerse     []BibleVerse       `bson:"bible_verse,truncate" json:"bible_verse"`
	DateCreated    time.Time          `bson:"date_created" json:"date_created"`
	DateUpdated    time.Time          `bson:"date_updated" json:"date_updated"`
	Favorite       bool               `bson:"favorite" json:"favorite"`
}

type KeywordsResponse struct {
	Keyword        string       `json:"keyword"`
	ForSubscribers bool         `json:"for_subscribers"`
	BibleVerse     []BibleVerse `json:"bible_verse"`
	DateCreated    time.Time    `json:"date_created"`
	Favorite       bool         `json:"favorite"`
}

type BibleVerse struct {
	ID            primitive.ObjectID `bson:"_id, omitempty"  json:"id"`
	BibleVerse    string             `bson:"bible_verse" json:"bible_verse" validate:"required"`
	RelatedVerses []string           `bson:"related_verses" json:"related_verses"`
	Passage       string             `bson:"passage" json:"passage"`
	Explanation   string             `bson:"explanation" json:"explanation"`
	Like          bool               `bson:"like" json:"like"`
	DateCreated   time.Time          `bson:"date_created" json:"date_created"`
}

type BibleVerseResponse struct {
	BibleVerse    string   `json:"bible_verse"`
	RelatedVerses []string `json:"related_verse"`
	Passage       string   `json:"passage"`
	Explanation   string   `json:"explanation"`
	Like          bool     `json:"like"`
}

type BibleVerseInfoResponse struct {
	TotalVerses int `json:"total_verses"`
}

type KeywordsInfoResponse struct {
	TotalKeywords       int64 `json:"total_keywords"`
	SubscribersKeywords int64 `json:"subscribers_keywords"`
	TotalBibleVerse     int64 `json:"total_bible_verse"`
}

type CategoryInfoResponse struct {
	TotalCategory       int64 `json:"total_category"`
	SubscribersCategory int64 `json:"subscribers_category"`
	TotalKeyWords       int64 `json:"total_keywords"`
}
