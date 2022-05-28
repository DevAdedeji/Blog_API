package routers

import (
	"BLOG_API/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	// Create a blog
	router.HandleFunc("/api/blog", controllers.CreateBlog).Methods("POST")
	// Get all blogs
	router.HandleFunc("/api/blogs", controllers.GetAllBlogs).Methods("GET")
	// Get a blog
	router.HandleFunc("/api/blog/{id}", controllers.GetABlog).Methods("GET")

	return router
}
