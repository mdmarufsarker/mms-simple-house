package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)
const (
	CONN_HOST = "localhost"
	CONN_PORT = "5000"
)

func home(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		parsedTemplate, _ := template.ParseFiles("templates/index.html")
		parsedTemplate.Execute(w, nil)
	}
}
func about(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		parsedTemplate, _ := template.ParseFiles("templates/about.html")
		parsedTemplate.Execute(w, nil)
	}
}
func contact(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		parsedTemplate, _ := template.ParseFiles("templates/contact.html")
		parsedTemplate.Execute(w, nil)
	}
}

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/about", about).Methods("GET")
	router.HandleFunc("/contact", contact).Methods("GET")
	
	router.PathPrefix("/").Handler(http.StripPrefix("/public",
	http.FileServer(http.Dir("public/"))))

	fmt.Println("Running on http://localhost:5000/")
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
	if err != nil{
		log.Fatal("error starting http server : ", err)
		return
	}
}