package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)


type Article struct{

	//P Person `json:"Person"`
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json: "content"`


}

type Articles []Article

func AllArticle(w http.ResponseWriter,r *http.Request){
	articles := Articles{
		Article{Title:"title",Desc: "Description",Content: "Hallow world"},
	}

	fmt.Println("All article endpoint hit")
	json.NewEncoder(w).Encode(articles)

}




func main() {

	handleReq()


}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is the response of home page")
}

func testPostReaquest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is the response of Tests post page")
}

func handleReq(){

	myRouter:=mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/",homePage)
	myRouter.HandleFunc("/article",AllArticle).Methods("GET")
	myRouter.HandleFunc("/article",testPostReaquest).Methods("POST")
	http.ListenAndServe(":8080", myRouter)
}