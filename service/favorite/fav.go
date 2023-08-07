package favorite

import (
	"errors"
	"time"

	"github.com/lordscoba/bible_compass_backend/internal/constants"
	"github.com/lordscoba/bible_compass_backend/internal/model"
	"github.com/lordscoba/bible_compass_backend/pkg/repository/storage/mongodb"
	"github.com/lordscoba/bible_compass_backend/utility"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func LikeKeywordsService(email, keyword string) (model.Fav, string, int, error) {

	// check if email  exists in users
	emailsearch := map[string]any{
		"email": email,
	}
	emailCount, _ := mongodb.MongoCount(constants.UserCollection, emailsearch)
	if emailCount < 1 {
		return model.Fav{}, "email does not exist", 403, errors.New("email does not exist in database")
	}

	// check if keyword  exists in keyword table
	keysearch := map[string]any{
		"keyword": keyword,
	}
	keyCount, _ := mongodb.MongoCount(constants.KeywordCollection, keysearch)
	if keyCount < 1 {
		return model.Fav{}, "keyword does not exist", 403, errors.New("keyword does not exist in database")
	}
	// get keywordId from db
	var resultKeys model.Keywords
	result1, err := mongodb.MongoGetOne(constants.KeywordCollection, keysearch)
	if err != nil {
		return model.Fav{}, "Unable to get favourite from database", 500, err
	}
	result1.Decode(&resultKeys)

	var userFav model.Fav
	userFav.Email = email
	userFav.Keyword = keyword

	// check if user  exists in favourite table
	usersearch := map[string]any{
		"email": email,
	}
	userCount, _ := mongodb.MongoCount(constants.FavCollection, usersearch)
	if userCount < 1 {

		//assign important values
		userFav.ID = primitive.NewObjectID()
		userFav.DateCreated = time.Now().Local()
		userFav.DateUpdated = time.Now().Local()

		// save to DB
		_, err := mongodb.MongoPost(constants.FavCollection, userFav)
		if err != nil {
			return model.Fav{}, "Unable to save keyword to database", 500, err
		}
	}

	// get from db
	var resultFavs model.Fav
	result, err := mongodb.MongoGetOne(constants.FavCollection, usersearch)
	if err != nil {
		return model.Fav{}, "Unable to get favourite from database", 500, err
	}
	result.Decode(&resultFavs)

	found := false
	// var index int
	for _, v := range resultFavs.Fav {
		if v.Keyword == keyword {
			found = true
			// index = i
			break
		}
	}

	var newdata model.FavData
	if !found {

		newdata.ID = resultKeys.ID.Hex()
		newdata.Keyword = keyword
		//append to slice
		userFav.Fav = append(resultFavs.Fav, newdata)
		// save to DB
		_, err = mongodb.MongoUpdate(usersearch, userFav, constants.FavCollection)
		if err != nil {
			return model.Fav{}, "Unable to save user to database", 500, err
		}
	}

	resultFavs.TotalFavs = len(resultFavs.Fav)

	return resultFavs, "", 0, nil
}

func UnLikeKeywordServive(email, keyword string) (model.Fav, string, int, error) {
	// check if email  exists in users
	emailsearch := map[string]any{
		"email": email,
	}
	emailCount, _ := mongodb.MongoCount(constants.UserCollection, emailsearch)
	if emailCount < 1 {
		return model.Fav{}, "email does not exist", 403, errors.New("email does not exist in database")
	}

	// check if keyword  exists in keyword table
	keysearch := map[string]any{
		"keyword": keyword,
	}
	keyCount, _ := mongodb.MongoCount(constants.KeywordCollection, keysearch)
	if keyCount < 1 {
		return model.Fav{}, "keyword does not exist", 403, errors.New("keyword does not exist in database")
	}
	// get keywordId from db
	// var resultKeys model.Keywords
	// result1, err := mongodb.MongoGetOne(constants.KeywordCollection, keysearch)
	// if err != nil {
	// 	return model.Fav{}, "Unable to get favourite from database", 500, err
	// }
	// result1.Decode(&resultKeys)

	var userFav model.Fav
	userFav.Email = email
	userFav.Keyword = keyword

	// check if user  exists in favourite table
	usersearch := map[string]any{
		"email": email,
	}
	userCount, _ := mongodb.MongoCount(constants.FavCollection, usersearch)
	if userCount < 1 {

		//assign important values
		userFav.ID = primitive.NewObjectID()
		userFav.DateCreated = time.Now().Local()
		userFav.DateUpdated = time.Now().Local()

		// save to DB
		_, err := mongodb.MongoPost(constants.FavCollection, userFav)
		if err != nil {
			return model.Fav{}, "Unable to save keyword to database", 500, err
		}
	}

	// get from db
	var resultFavs model.Fav
	result, err := mongodb.MongoGetOne(constants.FavCollection, usersearch)
	if err != nil {
		return model.Fav{}, "Unable to get favourite from database", 500, err
	}
	result.Decode(&resultFavs)

	found := false
	var index int
	for i, v := range resultFavs.Fav {
		if v.Keyword == keyword {
			found = true
			index = i
			break
		}
	}

	if found {
		//deleting from slice
		userFav.Fav = utility.DeleteElement(resultFavs.Fav, index)
		// save to DB
		_, err = mongodb.MongoUpdateForArrayBug(usersearch, userFav, constants.FavCollection)
		if err != nil {
			return model.Fav{}, "Unable to save user to database", 500, err
		}
	}

	resultFavs.TotalFavs = len(resultFavs.Fav)

	return resultFavs, "", 0, nil
}

func GetFavStatus(email, keyword string) (model.FavStatus, string, int, error) {

	// check if email  exists in users
	emailsearch := map[string]any{
		"email": email,
	}
	emailCount, _ := mongodb.MongoCount(constants.UserCollection, emailsearch)
	if emailCount < 1 {
		return model.FavStatus{}, "email does not exist", 403, errors.New("email does not exist in database")
	}

	// check if keyword  exists in keyword table
	keysearch := map[string]any{
		"keyword": keyword,
	}
	keyCount, _ := mongodb.MongoCount(constants.KeywordCollection, keysearch)
	if keyCount < 1 {
		return model.FavStatus{}, "keyword does not exist", 403, errors.New("keyword does not exist in database")
	}

	var userFavRes model.FavStatus
	userFavRes.Email = email
	userFavRes.Keyword = keyword

	// check if user  exists in favourite table
	usersearch := map[string]any{
		"email": email,
	}
	userCount, _ := mongodb.MongoCount(constants.FavCollection, usersearch)
	if userCount < 1 {
		userFavRes.Status = false
		return userFavRes, "", 0, nil
	}

	// get from db
	var resultFavs model.Fav
	result, err := mongodb.MongoGetOne(constants.FavCollection, usersearch)
	if err != nil {
		return model.FavStatus{}, "Unable to get favourite from database", 500, err
	}
	result.Decode(&resultFavs)

	found := false
	for _, v := range resultFavs.Fav {
		if v.Keyword == keyword {
			found = true
			break
		}
	}

	if !found {
		userFavRes.Status = false
		return userFavRes, "", 0, nil
	} else {
		userFavRes.Status = true
		return userFavRes, "", 0, nil
	}

}

func GetFavKeywordsService(email string) ([]model.FavData, string, int, error) {

	// check if email  exists in users
	emailsearch := map[string]any{
		"email": email,
	}
	emailCount, _ := mongodb.MongoCount(constants.UserCollection, emailsearch)
	if emailCount < 1 {
		return []model.FavData{}, "email does not exist", 403, errors.New("email does not exist in database")
	}

	// check if user  exists in favourite table
	usersearch := map[string]any{
		"email": email,
	}
	userCount, _ := mongodb.MongoCount(constants.FavCollection, usersearch)
	if userCount < 1 {
		return []model.FavData{}, "favourites empty", 0, nil

	}
	// get from db
	var resultFavs model.Fav
	result, err := mongodb.MongoGetOne(constants.FavCollection, usersearch)
	if err != nil {
		return []model.FavData{}, "Unable to get favourite from database", 500, err
	}
	result.Decode(&resultFavs)

	return resultFavs.Fav, "", 0, nil
}
