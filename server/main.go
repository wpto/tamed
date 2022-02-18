package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pgeowng/tamed/config"
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

	router := gin.Default()

	postRoute := routes.NewPostRoute(services)
	// userRoute := userroute.NewUserRoute(services)
	router.Static("/media/", config.Get().MediaPath)
	api := router.Group("/api")
	{
		api.GET("/post/", postRoute.List)
		api.GET("/post/:id", postRoute.Get)
		api.POST("/post/", postRoute.Create)
		// 	api.PATCH("/post/:id", userRoute.Modify)
		// 	api.DELETE("/post/:id", userRoute.Delete)
	}

	router.Run(":1314")
	return nil
}
