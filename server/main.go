package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pgeowng/tamed/routes/viewroute"
	"github.com/pgeowng/tamed/service"
	"github.com/pgeowng/tamed/store"
	"github.com/pkg/errors"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	store, err := store.New()
	if err != nil {
		return errors.Wrap(err, "store.New failed")
	}

	services, err := service.NewManager(store)
	if err != nil {
		return errors.Wrap(err, "service.NewManager failed")
	}

	viewRoute := viewroute.NewViewRoute(services)

	router := gin.Default()
	view := router.Group("/view")
	{
		view.GET("/art/:id", viewRoute.ViewArt)
		view.GET("/user/:id", viewRoute.ViewUser)
		view.GET("/search", viewRoute.Search)
	}

	router.Run(":1314")
	return nil
}
