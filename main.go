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
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	// Type Product
	type Product struct {
	 ID	primitive.ObjectID `bson:"_id,omitempty"`
	 Name	string	`bson:"name,omitempty"`
	 Description	string `bson:"description,omitempty"`
	 Price	int	`bson:"price,omitempty"`
	}

	// Connect Database
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://matheus:zurik21@ds237989.mlab.com:37989/ecommerce_catalog?retryWrites=false"))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	// Database Name
	database := client.Database("ecommerce_catalog")

	// Instance Product
	product := Product{
		Name: "Frauda",
		Description: "Fraudas descartaveis Tam: P,M,G",
		Price: 10,
	}

	// Collection Name
	productsCollection := database.Collection("products")

	// Insert One Product
	insertResult, err := productsCollection.InsertOne(ctx, product)
	if err != nil {
		panic(err)
	}
	fmt.Println(insertResult.InsertedID)

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
