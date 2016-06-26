package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func Index(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
	fmt.Fprint(ctx, "Welcome to REAL ESTATE API!\n")
}

func Hello(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
	fmt.Fprintf(ctx, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	category := categoryHandlers{}

	router := fasthttprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.GET("/categories", category.find)

	if err := fasthttp.ListenAndServe(":9001", router.Handler); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("fail start server")
	}
}
