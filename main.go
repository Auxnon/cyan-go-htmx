package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {

	http.Handle("/static",
		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./partials/index.html"))
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./partials/frags/results.html"))
		data := map[string][]Stock{
			"Results": SearchTicker(r.URL.Query().Get("key")),
		}
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/stock", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			ticker := r.PostFormValue("ticker")
			stock := SearchTicker(ticker)[0]
			val := 0. //GetDailyValues(ticker)`
			tmpl := template.Must(template.ParseFiles("./partials/index.html"))
			tmpl.ExecuteTemplate(w, "stock-element", Stock{Ticker: stock.Ticker, Name: stock.Name, Price: val})
		}
	})

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
