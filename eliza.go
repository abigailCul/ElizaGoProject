/*
Abigail Culkin
Go page for Eliza chatBot
Ref: */

package main

import (
	"fmt"
	"net/http"
	//"html/template"
	//"log"
	"regexp"
	"math/rand"
"time"
//"io"

)
type Chat struct{
	Message string
}
var Intro = []string{
    "Hello. How are you feeling today?",
    "How are you? Do you want to talk",
    "Is there something bothering you?",
    "How about we have a little chat?",
}

  var Help = map[string]string{
	"am": "are",
    "was": "were",
    "i": "you",
    "i'd": "you would",
    "i've": "you have",
    "i'll": "you will",
    "my": "your",
    "are": "am",
    "you've": "I have",
    "you'll": "I will",
    "your": "my",
    "yours": "mine",
    "you": "me",
	"me": "you",
}
var Bye = []string{
    "goodbye",
    "bye",
    "Please come back and talk soon",
	"It was nice to talk to you",

}
var response = []string{
	"I’m not sure what you’re trying to say. Could you explain it to me?",
	"How does that make you feel?",
	"Why do you say that?",
	"I really like talking to you",
	"Please tell me more.",
	"Very interesting.",
  }

  
func userHandler(w http.ResponseWriter, r * http.Request) {

			input := r.URL.Query().Get("input")
			out := Eliza(input)
			fmt.Fprintf(w, out)


} // Handler
  

func Eliza(inputMessage string) string {
	
	
	if matched, _ := regexp.MatchString(`(?i).*\bfather\b.*`, inputMessage); matched{
		return "Why dont you tell me about your Father?"
		}
		//(?i) case insensitive

		if matched, _ := regexp.MatchString(`(?i).*\bmother\b.*`, inputMessage); matched{
			return "Why dont you tell me about your mother?"
			}
   
	  re := regexp.MustCompile("(?i)I am ([^.!?]*)[.!?]?")
	   if re.MatchString(inputMessage){
	 
	  return re.ReplaceAllString(inputMessage, "How do you know you are $1?")
   
	   }

	   rep := regexp.MustCompile("(?i)I have ([^.!?]*)[.!?]?")
	   if rep.MatchString(inputMessage){
	 
	  return rep.ReplaceAllString(inputMessage, "Why dont you tell me more about your $1?")
   
	   }

	   name := regexp.MustCompile("(?i)Hi my name is ([^.!?]*)[.!?]?")
	   if name.MatchString(inputMessage){
	 
	  return name.ReplaceAllString(inputMessage, "Nice to meet you $1!")
   
	   }
	 

	res := regexp.MustCompile("(?i)I really like ([^.!?]*)[.!?]?")
	if res.MatchString(inputMessage){
  
   return res.ReplaceAllString(inputMessage, "Why do you really like $1?")

	}
	return response[rand.Intn(len(response))]

}//Close Eliza function

func ElizaHi() string {
    return  Intro[rand.Intn(len(Intro))]
}
func ElizaBye() string {
    return  Bye[rand.Intn(len(Bye))]
}
func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	http.Handle("/", http.FileServer(http.Dir("./src")))
	http.HandleFunc("/chat/", userHandler)

	http.ListenAndServe(":8080", nil)
}
