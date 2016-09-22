package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ryutoyasugi/go-webapp-sample/logger"
	"github.com/ryutoyasugi/go-webapp-sample/router"
)

func main() {

	logger.Logger()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}
	log.Println("server listen :", port)
	log.Fatal("alert: ", http.ListenAndServe(":"+port, router.Router()))
}
