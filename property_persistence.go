package main

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	mongo "github.com/TripolisSolutions/go-helper/mgojuice"
)

const PropertyCollection = "properties"
const PropertyTrashCollection = "properties_trash"

func EnsureIndexProperty() error {
	if err := ensureIndex(PropertyCollection, mgo.Index{
		Name:             "text",
		Key:              []string{"$text:name.text", "$text:desc.text", "$text:address.name.text"},
		DefaultLanguage:  "english",
		LanguageOverride: "english",
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

func FindProperties(filterers bson.M, limit, offset int) ([]Property, error) {
	var properties []Property
	if err := mongo.Execute("monotonic", PropertyCollection,
		func(collection *mgo.Collection) error {
			return collection.Find(filterers).Limit(limit).Skip(offset).All(&properties)
		}); err != nil {
		return properties, err
	}

	return properties, nil
}

func CountProperties() (int, error) {
	var result int
	if err := mongo.Execute("monotonic", PropertyCollection,
		func(collection *mgo.Collection) error {
			count, err := collection.Find(nil).Count()
			result = count
			return err
		}); err != nil {
		return result, err
	}

	return result, nil
}

func (property *Property) FindByID(ID string) error {
	if err := mongo.Execute("monotonic", PropertyCollection,
		func(collection *mgo.Collection) error {
			return collection.FindId(bson.ObjectIdHex(ID)).One(&property)
		}); err != nil {
		return err
	}

	return nil
}

func (property *Property) Insert() error {
	property.CAt = time.Now()
	property.UAt = time.Now()

	for i := range property.Desc {
		sanitizedHtml, err := sanitizeHtml(property.Desc[i].Text)
		if err != nil {
			return err
		}

		property.Desc[i].Text = sanitizedHtml
	}

	if err := mongo.Execute("monotonic", PropertyCollection,
		func(collection *mgo.Collection) error {
			return collection.Insert(property)
		}); err != nil {
		return err
	}

	return nil
}

func (property *Property) Update() error {
	property.UAt = time.Now()

	for i := range property.Desc {
		sanitizedHtml, err := sanitizeHtml(property.Desc[i].Text)
		if err != nil {
			return err
		}

		property.Desc[i].Text = sanitizedHtml
	}

	if err := mongo.Execute("monotonic", PropertyCollection,
		func(collection *mgo.Collection) error {
			return collection.UpdateId(property.ID, property)
		}); err != nil {
		return err
	}

	return nil
}

func deletePropertyByID(ID string) error {
	property := Property{}
	if err := property.FindByID(ID); err != nil {
		return err
	}

	if err := mongo.Execute("monotonic", PropertyTrashCollection,
		func(collection *mgo.Collection) error {
			return collection.Insert(property)
		}); err != nil {
		return err
	}

	if err := mongo.Execute("monotonic", PropertyCollection,
		func(collection *mgo.Collection) error {
			return collection.RemoveId(bson.ObjectIdHex(ID))
		}); err != nil {
		return err
	}

	return nil
}
