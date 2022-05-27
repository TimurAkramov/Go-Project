package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "index", nil)
}

func create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "create", nil)
}

func authorization(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/authorization.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "authorization", nil)
	/*
		mail := r.FormValue("e-mail")
		password1 := r.FormValue("password1")
		db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/golang")
		if err != nil {
			panic(err)
		}
		defer db.Close()
		password := password1
		insert, err := db.Query(fmt.Sprintf("INSERT INTO 'GoProject'('mail', 'password') VALUES('%s', '%s')"), mail, password)
		if err != nil {
			panic(err)
		}
		defer insert.Close()

		http.Redirect(w, r, "/", 301)*/
}

func LogIn(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/LogIn.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "login", nil)
	/*
		mail := r.FormValue("e-mail")
		password := r.FormValue("password")
		db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/golang")
		checkIfExist, exists := db.Query(fmt.Sprintf("SELECT email FROM users WHERE EXISTS(SELECT user.email='%s' AND user.password='%s' )"), mail, password)
		if exists != nil {
			panic(exists)
		}
		if err != nil {
			panic(err)
		}
		defer checkIfExist.Close()
		defer db.Close()

		http.Redirect(w, r, "/", 301)*/
}

func creator(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/creator.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	t.ExecuteTemplate(w, "creator", nil)
}

func handleFunc() {
	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/creator", creator)
	http.HandleFunc("/authorization", authorization)
	http.HandleFunc("/authorization/login", LogIn)
	http.ListenAndServe(":8081", nil)
}

func main() {
	handleFunc()
}
