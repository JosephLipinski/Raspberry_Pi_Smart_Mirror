window.onload=setJsonObject;

var JsonResponse = null;
var JsonObject = null;

var setJsonObject = function(){
    $.get("http://localhost:8000/forecast", function(data, status){
      var JsonResponse = data;
      JsonObject = JSON.parse(JsonResponse);
      replaceText();
    });
}

setJsonObject();

function replaceText(){
    document.getElementById("Date").innerHTML=JsonObject.Forecast.Simpleforecast.Forecastday[0].Date.Weekday + ", " + JsonObject.Forecast.Simpleforecast.Forecastday[0].Date.Monthname + " " + JsonObject.Forecast.Simpleforecast.Forecastday[0].Date.Day;
    document.getElementById("Conditions").innerHTML=JsonObject.Forecast.Simpleforecast.Forecastday[0].Conditions
    document.getElementById("High").innerHTML=JsonObject.Forecast.Simpleforecast.Forecastday[0].High.Fahrenheit + " &#176;F"
    document.getElementById("Low").innerHTML=JsonObject.Forecast.Simpleforecast.Forecastday[0].Low.Fahrenheit + " &#176;F"

}