package main

import (
	"net/http"
	"fmt"
	"log"
	"time"
	"github.com/gorilla/mux"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Index: %w\n", vars["Handler"])
}

func TasksHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Tasks: %w\n", vars["refactory code"])
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/tasks", TasksHandler)
	http.Handle("/", r)

	srv := &http.Server {
		Handler: r,
		Addr: "127.0.0.1:4000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
