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
var response = []string{
	"I’m not sure what you’re trying to say. Could you explain it to me?",
	"How does that make you feel?",
	"Why do you say that?",
  }
  
func userHandler(w http.ResponseWriter, r * http.Request) {
	
	
	
			input := r.URL.Query().Get("input")
			out := Eliza(input)
			fmt.Fprintf(w, out)

} // chatHandler
/*
func templatehandler(w http.ResponseWriter, r *http.Request){
	//fmt.Println("input:", r.URL.Query()); 

	m := Chat{Message: ""}

	t, _ := template.ParseFiles("chatPage.html")

	t.Execute(w, m)

	//r.ParseForm()

	//User input messages 
	//fmt.Println(r.Form["userInput"])
	io.WriteString(w, "Eliza: "+ r.FormValue("input"))


		input := r.URL.Query().Get("input")
        out := Eliza(input)
		fmt.Fprintf(w, out)

  }// handler
  */
  

func Eliza(input string) string {
	
	if matched, _ := regexp.MatchString(`(?i).*\bfather\b.*`, input); matched{
		return "Why dont you tell me about your Father?"
		}
		//(?i) case insensitive

		if matched, _ := regexp.MatchString(`(?i).*\bmother\b.*`, input); matched{
			return "Why dont you tell me about your mother?"
			}
   
	  re := regexp.MustCompile("I am (?i)([^.!?]*)[.!?]?")
	   if re.MatchString(input){
	 
	  return re.ReplaceAllString(input, "How do you know you are $1?")
   
	   }

/*	   pattern := "name is (.*)"

	   rep := regexp.MustCompile(pattern)

	   if rep.MatchString(input) { // the input and regular expression match.
		match := rep.FindStringSubmatch(input)
		name := match[1]
		return "Hello " + name + " it's nice to meet you."
	}*/

	res := regexp.MustCompile("I really like ([^.!?]*)[.!?]?")
	if res.MatchString(input){
  
   return res.ReplaceAllString(input, "Why do you really like $1?")

	}
	return response[rand.Intn(len(response))]

}

func main() {
		


//	fmt.Println()
//	fmt.Println(Eliza(response[rand.Intn(len(response))]))

	rand.Seed(time.Now().UTC().UnixNano())

	//http.HandleFunc("/chat/", templatehandler)
	http.Handle("/", http.FileServer(http.Dir("./src")))
	http.HandleFunc("/chat/", userHandler)

	http.ListenAndServe(":8080", nil)
}
/*func getInput() string {
    fmt.Print("You: ")
    reader := bufio.NewReader(os.Stdin)
    userinput, _ := reader.ReadString('\n')
    return userinput
}*/