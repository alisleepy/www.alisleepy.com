$(function(){
    var bId = $("#bId").val();
    getTicketInfo(bId);
});

function getTicketInfo(bId){
    var url = "/home/getBlogInfo";
    $.get(url, function (json) {
        if(json.code == 200){

        }else{
            
        }
    }, "json");
}