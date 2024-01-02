package model

type Scripture struct {
	Translation  string   `json:"translation"`
	Abbreviation string   `json:"abbreviation"`
	Lang         string   `json:"lang"`
	Language     string   `json:"language"`
	Direction    string   `json:"direction"`
	Encoding     string   `json:"encoding"`
	BookNr       int      `json:"book_nr"`
	BookName     string   `json:"book_name"`
	Chapter      int      `json:"chapter"`
	Name         string   `json:"name"`
	Ref          []string `json:"ref"`
	Verses       []struct {
		Chapter int    `json:"chapter"`
		Verse   int    `json:"verse"`
		Name    string `json:"name"`
		Text    string `json:"text"`
	} `json:"verses"`
}

// type Scripture struct {
// 	ID        primitive.ObjectID `bson:"_id, omitempty"  json:"id"`
// 	Book      []Book             `json:"book"`
// 	Direction string             `json:"direction"`
// 	Type      string             `json:"type"`
// 	Version   string             `json:"version"`
// }

// type Book struct {
// 	BookRef   string                `json:"book_ref"`
// 	BookName  string                `json:"book_name"`
// 	BookNr    string                `json:"book_nr"`
// 	ChapterNr string                `json:"chapter_nr"`
// 	Chapter   map[string]Verse[any] `json:"chapter"`
// 	// Chapter interface{} `json:"chapter"`
// }

// type Verse[r any] struct {
// 	VerseNr r      `json:"verse_nr"`
// 	Verse   string `json:"verse"`
// }

type RandomBibleVerseMain struct {
	Verse RandomVerse `json:"verse"`
}

type RandomVerse struct {
	Verse  RandomDetails `json:"details"`
	Notice string        `json:"notice"`
}

type RandomDetails struct {
	Text      string `json:"text"`
	Reference string `json:"reference"`
	Version   string `json:"version"`
	VerseURL  string `json:"verseurl"`
}
