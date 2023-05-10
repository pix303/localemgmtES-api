package main

import (
	"log"
	"net/http"
	"os"

	"github.com/pix303/localemgmtES-api/router"
)

func main() {
	log.Println("Init api router")
	r, _ := router.NewAPI()
	log.Println("going to run router")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatalln(http.ListenAndServe("0.0.0.0:"+port, r.Router))
	log.Println("something goes wrong")

}
