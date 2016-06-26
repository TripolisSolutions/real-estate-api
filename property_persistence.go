package main

import (
	"gopkg.in/mgo.v2"

	mongo "github.com/TripolisSolutions/go-helper/mgojuice"
)

const PropertyCollection = "properties"

func EnsureIndexPropertyCategory() error {
	if err := ensureIndex(PropertyCollection, mgo.Index{
		Name:       "",
		Key:        []string{"module", "user_ids"},
		Unique:     false,
		DropDups:   false,
		Background: false,
		Sparse:     false,
	}); err != nil {
		return err
	}
	return nil
}

func ensureIndex(collectionName string, index mgo.Index) error {
	if err := mongo.Execute("monotonic", collectionName,
		func(collection *mgo.Collection) error {
			return collection.EnsureIndex(index)
		}); err != nil {
		return err
	}
	return nil
}
