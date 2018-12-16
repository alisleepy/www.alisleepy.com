$(function(){
    var bId = $("#bId").val();
    console.log(bId);
    getTicketInfo(bId);
});

function getTicketInfo(bId){
    var url = "/home/getBlogInfo";
    $.get(url, {bId:bId}, function (json) {
        if(json.code == 200){
            var blog = json.data;
            var blogInfo_html_str = '';
            blogInfo_html_str +=
                '<h1 class="article-title" id="bTitle">'+blog.bTitle+'</h1>' +
                '<div class="article-meta">' +
                '    <span class="item article-meta-time">' +
                '       <time class="time" data-toggle="tooltip" data-placement="bottom" title="'+getLocalTime(blog.add_time)+'">' +
                '            <i class="glyphicon glyphicon-time" id="add_time"></i>'+getLocalTime(blog.add_time) +
                '        </time>' +
                '    </span>' +
                '    <span class="item article-meta-source" data-toggle="tooltip" data-placement="bottom" title="添加者：'+blog.AName+'">' +
                '        <i class="glyphicon glyphicon-globe" id="aId"></i>'+blog.AName +
                '    </span>' +
                '    <span class="item article-meta-category" data-toggle="tooltip" data-placement="bottom" title="分类：'+blog.catName+'">' +
                '        <i class="glyphicon glyphicon-list" id="bCategory"></i>'+blog.catName +
                '    </span>' +
                '    <span class="item article-meta-views" data-toggle="tooltip" data-placement="bottom" title="浏览数：'+blog.vViews+'">' +
                '        <i class="glyphicon glyphicon-eye-open" id="bViews"></i>'+blog.vViews +
                '    </span>' +
                '    <span class="item article-meta-comment" data-toggle="tooltip" data-placement="bottom" title="回复数：'+blog.vReply_num+'">' +
                '       <i class="glyphicon glyphicon-comment" b="bReplys"></i>'+blog.vReply_num +
                '    </span>' +
                '</div>';
            $("#blog_header").append(blogInfo_html_str);
            //添加主体内容
            $("#bContent").append(blog.bContent);
            //文章标签
            $("#bLabel").text(blog.lName);
        }else{
            
        }
    }, "json");
}