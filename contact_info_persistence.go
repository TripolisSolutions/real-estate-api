package main

import (
	"gopkg.in/mgo.v2"

	mongo "github.com/TripolisSolutions/go-helper/mgojuice"
)

const propertyDefaultContactInfoCollection = "default_contact_info"

func seedDefaultContactInfo() error {
	info := PropertyContactInfo{
		Phone:       "(+84) 981 688 076",
		OwnerName:   "Sonia-Phuong Tran",
		OwnerAvatar: "http://lorempixel.com/120/120/",
	}

	if err := mongo.Execute("monotonic", propertyDefaultContactInfoCollection,
		func(collection *mgo.Collection) error {
			return collection.Insert(info)
		}); err != nil {
		return err
	}

	return nil
}

func findDefaultContactInfo() ([]PropertyContactInfo, error) {
	var result []PropertyContactInfo
	if err := mongo.Execute("monotonic", propertyDefaultContactInfoCollection,
		func(collection *mgo.Collection) error {
			return collection.Find(nil).All(&result)
		}); err != nil {
		return result, err
	}
	return result, nil
}
