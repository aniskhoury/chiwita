

function showListChannels(){


  $.ajax({
        url: '/listChannels',
    dataType: 'application/json',
    complete: function(data){
        console.log(data.responseText)
        var result = data.responseText;
        result = result.replaceAll('&#34;','"');
        console.log(result);
        var resultJSON = JSON.parse(result);
        console.log(resultJSON);
        var content = document.getElementById("showChannels") ;

        for (var i =0;i<resultJSON.length;i++){
            content.innerHTML =content.innerHTML+ resultJSON[i]["channel"]+" users online: "+resultJSON[i]["counterUsers"]+"<br>";
        }
    },
    success: function(data){
        console.log(data)
    }
  })
}