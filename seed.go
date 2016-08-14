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
				Text:     "Apartments",
			},
		},
	}
	if err := apartment.Insert(); err != nil {
		return err
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
				Text:     "Luxury apartments",
			},
		},
	}

	if err := luxuryApartment.Insert(); err != nil {
		return err
	}

	if err := (&PropertyCategory{
		ID: bson.NewObjectId(),
		Name: []TranslatableText{
			{
				Language: Vietnamese,
				Text:     "Nhà ở",
			},
			{
				Language: English,
				Text:     "Houses",
			},
		},
	}).Insert(); err != nil {
		return err
	}

	if err := (&PropertyCategory{
		ID: bson.NewObjectId(),
		Name: []TranslatableText{
			{
				Language: Vietnamese,
				Text:     "Nhà phố",
			},
			{
				Language: English,
				Text:     "Townhouses",
			},
		},
	}).Insert(); err != nil {
		return err
	}

	if err := (&PropertyCategory{
		ID: bson.NewObjectId(),
		Name: []TranslatableText{
			{
				Language: Vietnamese,
				Text:     "Biệt thự",
			},
			{
				Language: English,
				Text:     "Villa",
			},
		},
	}).Insert(); err != nil {
		return err
	}

	if err := (&PropertyCategory{
		ID: bson.NewObjectId(),
		Name: []TranslatableText{
			{
				Language: Vietnamese,
				Text:     "Phòng",
			},
			{
				Language: English,
				Text:     "Rooms",
			},
		},
	}).Insert(); err != nil {
		return err
	}

	if err := (&PropertyCategory{
		ID: bson.NewObjectId(),
		Name: []TranslatableText{
			{
				Language: Vietnamese,
				Text:     "Văn phòng",
			},
			{
				Language: English,
				Text:     "Offices",
			},
		},
	}).Insert(); err != nil {
		return err
	}

	if err := (&PropertyCategory{
		ID: bson.NewObjectId(),
		Name: []TranslatableText{
			{
				Language: Vietnamese,
				Text:     "Đất nền",
			},
			{
				Language: English,
				Text:     "Lands",
			},
		},
	}).Insert(); err != nil {
		return err
	}

	if err := (&PropertyCategory{
		ID: bson.NewObjectId(),
		Name: []TranslatableText{
			{
				Language: Vietnamese,
				Text:     "Cửa hàng, Ki ốt",
			},
			{
				Language: English,
				Text:     "Shops/Kiosks",
			},
		},
	}).Insert(); err != nil {
		return err
	}

	if err := (&PropertyCategory{
		ID: bson.NewObjectId(),
		Name: []TranslatableText{
			{
				Language: Vietnamese,
				Text:     "Kho, nhà xưởng",
			},
			{
				Language: English,
				Text:     "Warehouses",
			},
		},
	}).Insert(); err != nil {
		return err
	}

	if err := (&PropertyCategory{
		ID: bson.NewObjectId(),
		Name: []TranslatableText{
			{
				Language: Vietnamese,
				Text:     "Khách sạn",
			},
			{
				Language: English,
				Text:     "Hotels",
			},
		},
	}).Insert(); err != nil {
		return err
	}

	if err := (&PropertyCategory{
		ID: bson.NewObjectId(),
		Name: []TranslatableText{
			{
				Language: Vietnamese,
				Text:     "Loại bất động sản khác",
			},
			{
				Language: English,
				Text:     "Other types",
			},
		},
	}).Insert(); err != nil {
		return err
	}

	return nil
}
