package main

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/validator.v2"

	mongo "github.com/TripolisSolutions/go-helper/mgojuice"
)

const PropertyCategoryCollection = "property_categories"

// Insert is ..
func (category *PropertyCategory) Insert() error {

	if err := validator.Validate(category); err != nil {
		return err
	}

	category.UAt = time.Now()
	category.CAt = time.Now()

	if err := mongo.Execute("monotonic", PropertyCategoryCollection,
		func(collection *mgo.Collection) error {
			return collection.Insert(category)
		}); err != nil {
		return err
	}

	return nil
}

func FindPropertyCategories() ([]PropertyCategory, error) {
	var categories []PropertyCategory
	if err := mongo.Execute("monotonic", PropertyCategoryCollection,
		func(collection *mgo.Collection) error {
			return collection.Find(nil).All(&categories)
		}); err != nil {
		return categories, err
	}

	return categories, nil
}
