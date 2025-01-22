package main

import (
	"customermod/controllers"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/api/issues/", controllers.GetIssues).Methods("GET")
	mux.HandleFunc("/api/findissue", controllers.FindIssue).Methods("GET").Queries("issueId",  "{issueId}")
	mux.HandleFunc("/api/createissue/", controllers.CreateIssue).Methods("POST")
	mux.HandleFunc("/api/deleteissue", controllers.DeleteIssue).Methods("DELETE").Queries("issueId",  "{issueId}")
	fmt.Println("server initiated")
	log.Fatal(http.ListenAndServe(":3000", mux))
}