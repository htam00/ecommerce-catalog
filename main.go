package main

import (
	"net/http"
	"context"
	"fmt"
	"log"
	"time"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Index: %w\n", vars["Handler"])
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Tasks: %w\n", vars["refactory code"])
}

func ProductsIDHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://matheus:zurik21@ds237989.mlab.com:37989/ecommerce_catalog"))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/products", ProductsHandler)
	r.HandleFunc("/products/{id}", ProductsIDHandler)
	http.Handle("/", r)

	srv := &http.Server {
		Handler: r,
		Addr: "127.0.0.1:4000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
