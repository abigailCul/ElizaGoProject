# ElizaGoProject
A go project for my third year project in Data Representation and Querying, using eliza chatbot. This project is a messenger chatbot that you can talk to and she will respond to you. The project is done using GoLang and made into a web browser using html to serve the go file. Once you send your message Eliza will respond after a second and the page wont refresh as i have used Ajax in my project

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

You can then run my code using "go run eliza.go" from the command once you are in the github folder that contains the project.
Then in the browser you run the page by using 127.0.0.1:8080 as the url for my project.

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


## License

This project is licensed under the MIT License - see the LICENSE.md file for details

  





