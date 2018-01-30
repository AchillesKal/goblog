package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fm, _ := GetFlash(w, r, "message")

	if fm == nil {
		fmt.Fprint(w, "No flash messages")
		return
	}
	fmt.Fprintf(w, "%s", fm)

	t, _ := template.ParseFiles("templates/base_front.gtpl", "templates/index.gtpl")
	t.Execute(w, nil)
}

func about(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/base_front.gtpl", "templates/about.gtpl")
	t.Execute(w, nil)
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/base_front.gtpl", "templates/contact.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Println("username:", r.FormValue("email"))
		email := []byte(r.FormValue("email"))

		SetFlash(w, "message", email)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func admin(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/admin.gtpl")
	t.Execute(w, nil)
}

func addPost(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/add_post.gtpl")
	t.Execute(w, nil)
}

func main() {
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)

	http.HandleFunc("/admin", admin)
	http.HandleFunc("/admin/add/post", addPost)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
