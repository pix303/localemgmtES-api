package main

import (
	"log"
	"net/http"

	"github.com/pix303/localemgmtES-api/router"
)

func main() {
	log.Println("Init api router")
	r, _ := router.NewAPI()
	log.Println("going to run router")
	log.Fatalln(http.ListenAndServe(":8080", r.Router))
	log.Println("something goes wrong")

}
