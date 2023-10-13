package verses

import (
	"encoding/json"
	"fmt"

	"github.com/lordscoba/bible_compass_backend/internal/model"
	"github.com/lordscoba/bible_compass_backend/pkg/repository/randombible"
)

func RandomBibleService() (model.RandomBibleVerseMain, string, int, error) {

	message, err := randombible.RandomBible()

	if err != nil {
		return model.RandomBibleVerseMain{}, err.Error(), 403, err

	}

	// Parse the trimmed input as JSON
	var data model.RandomBibleVerseMain
	err = json.Unmarshal(message.Body(), &data)

	if err != nil {
		fmt.Println("Error:", err)
		return model.RandomBibleVerseMain{}, err.Error(), 403, err
	}

	return data, "", 0, nil
}
