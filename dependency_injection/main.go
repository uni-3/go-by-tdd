package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	Greet(os.Stdout, "elodie\n")

	fmt.Println("server started on http://localhost:5000")
	if err := http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler)); err != nil {
		log.Fatal(err)
	}
}

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "hello, %s", name)
}

// wにgreetを書き込んで、rで返す
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}
