package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//we have to create http handlers
	// url formation
	// // http.HandleFunc("/api/hello", handleHello)
	http.HandleFunc("/", handleHello)
	http.HandleFunc("/products/", handlePro)
	http.HandleFunc("/contact/", handleContact)

	// port no is important
	// http.ListenAndServe(":3000",nil)
	fmt.Println("server running on port:2000")
	//listening and serve the entire app
	//port will be mapped to domain name

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
func handlePro(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	fmt.Println(req.URL)
	fmt.Println(req.URL.Path)

	fmt.Fprint(w, "<h1>welcome to products</h1>")

}
func handleContact(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	fmt.Println(req.URL)
	fmt.Println(req.URL.Path)

	fmt.Fprint(w, "<h1>email:thashmigaa03@gmail.com</h1>")

}
