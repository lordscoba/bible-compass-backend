package category

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

func AdminCreateCategory(category model.Category) (model.CategoryResponse, string, int, error) {

	// check if category  exists
	namesearch := map[string]any{
		"category_name": category.CategoryName,
	}
	idCount, _ := mongodb.MongoCount(constants.CategoryCollection, namesearch)
	if idCount >= 1 {
		return model.CategoryResponse{}, "category already exist", 403, errors.New("category already exists in database")
	}

	category.ID = primitive.NewObjectID()
	category.DateCreated = time.Now().Local()
	category.DateUpdated = time.Now().Local()
	category.ForSubscribers = false

	// save to DB
	_, err := mongodb.MongoPost(constants.CategoryCollection, category)
	if err != nil {
		return model.CategoryResponse{}, "Unable to save user to database", 500, err
	}

	CategoryResponse := model.CategoryResponse{
		CategoryName:   category.CategoryName,
		ForSubscribers: category.ForSubscribers,
		Keywords:       category.Keywords,
		DateCreated:    category.DateCreated,
	}
	return CategoryResponse, "", 0, nil
}

func AdminUpdatecategory(category model.Category, id string) (model.CategoryResponse, string, int, error) {

	idHash, _ := primitive.ObjectIDFromHex(id)

	idsearch := map[string]any{
		"_id": idHash,
	}

	// check if id exists
	idCount, _ := mongodb.MongoCount(constants.CategoryCollection, idsearch)
	if idCount < 1 {
		return model.CategoryResponse{}, "category does not exist", 403, errors.New("category does not exist in database")
	}

	category.DateUpdated = time.Now().Local()

	// save to DB
	_, err := mongodb.MongoUpdate(idsearch, category, constants.CategoryCollection)
	if err != nil {
		return model.CategoryResponse{}, "Unable to save category to database", 500, err
	}

	categoryResponse := model.CategoryResponse{
		CategoryName:   category.CategoryName,
		ForSubscribers: category.ForSubscribers,
		Keywords:       category.Keywords,
		DateCreated:    category.DateCreated,
	}
	return categoryResponse, "", 0, nil
}

func AdminGetCategory() ([]model.Category, string, int, error) {

	// get from db
	result, err := mongodb.MongoGetAll(constants.CategoryCollection)
	if err != nil {
		return []model.Category{}, "Unable to save category to database", 500, err
	}

	var category = make([]model.Category, 0)
	result.All(context.TODO(), &category)
	return category, "", 0, nil
}

func AdminGetCategorybyId(id string) (model.Category, string, int, error) {
	idHash, _ := primitive.ObjectIDFromHex(id)
	search := map[string]any{
		"_id": idHash,
	}
	// get from db
	var resultOne model.Category
	result, err := mongodb.MongoGetOne(constants.CategoryCollection, search)
	if err != nil {
		return model.Category{}, "Unable to get category from database", 500, err
	}
	result.Decode(&resultOne)
	// get from db end

	fmt.Println(resultOne)
	return resultOne, "", 0, nil
}

func AdminDeleteCategorybyId(id string) (int64, string, int, error) {

	idHash, _ := primitive.ObjectIDFromHex(id)
	search := map[string]any{
		"_id": idHash,
	}

	// check if id exists
	idCount, _ := mongodb.MongoCount(constants.CategoryCollection, search)
	if idCount < 1 {
		return 0, "Category does not exist", 403, errors.New("category does not exist in database")
	}

	// get from db
	var resultOne model.Category
	result1, err := mongodb.MongoGetOne(constants.CategoryCollection, search)
	if err != nil {
		return 0, "Unable to get category from database", 500, err
	}
	result1.Decode(&resultOne)
	// get from db end

	for _, v := range resultOne.Keywords {

		searchkey := map[string]any{
			"keyword": v,
		}

		// delete from db
		_, err := mongodb.MongoDelete(constants.KeywordCollection, searchkey)
		if err != nil {
			return 0, "Unable to delete  category from database", 500, err
		}
	}

	// delete from db
	result, err := mongodb.MongoDelete(constants.CategoryCollection, search)
	if err != nil {
		return 0, "Unable to delete category from database", 500, err
	}

	fmt.Println(result.DeletedCount)
	return result.DeletedCount, "", 0, nil
}

func AdminCategoryInfo() (model.CategoryInfoResponse, string, int, error) {
	// total users
	search := map[string]any{}
	TotalCategory, err := mongodb.MongoCount(constants.CategoryCollection, search)
	if err != nil {
		return model.CategoryInfoResponse{}, "Unable to get count", 500, err
	}

	// total subscribers category
	searchSubscribed := map[string]any{
		"for_subscribers": true,
	}
	SubscribersCategory, err := mongodb.MongoCount(constants.CategoryCollection, searchSubscribed)
	if err != nil {
		return model.CategoryInfoResponse{}, "Unable to get count", 500, err
	}

	// total keywords
	keywordsSearch := map[string]any{}
	TotalKeyWords, err := mongodb.MongoCount(constants.CategoryCollection, keywordsSearch)
	if err != nil {
		return model.CategoryInfoResponse{}, "Unable to get count", 500, err
	}

	CategoryInfo := model.CategoryInfoResponse{
		TotalCategory:       TotalCategory,
		SubscribersCategory: SubscribersCategory,
		TotalKeyWords:       TotalKeyWords,
	}

	return CategoryInfo, "", 0, nil
}
