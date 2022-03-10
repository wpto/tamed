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

	// router.LoadHTMLGlob("templates/*.tmpl.html")
	// router.Static("/static", "static")
	// router.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.tmpl.html", nil)
	// })

	router.StaticFile("/", "client/public/index.html")
	router.StaticFile("/global.css", "client/public/global.css")
	router.Static("/build/", "client/public/build")

	postRoute := routes.NewPostRoute(services)
	router.Static("/media/", config.Get().FsMediaPath)
	api := router.Group("/api")
	{
		api.GET("/posts", postRoute.List)
		api.GET("/posts/:id", postRoute.Get)
		api.POST("/posts", postRoute.Create)
		api.PATCH("/posts/:id", postRoute.Modify)
		api.DELETE("/posts/:id", postRoute.Delete)
	}

	router.Run(":" + config.Get().Port)
	return nil
}
