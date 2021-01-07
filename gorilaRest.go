package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"strconv"
	"sync"
	"github.com/gorila/mux"
)

type Article struct {
	id string `json:"id"`
	title string `json: "title"`
	description string `json: "description"`
	content string `json: "content"`
}

var counter int
var mutex = &sync.Mutex{}

var Articles []Article //Article Var

func returnSingleArticle(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Fprintf(w,"Current Key : %s\n", key)

	for _, article := range Articles {
		if article.id == key{
			fmt.Println(article)
			json.NewEncoder(w).Encode(article)
		}
	}
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Return All Articles")
	fmt.Println(Articles)
	json.NewEncoder(w).Encode(Articles)

}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}

func showHomepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Go test server!")
	fmt.Println("Endpoint Hit: Homepage of Go test server")
}

func handleRequests(){
	//Using Gorila/mux router to handle requests
	//Creating new instance of mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	//Function Handling
	myRouter.HandleFunc("/",showHomepage)
	myRouter.HandleFunc("/increment",incrementCounter)
	myRouter.HandleFunc("/articles",returnAllArticles)
	myRouter.HandleFunc("/article/{id}",returnSingleArticle)

	log.Fatal(http.ListenAndServe(":3000",myRouter))

}

func main(){
	fmt.Println("REST API - Mux Routers")
	fmt.Println("http://localhost:3000/")
	
	Articles = []Article{
		Article{id: "1", title: "This is the first title", description: "This is the first description", content: "This is the first content"},
		Article{id: "2", title: "This is the second title", description: "This is the second description", content: "This is the second content"},
		Article{id: "50", title: "This is the title for test", description: "This is the description for test", content: "This is the content for test"},
	}

	fmt.Println(Articles)
	
	handleRequests()
}