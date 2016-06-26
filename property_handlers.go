package main

import (
	"encoding/json"
	"net/http"
	"net/url"
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

	limit := ParseIntWithFallback(queries.Get("limit"), 50)
	offset := ParseIntWithFallback(queries.Get("offset"), 0)

	var filterers bson.M

	q := strings.TrimSpace(queries.Get("q"))
	language := strings.TrimSpace(queries.Get("language"))
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
		filterers["category_id"] = bson.ObjectIdHex(categoryID)
	}

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

	total, err := CountProperties()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Errorf("fail to count properties")
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		return
	}

	ctx.Response.SetStatusCode(http.StatusOK)
	ctx.Response.SetBody(utilities.ToJSON(struct {
		Docs  []Property
		Total int
	}{
		Docs:  properties,
		Total: total,
	}))

}

func (*propertyHandlers) create(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
	var property Property
	json.Unmarshal(ctx.Request.Body(), &property)

	property.ID = bson.NewObjectId()
	if err := property.Insert(); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Errorf("fail to insert property")
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		return
	}

	ctx.Response.SetStatusCode(http.StatusCreated)
	ctx.Response.SetBody(utilities.ToJSON(struct {
		Doc Property
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
		Doc Property
	}{
		Doc: property,
	}))
}

func (*propertyHandlers) remove(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {

}
