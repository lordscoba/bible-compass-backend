package verses

import (
	"encoding/json"
	"fmt"

	"github.com/lordscoba/bible_compass_backend/internal/model"
	gptbible "github.com/lordscoba/bible_compass_backend/pkg/repository/aibible"
)

func AiBibleService(aibible string) (model.Scripture, string, int, error) {

	message, err := gptbible.GptBible(aibible)

	if err != nil {
		return model.Scripture{}, err.Error(), 403, err

	}

	var data map[string]model.Scripture
	err = json.Unmarshal([]byte(message.Body()), &data)

	if err != nil {
		fmt.Println("Error:", err)
		return model.Scripture{}, err.Error(), 403, err
	}

	var Jsondata model.Scripture
	// Access the dynamic key and print the values
	for _, value := range data {
		Jsondata = value
	}

	return Jsondata, "", 0, nil
}
