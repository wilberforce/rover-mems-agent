package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pkg/browser"
)

func main() {

	fs := http.FileServer(http.Dir("./web-static"))
	http.Handle("/", fs)

	fmt.Println("Server listening on port 8080")
	log.Panic(
		http.ListenAndServe(":8080", nil),
	)
	browser.OpenURL("http://localhost:8080/index.html")
}