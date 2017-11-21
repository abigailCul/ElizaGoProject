
const form = $("#userInput");
const list = $("#output");

//Key function press enter to send your message
form.keypress(function(event){
    if(event.keyCode != 13){ 
        return;
    }

    event.preventDefault(); // Wont refresh

    const userInput = form.val(); // get the text from the input form and will clear your input box
    form.val(" "); 
    

    const userPrompt = {"input" : userInput }
    // Your input message will appear on the screen
    //Will see your message as You:
    list.append("<li class='list-group-item list-group-item'>"+"You : " + userInput + "</li>");

    // GET/POST
  //const queryParams = {"input" : userInput }
  $.get("/chat/", userPrompt)
        .done(responses => {
            const out = "<li class='list-group-item list-secondary' id='resp' >"+"ELiza : " + responses +  "</li>";
            setTimeout(function(){
                list.append(out)
            }, 1000);
        })
});