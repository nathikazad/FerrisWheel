ws = new WebSocket('ws://' + window.location.host + '/ws');
ws.addEventListener('message', function(e) {
    var msg = JSON.parse(e.data);
    if(msg.method == "spin") {
        rotateAnimation(msg.arg0);
    }
});

var looper;
function rotateAnimation(degrees){
    document.getElementById("Wheel").style.transform = "rotate("+degrees+"deg)";
    document.getElementById("Chair1").style.transform = "rotate(-"+degrees+"deg)";
    document.getElementById("Chair2").style.transform = "rotate(-"+degrees+"deg)";
    document.getElementById("Chair3").style.transform = "rotate(-"+degrees+"deg)";
    document.getElementById("Chair4").style.transform = "rotate(-"+degrees+"deg)";
    document.getElementById("Chair5").style.transform = "rotate(-"+degrees+"deg)";
    document.getElementById("Chair6").style.transform = "rotate(-"+degrees+"deg)";
    document.getElementById("Chair7").style.transform = "rotate(-"+degrees+"deg)";
    document.getElementById("Chair8").style.transform = "rotate(-"+degrees+"deg)";
    document.getElementById("degrees").innerHTML = degrees;
}

