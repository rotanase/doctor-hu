function make_api_request(request) {
    var xmlhttp = new XMLHttpRequest();

    xmlhttp.onreadystatechange = function() {
        if (xmlhttp.readyState == XMLHttpRequest.DONE) {
           if (xmlhttp.status == 200) {
           		const resp = JSON.parse(xmlhttp.responseText);
              var filtered;
              console.log(resp);
              //aici se face parsarea JSONULUI daca e nevoie si se apeleaza callbackuri daca sunt
              document.getElementById("response-box").innerHTML = filtered;
              document.getElementById("line").innerHTML = indicator;
           }
           else if (xmlhttp.status == 400) {
              alert('There was an error 400');
           }
           else {
              alert('something else other than 200 was returned');
           }
        } else {
          document.getElementById("response-box").innerHTML = `<i class="fa fa-gear fa-spin fa-3x fa-fw"></i>`;
        }
    };
    switch(request){
      //aici se decide ce parte a api-ului e apelat eg:filtru , topic nou... etc.
      case 'weather':
        xmlhttp.open("GET", `http://api.openweathermap.org/data/2.5/weather?q=${"Bucharest"}&APPID=${myWeatherAPIKey}`, true);
        break;
      case 'location':
        xmlhttp.open("GET", `http://ipinfo.io/json?token=${myIPAPIKey}`, true);
        break;
    }
    xmlhttp.send();
}
