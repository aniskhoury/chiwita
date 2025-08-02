
function openChannel(evt,nameChannel){
        var content = document.getElementById("channellist") ;

        content.innerHTML = "";
				var channel = document.createElement("div");
				channel.setAttribute("id",nameChannel);
				channel.setAttribute("class","channel");
                channel.innerHTML = nameChannel


				content.appendChild(channel);

        var i, tabcontent, tablinks;
        tabcontent = document.getElementsByClassName("tabcontent");
        for (i = 0; i < tabcontent.length; i++) {
            tabcontent[i].style.display = "none";
        }
        tablinks = document.getElementsByClassName("tablinks");
        for (i = 0; i < tablinks.length; i++) {
            tablinks[i].className = tablinks[i].className.replace(" active", "");
        }
        document.getElementById(nameChannel).style.display = "block";
        evt.currentTarget.className += " active";

}

function showListChannels(event){

  $.ajax({
        url: '/listChannels',
    dataType: 'application/json',
    complete: function(data){
        var result = data.responseText;
        result = result.replaceAll('&#34;','"');
        var resultJSON = JSON.parse(result);

        var linkToChannel = "";
        for (var i =0;i<resultJSON.length;i++){
            linkToChannel = linkToChannel +'<a id="myLink" onClick="openChannel(event,\''+resultJSON[i]["channel"]+'\');">'+resultJSON[i]["channel"]+" users online: "+resultJSON[i]["counterUsers"]+'</a><br>';
        }

         var content = document.getElementById("showChannels") ;
        if (document.getElementById("channellist") == null) {
            var listChannel = document.createElement("div");
            listChannel.setAttribute("id","channellist");
            listChannel.setAttribute("class","channel");
            listChannel.innerHTML = linkToChannel;
            content.appendChild(listChannel);
        }else{
            document.getElementById("channellist").innerHTML = linkToChannel;
        }

    },
    success: function(data){
        console.log(data)
    }
  })
}