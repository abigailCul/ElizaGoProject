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
	"regexp"
	"math/rand"
"time"
)
type Chat struct{
	Message string
}
var responses = []string{
	"I’m not sure what you’re trying to say. Could you explain it to me?",
	"How does that make you feel?",
	"Why do you say that?",
  }
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

//Previous tries to get my user input
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
	if matched, _ := regexp.MatchString(`(?i).*\bfather\b.*`, input); matched{
		return "Why dont you tell me about your Father?"
		}
		//(?i) case insensitive

		if matched, _ := regexp.MatchString(`(?i).*\bmother\b.*`, input); matched{
			return "Why dont you tell me about your mother?"
			}
   
	  re := regexp.MustCompile("I am ([^.!?]*)[.!?]?")
	   if re.MatchString(input){
	 
	  return re.ReplaceAllString(input, "How do you know you are $1?")
   
	   }

	   pattern := "name is (.*)"

	   rep := regexp.MustCompile(pattern)

	   if rep.MatchString(input) { // the input and regular expression match.
		match := rep.FindStringSubmatch(input)
		name := match[1]
		return "Hello " + name + " it's nice to meet you."
	}

	res := regexp.MustCompile("I really like ([^.!?]*)[.!?]?")
	if res.MatchString(input){
  
   return res.ReplaceAllString(input, "Why do you really like $1?")

	}

	//Code will be here for elizas responses
/*pattern := "name is (.*)"
	likePattern := "I like (.*)"

	re := regexp.MustCompile(pattern)

	if re.MatchString(input) { // the input and regular expression match.
		match := re.FindStringSubmatch(input)
		name := match[1]
		return "Hello " + name + " it's nice to meet you."
	}
	re = regexp.MustCompile(likePattern)
	if re.MatchString(input) { // the input and regular expression match.
		match := re.FindStringSubmatch(input)
		like := match[1]
		return "Why do you like " + like + "?"
	}

return "And how does that make you feel?"*/
	return responses[rand.Intn(len(responses))]
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	http.HandleFunc("/", templatehandler)

	http.HandleFunc("/chatPage", userinputhandler)
	http.ListenAndServe(":8080", nil)
}