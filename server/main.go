package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pgeowng/tamed/routes/artroute"
	"github.com/pgeowng/tamed/routes/userroute"
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

	artRoute := artroute.NewArtRoute(services)
	userRoute := userroute.NewUserRoute(services)
	// mediaRoute := mediaroute.NewMediaRoute(services)

	router := gin.Default()
	// v1 := router.Group("/media")
	// {
	// 	// v1.GET("/:id/:quality", mediaRoute.Get)
	// }

	art := router.Group("/art")
	{
		// art.POST("/", artRoute.Create)
		art.GET("/:id", artRoute.Get)
		// art.PUT("/:id", artRoute.Update)
		// art.DELETE("/:id", artRoute.Delete)
	}

	user := router.Group("/user")
	{
		user.GET("/:id", userRoute.Get)

	}

	router.Run(":1314")
	return nil
}
