package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	online_diler "github.com/cora23tt/online-diler"
	"github.com/cora23tt/online-diler/pkg/handler"
	"github.com/cora23tt/online-diler/pkg/repository"
	"github.com/cora23tt/online-diler/pkg/service"
)

func main() {

	db, err := repository.NewPostgresDB()

	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(online_diler.Server)
	go func() {
		if err := srv.Run(handlers.InitRouts()); err != nil {
			log.Fatalf("error ocured while running http server: %s", err.Error())
		}
	}()

	log.Print("ToDo-app started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("ToDo-app Shutting Down")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("error ocured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Printf("error ocured on db connection close: %s", err.Error())
	}

	// router := gin.Default()
	// router.SetHTMLTemplate(template.Must(template.ParseGlob("templates/*.html")))
	// router.Static("../static", "./static")

	// router.GET("/", func(c *gin.Context) {
	// 	c.HTML(200, "index.html", gin.H{})
	// })

	// router.GET("/cart", func(c *gin.Context) {
	// 	c.HTML(200, "cart.html", gin.H{})
	// })

	// router.GET("/appsettings", func(c *gin.Context) {
	// 	c.HTML(200, "appsettings.html", gin.H{})
	// })

	// router.GET("/product", func(c *gin.Context) {
	// 	c.HTML(200, "product.html", gin.H{})
	// })

	// router.GET("/signin", func(c *gin.Context) {
	// 	c.HTML(200, "signin.html", gin.H{})
	// })

	// router.GET("/signout", func(c *gin.Context) {
	// 	// remove user session
	// 	c.Redirect(http.StatusFound, "/signin")
	// })

	// router.Run(":8080")
}
