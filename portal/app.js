
var userName;
var questions;
$('#newQuestion').on('show.bs.modal', function (event) {
        var button = $(event.relatedTarget) // Button that triggered the modal// Extract info from data-* attributes
        // If necessary, you could initiate an AJAX request here (and then do the updating in a callback).
        // Update the modal's content. We'll use jQuery here, but you could use a data binding library or other methods instead.
      })

$.getJSON('https://cors-anywhere.herokuapp.com/'+'http://18.195.43.151:8080/questions', function(data){
  for(var question in data) {
    /*$.getJSON('https://cors-anywhere.herokuapp.com/'+'http://18.195.43.151:8080/users/' + data[question]["user_id"], function(user){
      userName = user["firstname"] + " " + user["lastname"];
    });*/
    //$("#content-body").append("<tr><td>" + user["firstname"] + " " + user["lastname"] + "</td><td>" + data[question]["content"] + '</td><td><button class="btn btn-outline-success btn-space" type="button" data-toggle="modal" data-target="#newAnswer" data-whatever="@mdo">Answer</button></td></tr>');
    $("#content-body").append("<tr><td>" + userName + "</td><td><p class='links' id='p"+data[question]["id"]+"'>" + data[question]["content"] + '</p></td><td><button class="btn btn-outline-success btn-space" type="button" data-toggle="modal" data-target="#newAnswer" data-whatever="@mdo">Answer</button></td><td class="hidden"></td></tr>');
  }
});

$("#questionSubmit").click(function(){
  var content = $("#question-text").val();
  var b = JSON.stringify({'user_id': '2', 'content': content});
  $.ajax({
    contentType: false,
    data: b,
    dataType: 'json',
    success: function(data){
        console.log(data);
    },
    error: function(data){
        console.log(data);
    },
    processData: false,
    type: 'POST',
    url: 'https://cors-anywhere.herokuapp.com/'+ "http://18.195.43.151:8080/questions/70"
});
  /*$.post('https://cors-anywhere.herokuapp.com/'+ "http://18.195.43.151:8080/questions/70", {"user_id":"1","content":"test question 71"}, function(result){
    console.log(result);
  });*/
  $("#question-text").val('');
//});
});

$("#answerSubmit").click(function(){
  var content = $("#answer-text").val();
  var b = JSON.stringify({'user_id': '2', 'content': content});
  $.ajax({
    contentType: false,
    data: b,
    dataType: 'json',
    success: function(data){
        console.log(data);
    },
    error: function(data){
        console.log(data);
    },
    processData: false,
    type: 'POST',
    url: 'https://cors-anywhere.herokuapp.com/'+ "http://18.195.43.151:8080/responses/71"
});
  $("#answer-text").val('');
});

$('#p1').click(function() {
  console.log("whatever");
  /*$.getJSON('https://cors-anywhere.herokuapp.com/'+'http://18.195.43.151:8080/response/' + this.id, function(data){
      $("hidden").addClass("unhidden").removeClass("hidden");
   });*/
});