package main

import (
	"fmt"
	"net/http"
	"html/template"
)
type Chat struct{
	Message string
}
func templatehandler(w http.ResponseWriter, r *http.Request){

	m := Chat{Message: "Eliza"}

	t, _ := template.ParseFiles("chatPage.html")

	t.Execute(w, &m)

	r.ParseForm()

	//User input messages

	fmt.Println(r.Form["userInput"])

}

func userinputhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Query().Get("value")) //.Path[1:])
}

func main() {

	// Adapted from: http://www.alexedwards.net/blog/serving-static-sites-with-go
	//fs:=  http.FileServer(http.Dir("chatPage"))
	//http.HandleFunc("/", userinputhandler)
	http.HandleFunc("/", templatehandler)

	//http.HandleFunc("/chatPage", userinputhandler)
	http.ListenAndServe(":8080", nil)
}