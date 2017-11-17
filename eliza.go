/*
Abigail Culkin
Go page for Eliza chatBot
Ref: */

package main

import (
	"fmt"
	"net/http"
	"html/template"
	//"log"
)
type Chat struct{
	Message string
}
// template message struct
/*type input struct {
    Input string
}// message_struct*/

func templatehandler(w http.ResponseWriter, r *http.Request){

	m := Chat{Message: ""}

	t, _ := template.ParseFiles("chatPage.html")

	t.Execute(w, m)

	r.ParseForm()

	//User input messages 
	fmt.Println(r.Form["userInput"])

  }// handler

func userinputhandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello, %s!", r.URL.Query().Get("value")) //.Path[1:])
/*	m := Chat{Message: ""}
	
		t, _ := template.ParseFiles("chatPage.html")
	
    // checking for guess URL encoded variable
    userInput := r.URL.Query().Get("userInput")
    // if not found execute the template and exit
    if  len(userInput) < 1 {
        log.Println("Url Param 'guess' is missing")
        // execute the template with the message
        t.Execute(w, m)
        return
	}// if*/
	
	userinput := r.URL.Query().Get("userInput")
	reply := Eliza(userinput)
	fmt.Fprintf(w, reply)
    

}

func Eliza(input string) string {
	//Code will be here for elizas responses
	
	return "Elizas response"
}

func main() {

	http.HandleFunc("/", templatehandler)

	http.HandleFunc("/chatPage", userinputhandler)
	http.ListenAndServe(":8080", nil)
}