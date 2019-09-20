package main

import (
	"net/http"
	"fmt"
	"time"
	"html/template"
)

type Welcome struct {
	Name string
	Time string
}

func main() {


	templates := template.Must(template.ParseFiles("template/welcome-template.html"))

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		welcome := Welcome{"Hakan", time.Now().Format(time.Stamp)}

		if name := r.FormValue("name"); name != "" {
			welcome.Name = name
		}

		if err := templates.ExecuteTemplate(w, "welcome-template.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
