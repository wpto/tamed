package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pgeowng/tamed/routes"
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

	mediaRoute := routes.NewMediaRoute(services)

	router := gin.Default()
	v1 := router.Group("/media")
	{
		v1.GET("/:id", mediaRoute.Get)
		v1.POST("/", mediaRoute.Upload)
	}

	router.Run(":1314")
	return nil
}
