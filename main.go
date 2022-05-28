package main

import (
	"BLOG_API/routers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("My Blog API")
	r := routers.Router()
	fmt.Println("Connection started")
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("Server Started.....")
}
