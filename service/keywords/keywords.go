package keywords

import (
	"errors"
	"time"

	"github.com/lordscoba/bible_compass_backend/internal/constants"
	"github.com/lordscoba/bible_compass_backend/internal/model"
	"github.com/lordscoba/bible_compass_backend/pkg/repository/storage/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AdminCreateKeywords(keywords model.Keywords, id string) (model.KeywordsResponse, string, int, error) {

	idHash, _ := primitive.ObjectIDFromHex(id)

	// check if id  exists in category
	idsearch := map[string]any{
		"_id": idHash,
	}
	idCount, _ := mongodb.MongoCount(constants.CategoryCollection, idsearch)
	if idCount < 1 {
		return model.KeywordsResponse{}, "category does not exist", 403, errors.New("user does not exist in database")
	}

	keywords.ForSubscribers = false
	keywords.Favorite = false
	keywords.CategoryID = idHash
	keywords.ID = primitive.NewObjectID()
	keywords.DateCreated = time.Now().Local()
	keywords.DateUpdated = time.Now().Local()

	// save to DB
	_, err := mongodb.MongoPost(constants.KeywordCollection, keywords)
	if err != nil {
		return model.KeywordsResponse{}, "Unable to save keyword to database", 500, err
	}

	KeywordsResponse := model.KeywordsResponse{
		Keyword:        keywords.Keyword,
		DateCreated:    keywords.DateCreated,
		ForSubscribers: keywords.ForSubscribers,
		Favorite:       keywords.Favorite,
	}
	return KeywordsResponse, "", 0, nil
}

func AdminUpdateKeywords(keywords model.Keywords, id string) (model.KeywordsResponse, string, int, error) {

	idHash, _ := primitive.ObjectIDFromHex(id)

	idsearch := map[string]any{
		"_id": idHash,
	}
	// check if id exists
	idCount, _ := mongodb.MongoCount(constants.KeywordCollection, idsearch)
	if idCount < 1 {
		return model.KeywordsResponse{}, "keyword does not exist", 403, errors.New("keyword does not exist in database")
	}
	// get from db
	var resultOne model.Keywords
	result, err := mongodb.MongoGetOne(constants.KeywordCollection, idsearch)
	if err != nil {
		return model.KeywordsResponse{}, "Unable to get user from database", 500, err
	}
	result.Decode(&resultOne)
	// get from db end

	keywords.CategoryID = resultOne.CategoryID
	keywords.DateUpdated = time.Now().Local()

	// save to DB
	_, err = mongodb.MongoUpdate(idsearch, keywords, constants.KeywordCollection)
	if err != nil {
		return model.KeywordsResponse{}, "Unable to save user to database", 500, err
	}

	KeywordsResponse := model.KeywordsResponse{
		Keyword:        keywords.Keyword,
		DateCreated:    keywords.DateCreated,
		ForSubscribers: keywords.ForSubscribers,
		Favorite:       keywords.Favorite,
	}
	return KeywordsResponse, "", 0, nil
}
