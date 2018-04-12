ws = new WebSocket('ws://' + window.location.host + '/ws');
ws.addEventListener('message', function(e) {
    var msg = JSON.parse(e.data);
    if(msg.method == "spin") {
        rotateAnimation("Wheel",10);
    }
});

var looper;
var degrees = 0;
var time = 0;
function rotateAnimation(el,speed){
    var elem = document.getElementById(el);
    elem.style.transform = "rotate("+degrees+"deg)";
    document.getElementById("Chair1").style.transform = "rotate(-"+degrees+"deg)";
    document.getElementById("Chair2").style.transform = "rotate(-"+degrees+"deg)";
    document.getElementById("Chair3").style.transform = "rotate(-"+degrees+"deg)";
    document.getElementById("Chair4").style.transform = "rotate(-"+degrees+"deg)";
    document.getElementById("Chair5").style.transform = "rotate(-"+degrees+"deg)";
    document.getElementById("Chair6").style.transform = "rotate(-"+degrees+"deg)";
    document.getElementById("Chair7").style.transform = "rotate(-"+degrees+"deg)";
    document.getElementById("Chair8").style.transform = "rotate(-"+degrees+"deg)";
    degrees += 1;
    if(degrees > 359){
        degrees = 1;
    }
    time++;
    if(time < 1080) {
        looper = setTimeout('rotateAnimation(\''+el+'\','+speed+')',speed);
    } else {
        time = 0;
        degrees = 0;
    }
}

