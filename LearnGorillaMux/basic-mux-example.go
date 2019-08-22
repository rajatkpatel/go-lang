package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()
	route.HandleFunc("/", HomeHandler)
	//r.HandleFunc("/articles", ArticlesHandler)
	route.HandleFunc("/articles/{category}", ArticlesCategoryHandler)
	//r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)

	product := route.PathPrefix("/products").Subrouter()

	// same as /products
	product.HandleFunc("", ProductsHandler)

	// same as /products/{key}
	product.HandleFunc("/{key}", ProductHandler)

	http.Handle("/", route)

	http.ListenAndServe(":8000", route)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome to the Home Page")
}

func ArticlesCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "No Products Available")
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%v product not available\n", vars["key"])
}
