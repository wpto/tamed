package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pgeowng/tamed/routes/userroute"
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

	router := gin.Default()

	viewRoute := viewroute.NewViewRoute(services)
	view := router.Group("/view")
	{
		view.GET("/art/:id", viewRoute.ViewArt)
		view.GET("/user/:id", viewRoute.ViewUser)
		view.GET("/search", viewRoute.Search)
	}

	userRoute := userroute.NewUserRoute(services)
	user := router.Group("/user")
	{
		user.POST("/art", userRoute.CreateArt)
	}

	router.Run(":1314")
	return nil
}
