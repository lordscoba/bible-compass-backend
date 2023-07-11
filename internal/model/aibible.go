package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Scripture struct {
	ID        primitive.ObjectID `bson:"_id, omitempty"  json:"id"`
	Book      []Book             `json:"book"`
	Direction string             `json:"direction"`
	Type      string             `json:"type"`
	Version   string             `json:"version"`
}

type Book struct {
	BookRef   string                `json:"book_ref"`
	BookName  string                `json:"book_name"`
	BookNr    string                `json:"book_nr"`
	ChapterNr string                `json:"chapter_nr"`
	Chapter   map[string]Verse[any] `json:"chapter"`
	// Chapter interface{} `json:"chapter"`
}

type Verse[r any] struct {
	VerseNr r      `json:"verse_nr"`
	Verse   string `json:"verse"`
}
