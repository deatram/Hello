package main

import (
	"fmt"
	"net/http"
	"html/template"
)

func main() {
	templates :=template.Must(template.ParseFiles("templates/index.html"))
	fmt.Println("Hello, Go Web Development")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		/*if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)*/
		}
		//fmt.Fprintf(w, "Hello,, Go Web Development")

	})
	fmt.Println(http.ListenAndServe(":9090", nil))

}