package utility

import (
	"reflect"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func IsValidPassword(userPassword, providedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword)) == nil
}

func StructToMap(inputStruct interface{}) map[string]interface{} {
	structType := reflect.TypeOf(inputStruct)
	structValue := reflect.ValueOf(inputStruct)

	if structType.Kind() != reflect.Struct {
		return nil
	}
	resultMap := make(map[string]interface{})

	for i := 1; i < structType.NumField(); i++ {
		field := structType.Field(i)
		jsonTag := field.Tag.Get("json")
		value := structValue.Field(i).Interface()
		if !IsEmpty(value) {
			resultMap[jsonTag] = value
		}
	}

	return resultMap
}

func MapToBsonD[T any](inputMap map[string]T) bson.D {
	elements := make([]bson.E, 0, len(inputMap))

	for key, value := range inputMap {
		element := bson.E{Key: key, Value: value}
		elements = append(elements, element)
	}

	return elements
}

func IsEmpty(value interface{}) bool {
	if value == nil {
		return true
	}
	switch v := reflect.ValueOf(value); v.Kind() {
	case reflect.String:
		return strings.TrimSpace(v.String()) == ""
	case reflect.Array, reflect.Slice, reflect.Map:
		return v.Len() == 0
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	default:
		return reflect.DeepEqual(value, reflect.Zero(v.Type()).Interface())
	}
}

func ComparingUpdate[T any](fromCurrent, fromDatabase T) T {
	if v := fromCurrent; !IsEmpty(v) {
		return v
	}
	return fromDatabase
}

func DeleteElement[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}
