package main

import (
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"gopkg.in/mgo.v2/bson"

	"github.com/TripolisSolutions/go-helper/utilities"
)

type categoryHandlers struct {
}

func (*categoryHandlers) find(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
	if categories, err := FindPropertyCategories(); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Errorf("fail to find property category")
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		return
	} else {
		ctx.Response.SetStatusCode(http.StatusOK)
		ctx.SetContentType("application/json")
		ctx.Response.SetBody(utilities.ToJSON(struct {
			Docs []PropertyCategory `json:"docs"`
		}{Docs: categories}))
	}
}

func (*categoryHandlers) create(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
	var category PropertyCategory
	json.Unmarshal(ctx.Request.Body(), &category)

	category.ID = bson.NewObjectId()
	if err := category.Insert(); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Errorf("fail insert property category")
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		return
	}

	ctx.Response.SetStatusCode(http.StatusCreated)
	ctx.Response.SetBody(utilities.ToJSON(category))
}
