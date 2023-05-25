package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloServer)
	fmt.Println("Application running on port 8000...")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Simple app")
	log.Println(r.URL)
}
