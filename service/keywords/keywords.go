package keywords

import (
	"context"
	"errors"
	"fmt"
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

	// check if key exists
	keysearch := map[string]any{
		"keyword": keywords.Keyword,
	}
	keyCount, _ := mongodb.MongoCount(constants.KeywordCollection, keysearch)
	if keyCount >= 1 {
		return model.KeywordsResponse{}, "key already exist", 403, errors.New("key already exist in database")
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

	// save keyword to category
	// get from db
	var resultOne model.Category
	result, err := mongodb.MongoGetOne(constants.CategoryCollection, idsearch)
	if err != nil {
		return model.KeywordsResponse{}, "Unable to get category from database", 500, err
	}
	result.Decode(&resultOne)

	var category model.Category
	category.Keywords = append(resultOne.Keywords, keywords.Keyword)

	// save to DB
	_, err = mongodb.MongoUpdate(idsearch, category, constants.CategoryCollection)
	if err != nil {
		return model.KeywordsResponse{}, "Unable to save user to database", 500, err
	}
	// save keyword to category ends

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

func AdminGetKeywords() ([]model.Keywords, string, int, error) {

	// get from db
	result, err := mongodb.MongoGetAll(constants.KeywordCollection)
	if err != nil {
		return []model.Keywords{}, "Unable to save keywords to database", 500, err
	}

	var keywords = make([]model.Keywords, 0)
	result.All(context.TODO(), &keywords)
	return keywords, "", 0, nil
}

func AdminGetkeywordsbyId(id string) (model.Keywords, string, int, error) {
	idHash, _ := primitive.ObjectIDFromHex(id)
	search := map[string]any{
		"_id": idHash,
	}
	// get from db
	var resultOne model.Keywords
	result, err := mongodb.MongoGetOne(constants.KeywordCollection, search)
	if err != nil {
		return model.Keywords{}, "Unable to get keywords from database", 500, err
	}
	result.Decode(&resultOne)
	// get from db end

	fmt.Println(resultOne)
	return resultOne, "", 0, nil
}

func AdminDeletekeywordsbyId(id string) (int64, string, int, error) {

	idHash, _ := primitive.ObjectIDFromHex(id)
	search := map[string]any{
		"_id": idHash,
	}

	// check if id exists
	idCount, _ := mongodb.MongoCount(constants.KeywordCollection, search)
	if idCount < 1 {
		return 0, "keywords does not exist", 403, errors.New("keywords does not exist in database")
	}

	// get from db
	result, err := mongodb.MongoDelete(constants.KeywordCollection, search)
	if err != nil {
		return 0, "Unable to save keywords to database", 500, err
	}

	fmt.Println(result.DeletedCount)
	return result.DeletedCount, "", 0, nil
}

func AdminKeywordsInfo() (model.KeywordsInfoResponse, string, int, error) {
	// total users
	search := map[string]any{}
	TotalKeywords, err := mongodb.MongoCount(constants.KeywordCollection, search)
	if err != nil {
		return model.KeywordsInfoResponse{}, "Unable to get count", 500, err
	}

	// total subscribers category
	StatusSearch1 := map[string]any{
		"for_subscribers": true,
	}
	SubscribersKeywords, err := mongodb.MongoCount(constants.KeywordCollection, StatusSearch1)
	if err != nil {
		return model.KeywordsInfoResponse{}, "Unable to get count", 500, err
	}

	// total keywords
	StatusSearch2 := map[string]any{
		"status": false,
	}
	TotalBibleVerse, err := mongodb.MongoCount(constants.KeywordCollection, StatusSearch2)
	if err != nil {
		return model.KeywordsInfoResponse{}, "Unable to get count", 500, err
	}

	KeywordsInfo := model.KeywordsInfoResponse{
		TotalKeywords:       TotalKeywords,
		SubscribersKeywords: SubscribersKeywords,
		TotalBibleVerse:     TotalBibleVerse,
	}

	return KeywordsInfo, "", 0, nil
}
