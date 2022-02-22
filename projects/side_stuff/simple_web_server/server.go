package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

/***
 * This is a simple web server that can return a static file,
 * print hello, POST to a form and return, otherwise return 404
 ***/

func main() {
	// curl http://localhost:8080/
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// curl http://localhost:8080/hello
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	io.WriteString(w, "Hello!")
}

// curl --data "name=Louis&address=edina" http://localhost:8080/form
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

// This was the first main function I created was to get
// some main points down. It just handled a simple hello
// world function. Can ignore.
func example() {
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Starting server on port 8080")

	// If there is an error, it will be logged to fatal
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/***
 * Resources:
 *
 * curl help: https://www.freecodecamp.org/news/how-to-start-using-curl-and-why-a-hands-on-introduction-ea1c913caaaa/
 * source page: https://blog.logrocket.com/creating-a-web-server-with-golang/
 ***/
