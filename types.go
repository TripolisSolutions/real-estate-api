package main

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// PropertyFacingDirection
type PropertyFacingDirection string

const PFDNorth PropertyFacingDirection = "north"
const PFDNorthEast PropertyFacingDirection = "north-east"
const PFDEast PropertyFacingDirection = "east"
const PFDSouthEast PropertyFacingDirection = "south-east"
const PFDSouth PropertyFacingDirection = "south"
const PFDSouthWest PropertyFacingDirection = "south-west"
const PFDWest PropertyFacingDirection = "west"
const PFDNorthWest PropertyFacingDirection = "north-west"

// RentalPeriodUnit
type RentalPeriodUnit string

const RPUMonths RentalPeriodUnit = "months"
const RPUYears RentalPeriodUnit = "years"
const RPUDays RentalPeriodUnit = "days"

// PricingUnit
type PricingUnit string

const PUVietnamDong PricingUnit = "VND"
const PUUSD PricingUnit = "USD"
const PUEuro PricingUnit = "EURO"

type TranslatablePrice struct {
	Value    float32     `bson:"value" json:"value"`
	Currency PricingUnit `bson:"currency" json:"currency"`
}

// languages
const Vietnamese = "vietnamese"
const English = "english"

type TranslatableText struct {
	Language string `bson:"language" json:"language"`
	Text     string `bson:"text" json:"text"`
}

type PropertyCategory struct {
	ID   bson.ObjectId      `bson:"_id" json:"id"`
	Name []TranslatableText `bson:"name" json:"name"`

	CAt time.Time `bson:"c_at" json:"c_at"`
	UAt time.Time `bson:"u_at" json:"u_at"`
}

type Image struct {
	ID     string `bson:"_id" json:"id"`
	URL    string `bson:"url" json:"url"`
	Width  int    `bson:"width" json:"width"`
	Height int    `bson:"height" json:"height"`
}

type Property struct {
	ID            bson.ObjectId      `bson:"_id" json:"id"`
	Name          []TranslatableText `bson:"name" json:"name"`
	Thumbnail     Image              `bson:"thumbnailImage" json:"thumbnailImage"`
	GalleryImages []Image            `bson:"galleryImages" json:"galleryImages"`
	Desc          []TranslatableText `bson:"desc" json:"desc"`

	CategoryID     *bson.ObjectId `bson:"categoryID" json:"categoryID"` // PropertyCategory.ID
	SalesType      string         `bson:"salesType" json:"salesType"`
	AvailableUntil time.Time      `bson:"availableUntil" json:"availableUntil"`
	Size           struct {
		Width  float32 `bson:"width" json:"width"`
		Length float32 `bson:"depth" json:"depth"`
	} `bson:"size" json:"size"`
	Address struct {
		Name     []TranslatableText `bson:"name" json:"name"`
		Viewport struct {
			Lat  float64 `bson:"lat" json:"lat"`
			Lng  float64 `bson:"lng" json:"lng"`
			Zoom float64 `bson:"zoom" json:"zoom"`
		} `bson:"viewport" json:"viewport"`
		CircleMarker struct {
			Lat    float64 `bson:"lat" json:"lat"`
			Lng    float64 `bson:"lng" json:"lng"`
			Radius float64 `bson:"radius" json:"radius"`
		} `bson:"circleMarker" json:"circleMarker"`
		Visible bool `bson:"visible" json:"visible"`
	} `bson:"address" json:"address"`
	BedRoomCount    int                     `bson:"bedRoomCount" json:"bedRoomCount"`
	FacingDirection PropertyFacingDirection `bson:"facingDirection" json:"facingDirection"`

	RentalPeriod struct {
		Digits float32          `bson:"digits" json:"digits"`
		Unit   RentalPeriodUnit `bson:"unit" json:"unit"`
	} `bson:"rental_period" json:"rentalPeriod"`
	Price TranslatablePrice `bson:"price" json:"price"`

	Visible bool `bson:"visible" json:"visible"`

	CAt time.Time `bson:"c_at" json:"c_at"`
	UAt time.Time `bson:"u_at" json:"u_at"`
}
