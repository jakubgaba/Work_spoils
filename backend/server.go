package backend

import (
	"errors"
	"flag"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Server() {

	host := flag.String("host", "localhost", "Server host")
	log.Println(host)
	port := flag.Int("port", 8080, "Server port")
	docker := flag.Bool("docker", false, "Running in docker")
	log.Println(port)
	flag.Parse()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	service := Service{}

	//apis
	api := router.Group("/api")
	api.GET("/test", service.testService)

	// start server
	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalln(err)
		}
	}()

}
