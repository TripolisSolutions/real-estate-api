package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	log "github.com/Sirupsen/logrus"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

type categoryHandlers struct {
}

func (*categoryHandlers) find(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
	fmt.Fprintf(ctx, "hello, %s!\n", ps.ByName("name"))

	log.Infof("querystring %s", string(ctx.Request.URI().QueryString()))
	queries, err := url.ParseQuery(string(ctx.Request.URI().QueryString()))
	if err != nil {
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		return
	}

	coleID := queries.Get("col_id")

	log.WithFields(
		log.Fields{"cole_id": coleID},
	).Info("find categories")

	ctx.Response.SetStatusCode(200)
}

func (*categoryHandlers) create(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
	var payload struct {
		Name string `json:"name"`
	}
	json.Unmarshal(ctx.Request.Body(), &payload)
	
	if err := category.Insert() 
	//	log.WithFields(
	//		log.Fields{"cole_id": coleID},
	//	).Info("find categories")
}
