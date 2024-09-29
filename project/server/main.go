package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type User struct {
	Email    string
	Password string
}

// func Home(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "index.html")

// }
// func receiveAjax(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "POST" {
// 		fmt.Println("______________________")
// 	}
// }

func main() {

	var user User = User{
		Email:    "-",
		Password: "-",
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		tmpl, _ := template.ParseFiles("index.html")
		tmpl.Execute(w, user)
	})
	// mux.HandleFunc("/receive", receiveAjax)
	// mux.HandleFunc("/", Home)
	mux.HandleFunc("/postform", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("-------------------")
		fmt.Println(r.FormValue("email"), r.FormValue("password"))
		user = User{
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		}

		text, _ := json.MarshalIndent(user, "", "\t")
		file, _ := os.Create("data.json")
		defer file.Close()
		file.Write(text)
		//http.Redirect(w, r, "/", http.StatusSeeOther)

	})
	mux.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "data.json")
	})

	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", mux)
}
