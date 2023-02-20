package main

import (
	"fmt"
	"go-postgres-yt/router"
	"log"
	"net/http"
)

func main(){
	r := router.Router()
	fmt.Println("starting server on 8080")

	log.Fatal(http.ListenAndServe(":8080",r))
}