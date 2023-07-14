package verses

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/lordscoba/bible_compass_backend/internal/model"
	gptbible "github.com/lordscoba/bible_compass_backend/pkg/repository/aibible"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AiBibleService(aibible string) (model.Scripture, string, int, error) {

	message, err := gptbible.GptBible(aibible)

	// Remove surrounding parentheses
	trimmedInput := strings.TrimPrefix(strings.TrimSuffix(string(message.Body()), ");"), "(")

	if err != nil {
		return model.Scripture{}, err.Error(), 403, err

	}

	// Parse the trimmed input as JSON
	var data model.Scripture
	data.ID = primitive.NewObjectID()
	err = json.Unmarshal([]byte(trimmedInput), &data)

	if err != nil {
		fmt.Println("Error:", err)
		return model.Scripture{}, err.Error(), 403, err
	}

	return data, "", 0, nil
}

// func AiBibleService(aibible string) (map[string]interface{}, string, int, error) {

// 	message, err := gptbible.GptBible(aibible)

// 	// Remove surrounding parentheses
// 	trimmedInput := strings.TrimPrefix(strings.TrimSuffix(string(message.Body()), ");"), "(")

// 	if err != nil {
// 		return map[string]interface{}{}, err.Error(), 403, err

// 	}

// 	// Parse the trimmed input as JSON
// 	var data map[string]interface{}
// 	err = json.Unmarshal([]byte(trimmedInput), &data)

// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return map[string]interface{}{}, "get bible failed  failed", 0, err
// 	}

// 	return data, "", 0, nil
// }
