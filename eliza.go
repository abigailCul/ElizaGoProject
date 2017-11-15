/*
Abigail Culkin
Go page for Eliza chatBot
Ref: */

package main

import (
	"fmt"
	"net/http"
	"html/template"
	"log"
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
	m := Chat{Message: ""}
	
		t, _ := template.ParseFiles("chatPage.html")
	
    // checking for guess URL encoded variable
    userInput, err := r.URL.Query()["userInput"]
    // if not found execute the template and exit
    if !err || len(userInput) < 1 {
        log.Println("Url Param 'guess' is missing")
        // execute the template with the message
        t.Execute(w, m)
        return
    }// if
    
    // Query()["guess"] will return an array of items,
    // we only want the single item.
    str := userInput[0]
    
    //str = str + "</br>" + input.Input
    
    // adding the guess value to the template
    m = Chat{Message: "" + str}

}

func main() {

	// Adapted from: http://www.alexedwards.net/blog/serving-static-sites-with-go
	//fs:=  http.FileServer(http.Dir("chatPage"))
	//http.HandleFunc("/", userinputhandler)
	http.HandleFunc("/", templatehandler)

	http.HandleFunc("/chatPage", userinputhandler)
	http.ListenAndServe(":8080", nil)
}