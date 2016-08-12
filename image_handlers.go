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

type imageHandlers struct {
}

func (*imageHandlers) find(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
	if images, err := FindImages(); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Errorf("fail to find images")
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		return
	} else {
		ctx.Response.SetStatusCode(http.StatusOK)
		ctx.SetContentType("application/json")
		ctx.Response.SetBody(utilities.ToJSON(struct {
			Docs []Image `json:"docs"`
		}{Docs: images}))
	}
}

func (*imageHandlers) create(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
	var image Image
	json.Unmarshal(ctx.Request.Body(), &image)

	log.WithFields(log.Fields{
		"image": image,
	}).Debug("creating image")

	image.ID = bson.NewObjectId()
	if err := image.Insert(); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Errorf("fail insert property image")
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		return
	}

	log.WithFields(log.Fields{
		"image": image,
	}).Debug("created image")

	ctx.Response.SetStatusCode(http.StatusOK)
	ctx.Response.SetBody(utilities.ToJSON(image))
}

func (*imageHandlers) remove(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
	if err := deleteImageByID(ps.ByName("id")); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Errorf("fail to delete property")
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		return
	}

	ctx.Response.SetStatusCode(http.StatusOK)
}
