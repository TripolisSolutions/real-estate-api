package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"

	mongo "github.com/TripolisSolutions/go-helper/mgojuice"
)

func main() {
	if err := mongo.Startup(); err != nil {
		log.Fatalf("error[%s] while startup mongodb connection", err)
	}

	if err := EnsureIndexProperty(); err != nil {
		log.Fatalf("error[%s] while ensure index on properties collection", err)
	}

	if err := seedDataIfNeeded(); err != nil {
		log.Fatalf("error[%s] while seed data", err)
	}

	category := categoryHandlers{}
	property := propertyHandlers{}

	router := fasthttprouter.New()
	router.GET("/", Index)

	router.GET("/categories", category.find)
	router.POST("/categories", category.create)

	router.GET("/properties", property.find)
	router.POST("/properties", property.create)
	router.GET("/properties/:id", property.get)
	router.PUT("/properties/:id", property.update)
	router.DELETE("/properties/:id", property.remove)

	if err := fasthttp.ListenAndServe(":9001", router.Handler); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("fail start server")
	}
}

func Index(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
	fmt.Fprint(ctx, "Welcome to REAL ESTATE API!\n")
}
