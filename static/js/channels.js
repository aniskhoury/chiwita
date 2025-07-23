

function showListChannels(){


  $.ajax({
        url: 'http://localhost:8080/listChannels',
    dataType: 'application/json',
    complete: function(data){
        console.log(data.responseText)
        var result = data.responseText;
        result = result.replaceAll('&#34;','"');
        console.log(result);
        var resultJSON = JSON.parse(result);
        console.log(resultJSON);
    },
    success: function(data){
        console.log(data)
    }
  })
}