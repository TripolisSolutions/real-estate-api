package main

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	mongo "github.com/TripolisSolutions/go-helper/mgojuice"
)

const ImagesCollection = "images"

func FindImages() ([]Image, error) {
	var items []Image
	if err := mongo.Execute("monotonic", ImagesCollection,
		func(collection *mgo.Collection) error {
			return collection.Find(nil).All(&items)
		}); err != nil {
		return items, err
	}

	return items, nil
}

func (item *Image) Insert() error {
	item.CAt = time.Now()
	item.UAt = time.Now()

	if err := mongo.Execute("monotonic", ImagesCollection,
		func(collection *mgo.Collection) error {
			return collection.Insert(item)
		}); err != nil {
		return err
	}

	return nil
}

func (item *Image) Update() error {
	item.UAt = time.Now()

	if err := mongo.Execute("monotonic", ImagesCollection,
		func(collection *mgo.Collection) error {
			return collection.UpdateId(item.ID, item)
		}); err != nil {
		return err
	}

	return nil
}

func deleteImageByID(ID string) error {
	if err := mongo.Execute("monotonic", ImagesCollection,
		func(collection *mgo.Collection) error {
			return collection.RemoveId(bson.ObjectIdHex(ID))
		}); err != nil {
		return err
	}

	return nil
}
