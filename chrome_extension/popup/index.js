var tabId = 0;
chrome.tabs.query({active: true, currentWindow: true}, function(tabs) {
  tabId = tabs[0].id
});

console.log(tabId)

$("#clickme").on("click", function(){
  var message = $("#message").val()
  console.log(message)
  chrome.tabs.sendMessage(
    tabId,
    {message: message},
    function(response) {}
  )
  console.log("sent message")
});
