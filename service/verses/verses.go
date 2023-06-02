package verses

import (
	"time"

	"github.com/lordscoba/bible_compass_backend/internal/model"
)

func AdminCreateVerses(verses model.BibleVerse, id string) (model.BibleVerseResponse, string, int, error) {

	// check if email already exists
	// versesearch := map[string]any{
	// 	"email": verses.BibleVerse,
	// }
	// emailCount, _ := mongodb.MongoCount(constants.UserCollection, emailsearch)
	// if emailCount >= 1 {
	// 	return model.UserResponse{}, "email already exist", 403, errors.New("email already exist in database")
	// }

	verses.DateCreated = time.Now().Local()

	// // save to DB
	// _, err := mongodb.MongoUpdate(constants.UserCollection, user)
	// if err != nil {
	// 	return model.BibleVerseResponse{}, "Unable to save user to database", 500, err
	// }

	verseResponse := model.BibleVerseResponse{
		BibleVerse: verses.BibleVerse,
	}
	return verseResponse, "", 0, nil
}
