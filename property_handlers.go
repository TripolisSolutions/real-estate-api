package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"gopkg.in/mgo.v2/bson"

	"github.com/TripolisSolutions/go-helper/utilities"
)

type propertyHandlers struct {
}

func (*propertyHandlers) find(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {

	queries, err := url.ParseQuery(string(ctx.Request.URI().QueryString()))
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		return
	}

	limit := ParseIntWithFallback(queries.Get("limit"), 20)
	var offset int

	offset = ParseIntWithFallback(queries.Get("offset"), 0)

	if queries.Get("page") != "" {
		page := ParseIntWithFallback(queries.Get("page"), 0)
		offset = page * limit
	}

	if limit > 100 {
		limit = 100
	}

	var filterers = bson.M{}

	q := strings.TrimSpace(queries.Get("q"))
	language := strings.TrimSpace(queries.Get("language"))

	if queries.Get("lang") != "" {
		language = langCodeToLanguage(queries.Get("lang"))
	}

	if q != "" {
		if language == "" || !isLanguageSupported(language) {
			log.WithFields(log.Fields{
				"language": language,
			}).Infof("invalid language")
			ctx.Response.SetStatusCode(http.StatusBadRequest)
			ctx.Response.SetBodyString("invalid 'language' param")
			return
		}

		filterers["$text"] = bson.M{
			"$search":             q,
			"$language":           language,
			"$diacriticSensitive": false,
			"$caseSensitive":      false,
		}
	}

	categoryID := strings.TrimSpace(queries.Get("category"))
	if categoryID != "" {
		if !bson.IsObjectIdHex(categoryID) {
			ctx.Response.SetStatusCode(http.StatusBadRequest)
			ctx.Response.SetBodyString("invalid 'category' param")
			return
		}
		filterers["categoryID"] = bson.ObjectIdHex(categoryID)
	}

	salesType := strings.TrimSpace(queries.Get("salesType"))
	if salesType != "" {
		filterers["salesType"] = salesType
	}

	minBed := strings.TrimSpace(queries.Get("minBed"))
	if minBed != "" {
		minBedValue, err := strconv.Atoi(minBed)
		if err != nil {
			ctx.Response.SetStatusCode(http.StatusBadRequest)
			ctx.Response.SetBodyString("invalid 'minBed' param")
			return
		}
		filterers["bedRoomCount"] = bson.M{
			"$gte": minBedValue,
		}
	}

	maxBed := strings.TrimSpace(queries.Get("maxBed"))
	if maxBed != "" {
		maxBedValue, err := strconv.Atoi(maxBed)
		if err != nil {
			ctx.Response.SetStatusCode(http.StatusBadRequest)
			ctx.Response.SetBodyString("invalid 'maxBed' param")
			return
		}
		filterers["bedRoomCount"] = bson.M{
			"$lte": maxBedValue,
		}
	}

	currency := strings.TrimSpace(queries.Get("currency"))

	minPrice := strings.TrimSpace(queries.Get("minPrice"))
	if minPrice != "" {
		minPriceValue, err := strconv.Atoi(minPrice)
		if err != nil {
			ctx.Response.SetStatusCode(http.StatusBadRequest)
			ctx.Response.SetBodyString("invalid 'minPrice' param")
			return
		}

		if !isCurrencySupported(currency) {
			log.WithFields(log.Fields{
				"currency": currency,
			}).Infof("invalid currency")
			ctx.Response.SetStatusCode(http.StatusBadRequest)
			ctx.Response.SetBodyString("invalid 'currency' param")
			return
		}

		filterers["price"] = bson.M{
			"$elemMatch": bson.M{
				"currency": currency,
				"value": bson.M{
					"$gte": minPriceValue,
				},
			},
		}
	}

	maxPrice := strings.TrimSpace(queries.Get("maxPrice"))
	if maxPrice != "" {
		maxPriceValue, err := strconv.Atoi(maxPrice)
		if err != nil {
			ctx.Response.SetStatusCode(http.StatusBadRequest)
			ctx.Response.SetBodyString("invalid 'maxPrice' param")
			return
		}

		if !isCurrencySupported(currency) {
			log.WithFields(log.Fields{
				"currency": currency,
			}).Infof("invalid currency")
			ctx.Response.SetStatusCode(http.StatusBadRequest)
			ctx.Response.SetBodyString("invalid 'currency' param")
			return
		}

		filterers["price"] = bson.M{
			"$elemMatch": bson.M{
				"currency": currency,
				"value": bson.M{
					"$lte": maxPriceValue,
				},
			},
		}
	}

	district := strings.TrimSpace(queries.Get("district"))
	if district != "" {
		filterers["address.district"] = district
	}

	size := strings.TrimSpace(queries.Get("size"))
	// e.g. -30, 80-100, 500-
	if size != "" {
		frags := strings.Split(size, "-")
		if len(frags) != 2 {
			ctx.Response.SetStatusCode(http.StatusBadRequest)
			ctx.Response.SetBodyString("invalid 'size' param: " + size)
			return
		}

		if frags[0] != "" {
			from, err := strconv.Atoi(frags[0])
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusBadRequest)
				ctx.Response.SetBodyString("invalid 'size' param")
				return
			}

			filterers["size.area"] = bson.M{
				"$gte": from,
			}
		}

		if frags[1] != "" {
			to, err := strconv.Atoi(frags[1])
			if err != nil {
				ctx.Response.SetStatusCode(http.StatusBadRequest)
				ctx.Response.SetBodyString("invalid 'size' param")
				return
			}

			filterers["size.area"] = bson.M{
				"$lte": to,
			}
		}
	}

	log.WithFields(log.Fields{
		"filterers": string(utilities.ToJSON(filterers)),
		"limit":     limit,
		"offset":    offset,
	}).Info("finding properties")

	properties, err := FindProperties(filterers, limit, offset)
	if err != nil {
		log.WithFields(log.Fields{
			"error":     err,
			"filterers": filterers,
			"limit":     limit,
			"offset":    offset,
		}).Errorf("fail to find properties")
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		return
	}

	total, err := CountProperties(filterers)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Errorf("fail to count properties")
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		return
	}

	ctx.Response.SetStatusCode(http.StatusOK)
	ctx.Response.SetBody(utilities.ToJSON(struct {
		Docs  []Property `json:"docs"`
		Total int        `json:"total"`
	}{
		Docs:  properties,
		Total: total,
	}))

}

