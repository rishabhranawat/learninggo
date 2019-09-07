package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the Homepage!")
	fmt.Println("Endpoint Hit: homepage")
}

func returnName(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome Home rish!")
	fmt.Println("Endpoint Hit: returnName")
}

func returnGiveName(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	name := vars["name"]
	fmt.Fprintf(w, "Welcome home " + name)
	fmt.Println("Endpoint Hit: returnGiveName")
}

func handleRequests(){
	myRouter := mux.NewRouter().StrictSlash(true)

	// ordering has to be hierarchical
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/name", returnName)
	myRouter.HandleFunc("/name/{name}", returnGiveName)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main(){
	handleRequests()
}