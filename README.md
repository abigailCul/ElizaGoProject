# ElizaGoProject
A go project for my third year project in Data Representation and Querying, using eliza chatbot. This project is a messenger chatbot that you can talk to and she will respond to you. The project is done using GoLang and made into a web browser using html to serve the go file. Once you send your message Eliza will respond after a second and the page wont refresh as i have used Ajax in my project

## Description
My Eliza chatbot once run in the browser comes up as a messenget page. The message is send using the return key and is sent from there to the list shown as You:, Eliza will then respond under Eliza:. If you use a word that she recognises she will have certain responses for that word, otherwise her responses will be random. 


## Prerequisites

I used github for my project so it would not be lost and be easy for other people to access.

### Push to Github:

In order to submit my project changes to github from my github folder i used the following commands:
git add .
git commit -m "Initial commit"
git push

### Download from github:
For you to download my project you must clone my repository link from the command promp:

git clone "example.github/project"

### You can then run my code using:
"go run eliza.go" 
You do this from the command once you are in the github folder that contains the project.
Then in the browser you run the page by using 127.0.0.1:8080 as the url for my project.

The only problem with my project is when you run it you habe to click chatPage.html to run eliza.
Im not sure how i was to fix this.

## Coding Syle

In my project i used GoLang, HTML and JavaScript. 
I used Ajax with go and javaScript in the following way. This also allowed me to serve the html file. 

Give an example

The src folder is the folder that contains my html and javaScript.
My userhandler is my function that gets the input from my html and when it is sent it will get Elizas response.
```
http.Handle("/", http.FileServer(http.Dir("./src")))
	http.HandleFunc("/chat/", userHandler)

	http.ListenAndServe(":8080", nil)
```
Used to link my js to my go file in order to run inputs and outputs.
```
$.get("/chat/", userPrompt)
```

## Resources i used for my project.

https://golang.org/doc/articles/wiki/
https://github.com/abigailCul/Go-Problem3/blob/master/Prob3.go
https://github.com/data-representation/go-ajax
https://gist.github.com/ianmcloughlin/c4c2b8dc586d06943f54b75d9e2250fe
https://www.reddit.com/r/golang/comments/2tmllh/sending_an_ajax_request_to_go_web_server/?st=jabfsgyz&sh=385af765
http://blog.teamtreehouse.com/create-ajax-contact-form 
https://www.w3schools.com/jquery/ajax_get.a
http://getbootstrap.com/

## License

This project is licensed under the Apache License- see the LICENSE.md file for details

  





