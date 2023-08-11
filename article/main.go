package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// mapping go lang to json
// caps at start
type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Content string `json:"Content"`
	Summary string `json:"Summary"`
	Author  string `json:"Author"`
}

func main() {
	http.HandleFunc("/article/", handleArticle)
	http.HandleFunc("/article1/", getArticles)
	fmt.Println("server running on port:2000")
	log.Fatal(http.ListenAndServe(":2000", nil))

}

func handleArticle(w http.ResponseWriter, req *http.Request) {
	// if req.URL.Path != "/" {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	fmt.Fprintf(w, "<h1>404 error:Page not Found<h1>")
	// } else {
	// 	fmt.Fprint(w, "<h1>welcome to http</h1>")
	// }
	if req.Method == "POST" {
		reqBody, _ := io.ReadAll(req.Body)
		var post Article
		// unmarshal converting json object into go lang object
		// rev is called marshaling

		err := json.Unmarshal(reqBody, &post)
		post.Author = "Thash"

		// newData,err2:=json.Marshal(post)
		if err != nil {
			fmt.Println("got error")
			fmt.Fprintf(w, err.Error())
		} else {
			json.NewEncoder(w).Encode(post) // sending json structure
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(w, "unable to process your request")
	}
}

func getArticles(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		reqBody, _ := io.ReadAll(req.Body)
		var articles [ ]Article
		err := json.Unmarshal(reqBody, &articles)
		if err != nil {
			fmt.Println("got error")
			fmt.Fprintf(w, err.Error())
		} else {
			fmt.Println(articles)
			//appeding the data here
			articles = append(articles, Article{Id: "3", Title: "Whereabouts", Content: "book", Summary: "good book", Author: "Jhumpa Lahiri"})
			json.NewEncoder(w).Encode(articles) // sending json structure
		}
	}
}

// json file
/*"Article":[{
  "Id":"1",
  "Title":"harry potter",
  "Content":"book",
  "Summary":"good book",
  "Author":"J K Rowling"},
  {
  "Id":"2",
  "Title":"wings of fire",
  "Content":"book",
  "Summary":"good book",
  "Author":"A P J"
  }]*/
