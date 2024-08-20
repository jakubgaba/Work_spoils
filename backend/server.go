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
	log.Println(port)
	flag.Parse()

	router := gin.Default()

	// start server
	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalln(err)
		}
	}()

}
