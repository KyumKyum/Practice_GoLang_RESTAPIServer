package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"strconv"
	"sync"
	"github.com/gorilla/mux"
	"io/ioutil"
)

type Article struct {
	Id string `json:"Id"`
	Title string `json: "Title"`
	Description string `json: "Description"`
	Content string `json: "Content"`
}

var counter int
var mutex = &sync.Mutex{}

var Articles []Article //Article Var

func deleteArticle(w http.ResponseWriter, r *http.Request){
	//DELETE Method: Delete Designated Article
	vars := mux.Vars(r)
	key := vars["id"]

	for index, article := range Articles {
		if article.Id == key{
			Articles = append(Articles[:index], Articles[index+1:]...)
			fmt.Println("Key %s deleted", key)
		}
	}
}

func createArticle(w http.ResponseWriter, r *http.Request){
	//POST Method: Create New Articles

	reqBody, _ := ioutil.ReadAll(r.Body)
	//fmt.Fprintf(w,"%+v",string(reqBody)) Uncomment this before testing POST

	//Comment this area when testing fmt.Fprintf
	var article Article
	json.Unmarshal(reqBody, &article)
	//Update the global Article Array
	Articles = append(Articles, article)
	fmt.Println("Key Created")

	json.NewEncoder(w).Encode(article)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request){
	//GET MEthod: Read Single Article

	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Fprintf(w,"Current Key : %s\n", key)

	for _, article := range Articles {
		if article.Id == key{
			fmt.Println(article)
			json.NewEncoder(w).Encode(article)
		}
	}
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	//GET Method: Read All Articles

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
	myRouter.HandleFunc("/article",createArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}",deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}",returnSingleArticle)

	log.Fatal(http.ListenAndServe(":3000",myRouter))

}

func main(){
	fmt.Println("REST API - Mux Routers")
	fmt.Println("http://localhost:3000/")
	
	Articles = []Article{
		Article{Id: "1", Title: "This is the first title", Description: "This is the first description", Content: "This is the first content"},
		Article{Id: "2", Title: "This is the second title", Description: "This is the second description", Content: "This is the second content"},
		Article{Id: "50", Title: "This is the title for test", Description: "This is the description for test", Content: "This is the content for test"},
	}

	fmt.Println(Articles)
	
	handleRequests()
}