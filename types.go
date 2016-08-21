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
const RPUDays RentalPeriodUnit = "dayclassNames"

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

type PropertyContactInfo struct {
	Phone       string `bson:"phone" json:"phone"`
	OwnerName   string `bson:"ownerName" json:"ownerName"`
	OwnerAvatar Image  `bson:"ownerAvatar" json:"ownerAvatar"`
}

type Property struct {
	ID            bson.ObjectId      `bson:"_id" json:"id"`
	Name          []TranslatableText `bson:"name" json:"name"`
	Thumbnail     *Image             `bson:"thumbnailImage,omitempty" json:"thumbnailImage,omitempty"`
	GalleryImages []Image            `bson:"galleryImages" json:"galleryImages"`
	Desc          []TranslatableText `bson:"desc" json:"desc"`

	ContactInfo []PropertyContactInfo `bson:"contactInfos" json:"contactInfos"`

	CategoryID     *bson.ObjectId `bson:"categoryID,omitempty" json:"categoryID,omitempty"` // PropertyCategory.ID
	SalesType      string         `bson:"salesType" json:"salesType"`
	AvailableUntil *time.Time     `bson:"availableUntil,omitempty" json:"availableUntil,omitempty"`
	Size           *struct {
		Width  float32 `bson:"width" json:"width"`
		Length float32 `bson:"length" json:"length"`
		Area   float32 `bson:"area" json:"area"`
	} `bson:"size,omitempty" json:"size,omitempty"`
	Address *struct {
		Name     []TranslatableText `bson:"name" json:"name"`
		District string             `bson:"district" json:"district"`
		Viewport *struct {
			Lat  float64 `bson:"lat" json:"lat"`
			Lng  float64 `bson:"lng" json:"lng"`
			Zoom float64 `bson:"zoom" json:"zoom"`
		} `bson:"viewport,omitempty" json:"viewport,omitemptyt"`
		CircleMarker *struct {
			Lat    float64 `bson:"lat" json:"lat"`
			Lng    float64 `bson:"lng" json:"lng"`
			Radius float64 `bson:"radius" json:"radius"`
		} `bson:"circleMarker,omitempty" json:"circleMarker,omitempty"`
		Visible bool `bson:"visible" json:"visible"`
	} `bson:"address,omitempty" json:"address,omitempty"`
	BedRoomCount    int                     `bson:"bedRoomCount" json:"bedRoomCount"`
	FacingDirection PropertyFacingDirection `bson:"facingDirection" json:"facingDirection"`

	RentalPeriod struct {
		Negotiable bool             `bson:"negotiable" json:"negotiable"`
		Digits     float32          `bson:"digits" json:"digits"`
		Unit       RentalPeriodUnit `bson:"unit" json:"unit"`
	} `bson:"rentalPeriod" json:"rentalPeriod"`
	Price []TranslatablePrice `bson:"price" json:"price"`

	Visible bool `bson:"visible" json:"visible"`

	CAt time.Time `bson:"c_at" json:"c_at"`
	UAt time.Time `bson:"u_at" json:"u_at"`
}

type Image struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	URL      string        `bson:"url" json:"url"`
	Width    int           `bson:"width" json:"width"`
	Height   int           `bson:"height" json:"height"`
	FileName string        `bson:"fileName" json:"fileName"`
	CAt      time.Time     `bson:"c_at" json:"c_at"`
	UAt      time.Time     `bson:"u_at" json:"u_at"`
}
