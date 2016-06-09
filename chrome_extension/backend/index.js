chrome.runtime.onMessage.addListener(
  function(request, sender, sendResponse) {
    $("#issue_title").val(request.message);
  }
)