func (*propertyHandlers) create(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
	var property Property
	json.Unmarshal(ctx.Request.Body(), &property)

	property.ID = bson.NewObjectId()

	log.WithFields(log.Fields{
		"property": property,
	}).Info("creating")

	if err := property.Insert(); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Errorf("fail to insert property")
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		return
	}

	if err := property.ContactInfo[0].SaveAsDefault(); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Errorf("fail to save contact info as default")
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		return
	}

	ctx.Response.SetStatusCode(http.StatusCreated)
	ctx.Response.SetBody(utilities.ToJSON(struct {
		Doc Property `json:"doc"`
	}{
		Doc: property,
	}))
}

func (*propertyHandlers) update(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
	var property Property

	if err := property.FindByID(ps.ByName("id")); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Errorf("fail to find property")
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		return
	}

	json.Unmarshal(ctx.Request.Body(), &property)

	log.WithFields(log.Fields{
		"property": property,
	}).Info("updating")

	if err := property.Update(); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Errorf("fail to insert property")
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		return
	}

	if err := property.ContactInfo[0].SaveAsDefault(); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Errorf("fail to save contact info as default")
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		return
	}

	ctx.Response.SetStatusCode(http.StatusCreated)
	ctx.Response.SetBody(utilities.ToJSON(struct {
		Doc Property `json:"doc"`
	}{
		Doc: property,
	}))
}

func (*propertyHandlers) get(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
	var property Property
	if err := property.FindByID(ps.ByName("id")); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Errorf("fail to find property")
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		return
	}

	ctx.Response.SetStatusCode(http.StatusCreated)
	ctx.Response.SetBody(utilities.ToJSON(struct {
		Doc Property `json:"doc"`
	}{
		Doc: property,
	}))
}

func (*propertyHandlers) remove(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
	if err := deletePropertyByID(ps.ByName("id")); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Errorf("fail to delete property")
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		return
	}

	ctx.Response.SetStatusCode(http.StatusOK)
}
