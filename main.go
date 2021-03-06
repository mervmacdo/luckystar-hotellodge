package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/mervmacdo/luckystar-hotellodge/morestrings"
)

type Welcome struct {
	Name string
	Time string
}

func main() {

	welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}

	templates := template.Must(template.ParseFiles("templates/welcome-template.html"))

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if name := r.FormValue("name"); name != "" {
			welcome.Name = name
		}

		if err := templates.ExecuteTemplate(w, "welcome-template.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Serving on port 8080")

	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello WOrld", "Hello Go"))

	fmt.Println(http.ListenAndServe(":8080", nil))
}
