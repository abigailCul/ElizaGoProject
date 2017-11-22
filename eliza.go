/*
Abigail Culkin
Go page for Eliza chatBot
Ref:
 https://golang.org/doc/articles/wiki/
https://github.com/abigailCul/Go-Problem3/blob/master/Prob3.go
https://github.com/data-representation/go-ajax
https://gist.github.com/ianmcloughlin/c4c2b8dc586d06943f54b75d9e2250fe
*/

package main

//The imports i have used in my go 
import (
	"fmt"
	"net/http"
	"regexp"
	"math/rand"
	 "time"
	 "strings"

)
//If you ask Eliza "How are you? She will answer with the following"
var answers = []string {
	"Im really good. You?",
	"Im not feeling good today!",
	"I couldnt be better!!",
	"I have had better days..",
}// answers

//Same as above if you say bye or goodbye or time so say bye
var Bye = []string {
	"Bye, nice to talk to you..",
	"Goodbye, I hope you come back soon..",
	"I enjoyed your comany, talk soon!",
	"See you soon I hope!",
}// Bye
//Her many responses if you say hi or hello
var Hi = []string {
	"Hello, How are you?",
	"Hi, what is your name?",
	"Hiya, Want to talk?",
}// Hi

var Sorry = []string {
	"Please don't apologise!!",
	"Why are you sorry?",
	"You dont need to be sorry!!",
}// Sorry
var Thankyou = []string {
	"You are very welcome!",
	"Thats no problem, anytime!",
}// Sorry

//Elizas random responses to your messages
var response = []string{
	"I’m not sure what you’re trying to say. Could you explain it to me?",
	"How does that make you feel?",
	"Why do you say that?",
	"I really like talking to you",
	"Please tell me more.",
	"Very interesting.",
	"You seem like a smart person.",
	"I hope you are enjoying my company",
  }

  //Function that links to javascript and recognises your input from the text box
//It will then release her output.
//Used response and request from the Go language to be able to talk back and forth
func userHandler(w http.ResponseWriter, r * http.Request) {

			input := r.URL.Query().Get("input")
			out := Eliza(input)
			fmt.Fprintf(w, out)


} // Handler
  
//This function is where Elizas responses are run
//I have used regular expressions so she will recognise different words and this will trigger certain responses
func Eliza(inputMessage string) string {
	
	
	if matched, _ := regexp.MatchString(`(?i).*\bfather\b.*`, inputMessage); matched{
		return "Why dont you tell me about your Father?"
		}
		//(?i) case insensitive

		if matched, _ := regexp.MatchString(`(?i).*\bmother\b.*`, inputMessage); matched{
			return "Why dont you tell me about your mother?"
			}

	if matched, _ := regexp.MatchString(`(?i).*\bsister\b.*`, inputMessage); matched{
		return "Do you have a sister? Tell me more?"
		}

		if matched, _ := regexp.MatchString(`(?i).*\bbrother\b.*`, inputMessage); matched{
			return "Do you have a brother? Tell me more?"
			}
		
			rs := regexp.MustCompile("(?i)how are you(.*)")
			if rs.MatchString(inputMessage){

				return answers[rand.Intn(len(response))]
		}// if

		hiMes := regexp.MustCompile("(?i)hi(.*)")
		if hiMes.MatchString(inputMessage){

			return Hi[rand.Intn(len(response))]
	}// if

	hello := regexp.MustCompile("(?i)hello(.*)")
	if hello.MatchString(inputMessage){

		return Hi[rand.Intn(len(response))]
}// if

	bye := regexp.MustCompile("(?i)bye(.*)")
	if bye.MatchString(inputMessage){

		return Bye[rand.Intn(len(response))]
	}// if

	sor := regexp.MustCompile("(?i)sorry(.*)")
	if sor.MatchString(inputMessage){

		return Sorry[rand.Intn(len(response))]
	}// if
	
	thank := regexp.MustCompile("(?i)thank(.*)")
	if thank.MatchString(inputMessage){

		return Thankyou[rand.Intn(len(response))]
	}// if
	
	  re := regexp.MustCompile("(?i)I am ([^.!?]*)[.!?]?")
	   if re.MatchString(inputMessage){
	 
	  return re.ReplaceAllString(inputMessage, "How do you know you are $1?")
   
	   }

	   rep := regexp.MustCompile("(?i)I have ([^.!?]*)[.!?]?")
	   if rep.MatchString(inputMessage){
	 
	  return rep.ReplaceAllString(inputMessage, "Why dont you tell me more about your $1?")
   
	   }

	   name := regexp.MustCompile("(?i)My name is ([^.!?]*)[.!?]?")
	   if name.MatchString(inputMessage){
	 
	  return name.ReplaceAllString(inputMessage, "Nice to meet you $1, want to talk?!")
   
	   }

	   res := regexp.MustCompile("(?i)I like(.*)")
	   subString := res.FindStringSubmatch(inputMessage)
	   
	   if len(subString) > 1 {
		prounoun := helper(subString[1])
		   return "Why do you like " + prounoun + "?"
   }// if
	return response[rand.Intn(len(response))]

}//Close Eliza function

func helper(prounoun string) string {
	// replacing the full stop with a question mark
	prounoun = strings.Replace(prounoun, ".", "?", 1)
	// reflecting the pronouns in the captured groups
	prounoun = strings.Replace(prounoun, "you’re", "I'm", 1)
	prounoun = strings.Replace(prounoun, "your", "my", 1)
	prounoun = strings.Replace(prounoun, "you", "me", 1)
	prounoun = strings.Replace(prounoun, "I", "you", 1)
	prounoun = strings.Replace(prounoun, "i am", "you are", 1)
	return prounoun
}// helper func

//My main function that uses Ajax and serves the HTML file and connects java script
func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	http.Handle("/", http.FileServer(http.Dir("./src")))
	http.HandleFunc("/chat/", userHandler)

	http.ListenAndServe(":8080", nil)
}
