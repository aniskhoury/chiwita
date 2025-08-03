
function openChannel(evt,nameChannel){
        $(".userList").innerHTML = "hola";
        $(".userList").hide();
        $(".textArea").hide();
        var content = document.getElementById("channellist") ;
        $(".channelList").hide();

        if (document.getElementById("userList"+nameChannel)==null){

            document.getElementById("tab").innerHTML = document.getElementById("tab").innerHTML+ '<button id = "'+nameChannel+'" class="tablinks" onclick="openChannel(event,\''+nameChannel+'\')">'+nameChannel+'</button>';

            var userList = document.createElement("div");
            var textArea = document.createElement("div");

            userList.setAttribute("id","userList"+nameChannel);
            textArea.setAttribute("id","textArea"+nameChannel);
            userList.setAttribute("class","userList");
            textArea.setAttribute("class","textArea");

            document.getElementById("elements").appendChild(userList);
            document.getElementById("elements").appendChild(textArea);
            //userList.style = "background-color:blue";
            //textArea.style = "background-color:red";
        }else{
            $("#userList"+nameChannel).show();
            $("#textArea"+nameChannel).show();
        }
 

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
        $(".userList").innerHTML = "hola";
        $(".userList").hide();
        $(".textArea").hide();
        $(".channelList").hide();

        var result = data.responseText;
        result = result.replaceAll('&#34;','"');
        var resultJSON = JSON.parse(result);

        var linkToChannel = "";
        for (var i =0;i<resultJSON.length;i++){
            linkToChannel = linkToChannel +'<a id="myLink" onClick="openChannel(event,\''+resultJSON[i]["channel"]+'\');">'+resultJSON[i]["channel"]+" users online: "+resultJSON[i]["counterUsers"]+'</a><br>';
        }

        if (document.getElementById("channelList")==null){
            var channelList = document.createElement("div");
            channelList.setAttribute("id","channelList");
            channelList.setAttribute("class","channelList");
            document.getElementById("content").appendChild(channelList);
            document.getElementById("channelList").innerHTML = linkToChannel; 
        }else{
            document.getElementById("channelList").innerHTML = linkToChannel;    
            $(".channelList").show();
        }
        
         


        

    },
    success: function(data){
        console.log(data)
    }
  })
}