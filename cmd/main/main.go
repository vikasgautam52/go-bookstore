package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/vikasgautam52/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	http.Handle("/", r)
	routes.RegisterBookStoreRoutes(r)
	log.Fatal(http.ListenAndServe(":9010", r))
}