package main

import "fmt"
import "net/http"

type struct WordCollection {}

func words(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func main() {
	http.HandleFunc("/get_words", words)

	http.ListenAndServe(":8090", nil)
}
