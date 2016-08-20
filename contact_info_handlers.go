package main

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"

	"github.com/TripolisSolutions/go-helper/utilities"
)

type contactInfoHandlers struct {
}

func (*contactInfoHandlers) find(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
	if contactInfos, err := findDefaultContactInfo(); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Errorf("fail to find default contact infos")
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		return
	} else {
		ctx.Response.SetStatusCode(http.StatusOK)
		ctx.SetContentType("application/json")
		ctx.Response.SetBody(utilities.ToJSON(struct {
			Docs []PropertyContactInfo `json:"docs"`
		}{Docs: contactInfos}))
	}
}
