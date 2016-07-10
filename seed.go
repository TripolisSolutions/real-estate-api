package main

import (
	log "github.com/Sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"

	mongo "github.com/TripolisSolutions/go-helper/mgojuice"
)

func seedDataIfNeeded() error {
	mgoSession, err := mongo.CloneMasterSession()
	if err != nil {
		return err
	}
	exists := mongo.CollectionExists(mgoSession, "", PropertyCategoryCollection)

	if exists {
		log.Infof("PropertyCategoryCollection exists!")
		return nil
	}

	apartment := PropertyCategory{
		ID: bson.NewObjectId(),
		Name: []TranslatableText{
			{
				Language: Vietnamese,
				Text:     "Căn hộ chung cư",
			},
			{
				Language: English,
				Text:     "Apartment",
			},
		},
	}

	luxuryApartment := PropertyCategory{
		ID: bson.NewObjectId(),
		Name: []TranslatableText{
			{
				Language: Vietnamese,
				Text:     "Căn hộ chung cư cao cấp",
			},
			{
				Language: English,
				Text:     "Luxury apartment",
			},
		},
	}

	if err := apartment.Insert(); err != nil {
		return err
	}

	if err := luxuryApartment.Insert(); err != nil {
		return err
	}

	return nil
}
