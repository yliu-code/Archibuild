var $welcomeText = $("<div id=\"welcomeText\">Welcome!</div>")
$welcomeText.hide()

var $questionText = $("<div id=\"internalText\" class=\"row\"> Are you a </div>")
$questionText.hide()

var $internalContainer = $("<div id=\"innerContainer\"class=\"row\"> </div>")

var $clientButton = $("<div class=\"col-md-5\">\n" +
    "                <button id=\"cButton\">Client</button>\n" +
    "            </div>")

var $orText = $("<div id=\"orText\" class=\"col-md-2\">\n" +
    "                <div class=\"vertical\">or</div>\n" +
    "            </div>")

var $archButton = $("<div class=\"col-md-5\">\n" +
    "                <button id=\"cButton\">Architect</button>\n" +
    "            </div>")

$clientButton.on('click', function(){
    location.href = '/search';
})

$archButton.on('click', function(){
    location.href = '/author';
})

$internalContainer.append($clientButton)
$internalContainer.append($orText)
$internalContainer.append($archButton)

$internalContainer.hide()


$("#totalPane").append($welcomeText)
$("#totalPane").append($questionText)
$("#totalPane").append($internalContainer)


function showQuestion(){

    $questionText.show("slide", { direction: "left" }, 1000,  function(){
        $internalContainer.fadeIn(2000)
    });

}

window.onload=function(){
    $welcomeText.show("slide", { direction: "left"}, 1000, showQuestion);

}
