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
	//"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

// Type Product 
type Product struct {
	Name string
	Description string
	Price int
}


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

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://matheus:zurik21@ds237989.mlab.com:37989/ecommerce_catalog")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)


	if err != nil {
		log.Fatal(err)
	}

	// check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")

	//collection := client.Database("ecommerce_catalog").Collection("products")

	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed")

	// Router
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/products", ProductsHandler)
	r.HandleFunc("/products/{id}", ProductsIDHandler)
	http.Handle("/", r)

	srv := &http.Server {
		Handler: r,
		Addr: "127.0.0.1:5000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
