package verses

import (
	"errors"
	"fmt"
	"time"

	"github.com/lordscoba/bible_compass_backend/internal/constants"
	"github.com/lordscoba/bible_compass_backend/internal/model"
	"github.com/lordscoba/bible_compass_backend/pkg/repository/storage/mongodb"
	"github.com/lordscoba/bible_compass_backend/utility"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AdminCreateVerses(verses model.BibleVerse, kid string) (model.BibleVerseResponse, string, int, error) {

	kidHash, _ := primitive.ObjectIDFromHex(kid)
	var keyword model.Keywords

	// check if keyword  exists
	keywordsearch := map[string]any{
		"_id": kidHash,
	}

	keywordCount, _ := mongodb.MongoCount(constants.KeywordCollection, keywordsearch)
	if keywordCount < 1 {
		return model.BibleVerseResponse{}, "keyword doesnt exist", 403, errors.New("keyword doesnt exist in database")
	}

	// check if bible verse already exists
	// get from db
	var resultOne model.Keywords
	result, err := mongodb.MongoGetOne(constants.KeywordCollection, keywordsearch)
	if err != nil {
		return model.BibleVerseResponse{}, "Unable to get keywords from database", 500, err
	}
	result.Decode(&resultOne)
	// get from db end

	// filter array
	for _, v := range resultOne.BibleVerse {
		// fmt.Println(k, " here ", v.BibleVerse)
		if v.BibleVerse == verses.BibleVerse {
			return model.BibleVerseResponse{}, "Bible verse exist", 403, errors.New("bible verse exist in database")
		}
	}

	// append to bible verse to keyword
	bibleDetails := model.BibleVerse{
		ID:            primitive.NewObjectID(),
		BibleVerse:    verses.BibleVerse,
		RelatedVerses: verses.RelatedVerses,
		Passage:       verses.Passage,
		Explanation:   verses.Explanation,
		DateCreated:   time.Now().Local(),
	}

	keyword.BibleVerse = append(resultOne.BibleVerse, bibleDetails)

	// save to DB
	_, err = mongodb.MongoUpdate(keywordsearch, keyword, constants.KeywordCollection)
	if err != nil {
		return model.BibleVerseResponse{}, "Unable to save user to database", 500, err
	}

	// initiate response
	verseResponse := model.BibleVerseResponse{
		BibleVerse:  verses.BibleVerse,
		Passage:     verses.Passage,
		Like:        verses.Like,
		Explanation: verses.Explanation,
	}
	return verseResponse, "", 0, nil
}

func AdminUpdateVerse(verses model.BibleVerse, kid string, Bid string) (model.BibleVerseResponse, string, int, error) {
	kidHash, _ := primitive.ObjectIDFromHex(kid)
	BidHash, _ := primitive.ObjectIDFromHex(Bid)

	// check if keyword  exists
	keywordsearch := map[string]any{
		"_id": kidHash,
	}

	keywordCount, _ := mongodb.MongoCount(constants.KeywordCollection, keywordsearch)
	if keywordCount < 1 {
		return model.BibleVerseResponse{}, "keyword doesnt exist", 403, errors.New("keyword doesnt exist in database")
	}

	// get from db
	var resultOne model.Keywords
	result, err := mongodb.MongoGetOne(constants.KeywordCollection, keywordsearch)
	if err != nil {
		return model.BibleVerseResponse{}, "Unable to get keywords from database", 500, err
	}
	result.Decode(&resultOne)
	// get from db end

	// filter array to check if bible verse exists
	found := false
	for _, v := range resultOne.BibleVerse {
		if v.ID == BidHash {
			found = true
			break
		}
	}
	if !found {
		return model.BibleVerseResponse{}, "Bible verse does not exist", 403, errors.New("bible verser does not exist in database")
	}

	// filter array and assign new values
	var newBibleverse string
	var newRelatedverse []string
	var newPassage string
	var newExplanation string
	var newDateCreated time.Time
	var newLike bool
	var index int
	for n, v := range resultOne.BibleVerse {
		if v.ID == BidHash {
			newBibleverse = utility.ComparingUpdate[string](verses.BibleVerse, v.BibleVerse)
			newRelatedverse = utility.ComparingUpdate[[]string](verses.RelatedVerses, v.RelatedVerses)
			newPassage = utility.ComparingUpdate[string](verses.Passage, v.Passage)
			newExplanation = utility.ComparingUpdate[string](verses.Explanation, v.Explanation)
			newDateCreated = utility.ComparingUpdate[time.Time](verses.DateCreated, v.DateCreated)
			newLike = utility.ComparingUpdate[bool](verses.Like, v.Like)
			index = n
		}
	}

	// new bible verse details to keyword
	bibleDetails := model.BibleVerse{
		ID:            BidHash,
		BibleVerse:    newBibleverse,
		RelatedVerses: newRelatedverse,
		Passage:       newPassage,
		Explanation:   newExplanation,
		DateCreated:   newDateCreated,
		Like:          newLike,
	}

	var updateOne model.Keywords
	updateOne.ID = resultOne.ID
	updateOne.CategoryID = resultOne.CategoryID

	//deleting from slice in mongodb
	updateOne.BibleVerse = utility.DeleteElement(resultOne.BibleVerse, index)
	fmt.Println(updateOne.BibleVerse)
	// save to DB
	_, err = mongodb.MongoUpdate(keywordsearch, updateOne, constants.KeywordCollection)
	if err != nil {
		return model.BibleVerseResponse{}, "Unable to save user to database", 500, err
	}

	//appending to slice in mongodb
	updateOne.BibleVerse = append(updateOne.BibleVerse, bibleDetails)
	// save to DB
	_, err = mongodb.MongoUpdate(keywordsearch, updateOne, constants.KeywordCollection)
	if err != nil {
		return model.BibleVerseResponse{}, "Unable to save user to database", 500, err
	}

	// initiating response
	verseResponse := model.BibleVerseResponse{
		BibleVerse:  newBibleverse,
		Passage:     newPassage,
		Like:        newLike,
		Explanation: newExplanation,
	}
	return verseResponse, "", 0, nil
}

func AdminGetVerse(kid string) ([]model.BibleVerse, string, int, error) {
	kidHash, _ := primitive.ObjectIDFromHex(kid)

	// check if keyword  exists
	keywordsearch := map[string]any{
		"_id": kidHash,
	}

	keywordCount, _ := mongodb.MongoCount(constants.KeywordCollection, keywordsearch)
	if keywordCount < 1 {
		return []model.BibleVerse{}, "keyword doesnt exist", 403, errors.New("keyword doesnt exist in database")
	}

	// get from db
	var resultOne model.Keywords
	result, err := mongodb.MongoGetOne(constants.KeywordCollection, keywordsearch)
	if err != nil {
		return []model.BibleVerse{}, "Unable to get keywords from database", 500, err
	}
	result.Decode(&resultOne)

	allverses := resultOne.BibleVerse

	return allverses, "", 0, nil
}

func AdminGetVersebyId(kid, Bid string) (model.BibleVerse, string, int, error) {
	kidHash, _ := primitive.ObjectIDFromHex(kid)
	BidHash, _ := primitive.ObjectIDFromHex(Bid)

	// check if keyword  exists
	keywordsearch := map[string]any{
		"_id": kidHash,
	}

	keywordCount, _ := mongodb.MongoCount(constants.KeywordCollection, keywordsearch)
	if keywordCount < 1 {
		return model.BibleVerse{}, "keyword doesnt exist", 403, errors.New("keyword doesnt exist in database")
	}

	// get from db
	var resultOne model.Keywords
	result, err := mongodb.MongoGetOne(constants.KeywordCollection, keywordsearch)
	if err != nil {
		return model.BibleVerse{}, "Unable to get keywords from database", 500, err
	}
	result.Decode(&resultOne)

	// filter array to check if bible verse exists
	found := false
	for _, v := range resultOne.BibleVerse {
		if v.ID == BidHash {
			found = true
			break
		}
	}
	if !found {
		return model.BibleVerse{}, "Bible verse does not exist", 403, errors.New("bible verser does not exist in database")
	}

	// get the required verse
	var Oneverse model.BibleVerse
	for _, v := range resultOne.BibleVerse {
		if v.ID == BidHash {
			Oneverse = v
		}
	}

	return Oneverse, "", 0, nil
}

func AdminDeleteVersebyId(kid, Bid string) (model.BibleVerseResponse, string, int, error) {
	kidHash, _ := primitive.ObjectIDFromHex(kid)
	BidHash, _ := primitive.ObjectIDFromHex(Bid)

	// check if keyword  exists
	keywordsearch := map[string]any{
		"_id": kidHash,
	}

	keywordCount, _ := mongodb.MongoCount(constants.KeywordCollection, keywordsearch)
	if keywordCount < 1 {
		return model.BibleVerseResponse{}, "keyword doesnt exist", 403, errors.New("keyword doesnt exist in database")
	}

	// get from db
	var resultOne model.Keywords
	result, err := mongodb.MongoGetOne(constants.KeywordCollection, keywordsearch)
	if err != nil {
		return model.BibleVerseResponse{}, "Unable to get keywords from database", 500, err
	}
	result.Decode(&resultOne)

	// filter array to check if bible verse exists
	found := false
	for _, v := range resultOne.BibleVerse {
		if v.ID == BidHash {
			found = true
			break
		}
	}
	if !found {
		return model.BibleVerseResponse{}, "Bible verse does not exist", 403, errors.New("bible verser does not exist in database")
	}

	var index int
	for n, v := range resultOne.BibleVerse {
		if v.ID == BidHash {
			index = n
		}
	}

	var updateOne model.Keywords
	updateOne.ID = resultOne.ID
	updateOne.CategoryID = resultOne.CategoryID

	//deleting from slice in mongodb
	updateOne.BibleVerse = utility.DeleteElement(resultOne.BibleVerse, index)
	// save to DB
	_, err = mongodb.MongoUpdate(keywordsearch, updateOne, constants.KeywordCollection)
	if err != nil {
		return model.BibleVerseResponse{}, "Unable to save user to database", 500, err
	}

	// initiating response
	verseResponse := model.BibleVerseResponse{
		BibleVerse:  "",
		Passage:     "",
		Like:        false,
		Explanation: "",
	}
	return verseResponse, "", 0, nil
}

func AdminVerseInfo(kid string) (model.BibleVerseInfoResponse, string, int, error) {
	kidHash, _ := primitive.ObjectIDFromHex(kid)

	// check if keyword  exists
	keywordsearch := map[string]any{
		"_id": kidHash,
	}

	keywordCount, _ := mongodb.MongoCount(constants.KeywordCollection, keywordsearch)
	if keywordCount < 1 {
		return model.BibleVerseInfoResponse{}, "keyword doesnt exist", 403, errors.New("keyword doesnt exist in database")
	}

	// get from db
	var resultOne model.Keywords
	result, err := mongodb.MongoGetOne(constants.KeywordCollection, keywordsearch)
	if err != nil {
		return model.BibleVerseInfoResponse{}, "Unable to get keywords from database", 500, err
	}
	result.Decode(&resultOne)

	totalverses := len(resultOne.BibleVerse)

	// initiating response
	verseResponse := model.BibleVerseInfoResponse{
		TotalVerses: totalverses,
	}
	return verseResponse, "", 0, nil
}
