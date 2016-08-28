package main

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	mongo "github.com/TripolisSolutions/go-helper/mgojuice"
)

const propertyDefaultContactInfoCollection = "default_contact_info"

func seedDefaultContactInfo() error {
	avatar := Image{
		ID:  bson.NewObjectId(),
		URL: "http://lorempixel.com/120/120/",
		CAt: time.Now(),
		UAt: time.Now(),
	}

	if err := avatar.Insert(); err != nil {
		return err
	}

	info1 := PropertyContactInfo{
		Phone:       "(+84) 981 688 076",
		OwnerName:   "Sonia-Phuong Tran",
		OwnerAvatar: avatar,
	}

	if err := info1.Insert(); err != nil {
		return err
	}

	info2 := PropertyContactInfo{
		Phone:       "(+84) 981 688 075",
		OwnerName:   "Dean Walkerden",
		OwnerAvatar: avatar,
	}

	return info2.Insert()
}

func (c *PropertyContactInfo) Insert() error {
	if err := mongo.Execute("monotonic", propertyDefaultContactInfoCollection,
		func(collection *mgo.Collection) error {
			return collection.Insert(c)
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

func SaveAsDefault(contactInfos []PropertyContactInfo) error {
	if err := mongo.Execute("monotonic", propertyDefaultContactInfoCollection,
		func(collection *mgo.Collection) error {
			_, err := collection.RemoveAll(nil)
			return err
		}); err != nil {
		return err
	}

	for _, contactInfo := range contactInfos {
		if err := contactInfo.Insert(); err != nil {
			return err
		}
	}

	return nil
}
