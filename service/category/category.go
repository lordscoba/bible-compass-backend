package category

import (
	"errors"
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
