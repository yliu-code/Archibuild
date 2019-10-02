var connectionMade = false;
setInterval(function(){
    if(connectionMade){
        connection.send(JSON.stringify({content:"</Alive>"}));
    }
},1000)


var connection = new WebSocket('ws://final160-alexanderpark.codeanyapp.com:2323/commun');
connection.onopen = function(){
    connectionMade=true;
    connection.send("client");
}
connection.onmessage = function(receivedMessage){
    let packet = JSON.parse(receivedMessage.data)
    //packet["text"]
    if(packet["statusOfOther"]=="dead"){
        console.log("dead on the other side")
    }else{
        let tokens = packet["content"].split("</Delimiter>")



        //var windowtab=window.open('about:blank','image from canvas');
        //windowtab.document.write("<img src='"+tokens[0]+"' alt='from canvas'/>");

        /*
        var link = document.createElement('a');
        link.download = "myPainting.png";
        link.href = tokens[0]
        link.click();
*/
    }


}
$("#chatBoxText").css("width","100%")
$("#chatBoxText").css("height","5%")
$("#chatBoxText").css("font-size","1vw")

$("#textToSend").css("width","100%")
$("#textToSend").css("height","45%")

$("#textToSend").attr("placeholder", "Type here and hit enter to send to Architect");
$("#textToSend").css("resize","none");
$("#textToSend").attr("rows","1");
$("#textToSend").attr("cols","1");



$("#textToSend").keypress(function(e){

    if(e.keyCode===13){
        let textToSend = $(this).val()
        packet = {content:textToSend}
        connection.send(JSON.stringify(packet))
        connection.send(JSON.stringify({content:""}));
        $("#userTextSoFar").val($("#userTextSoFar").val()+"\n-------------\n"+textToSend)
        $(this).val("")
        return false
    }
});

$("#userTextSoFar").css("width","100%")
$("#userTextSoFar").css("height","50%")
$("#userTextSoFar").css("resize","none");
$("#userTextSoFar").attr("disabled","disabled");
$("#userTextSoFar").val("Previous Messages to Architect:");
$("#userTextSoFar").css("background-color","#000")
$("#userTextSoFar").css("color","#fff")


//$("#userTextSoFar").css("border-color","#fff")

// Add floorplan images as the background
$(".bath").css( "background-color", "#020a01");
$(".designBoard>.tab").css("visibility","hidden");
$("#bath").css("visibility","visible");
$("ul>li").click(function() {
    $("ul>li").css( "background-color", "#c1c1c1");
    $(this).css( "background-color", "#020a01");
    console.log(this.className);
    $(".designBoard>.tab").css("visibility","hidden");
    $("#"+ this.className).css("visibility","visible");
    currentDiv = this.className;
});

// Add floorplan images as the background
$(".bath").css( "background-color", "#020a01");
$("#bath").prepend("<img id = 'bath_floorplan' class='floorplanImg' src = '/assets/imgs/floorplans/Bedroom.svg' width = '800px' height = '800px'/>");
$("#bed").prepend("<img id = 'bed_floorplan' class='floorplanImg' src = '/assets/imgs/floorplans/Bedroom.svg' width = '800px' height = '800px'/>");
$("#live").prepend("<img id = 'live_floorplan' class='floorplanImg' src = '/assets/imgs/floorplans/LivingRoom.svg' width = '800px' height = '800px'/>");
$("#kitchen").prepend("<img id = 'kitchen_floorplan' class='floorplanImg' src = '/assets/imgs/floorplans/Kitchen.svg' width = '800px' height = '800px'/>");
$(".designBoard>.tab").css("visibility","hidden");
$("#bath").css("visibility","visible");





