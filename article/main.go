package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

//mapping go lang to json
// caps at start
type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Content string `json:"Content"`
	Summary string `json:"Summary"`
	Author  string `json:"Author"`
}

func main() {

	http.HandleFunc("/", handleHello)
	http.HandleFunc("/contact/", handleContact)
	http.HandleFunc("/article/", handleArticle)
	fmt.Println("server running on port:2000")
	log.Fatal(http.ListenAndServe(":2000", nil))

}

// parameters are response and pointer ref request
func handleHello(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	fmt.Println(req.URL)
	fmt.Println(req.URL.Path)
	if req.URL.Path != "/" {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "<h1>404 error:Page not Found<h1>")
	} else {
		fmt.Fprint(w, "<h1>welcome to http</h1>")
	}

}
func handleArticle(w http.ResponseWriter, req *http.Request) {
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
func handleContact(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	fmt.Println(req.URL)
	fmt.Println(req.URL.Path)
	fmt.Fprint(w, "<h1>email:thashmigaa03@gmail.com</h1>")

}

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