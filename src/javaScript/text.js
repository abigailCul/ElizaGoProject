
/*
Abigail Culkin
G00334291
Ref:
https://www.reddit.com/r/golang/comments/2tmllh/sending_an_ajax_request_to_go_web_server/?st=jabfsgyz&sh=385af765
http://blog.teamtreehouse.com/create-ajax-contact-form 
https://www.w3schools.com/jquery/ajax_get.asp
 */
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

 //get chat links with my main function in go for user handler i have /chat/
  $.get("/chat/", userPrompt)
        .done(responses => {
            const out = "<li class='list-group-item list-group-item-secondary' id='resp' >"+"ELiza : " + responses +  "</li>";
            setTimeout(function(){
                list.append(out)
            }, 1000);
        })

        //function to help scroll but does not work
        function scrollToBottom () {
            document.body.scrollTop = document.body.scrollHeight;
         }
});


