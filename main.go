package main

import (
	"fmt"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p1 := Person{Name: "Nick", Age: 32}
		fmt.Fprint(w, "Hello, ", p1.Name)
	})
	http.ListenAndServe(":4000", nil)
}
