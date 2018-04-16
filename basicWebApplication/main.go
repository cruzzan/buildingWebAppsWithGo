package main

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func main()  {
	router := httprouter.New()

	router.GET("/", HomeHandler)

	// Collections
	router.GET("/posts", PostsIndexHandler)
	router.POST("/posts", PostsCreateHandler)

	// Singulars
	router.GET("/posts/:id/edit", PostsEditFormHandler)
	router.PUT("/posts/:id", PostsUpdateHandler)
	router.GET("/posts/:id", PostsShowHandler)
	router.DELETE("/posts/:id", PostsDeleteHandler)

	fmt.Println("Starting server on :8181")
	http.ListenAndServe(":8181", router)
}

func HomeHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "You are home")
}

func PostsIndexHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintln(rw, "Get all posts")
}

func PostsCreateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	fmt.Fprintln(rw, "Create a new post")
}

func PostsEditFormHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	id := p.ByName("id")
	fmt.Fprintln(rw, "Edit form for post id:", id)
}

func PostsUpdateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	fmt.Fprintln(rw, "Updating post id:", id)
}

func PostsShowHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	fmt.Fprintln(rw, "Show post id:", id)
}

func PostsDeleteHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	id := p.ByName("id")
	fmt.Fprintln(rw, "Delete post id:", id)
}
