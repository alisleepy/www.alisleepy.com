$(function(){
    var bId = $("#bId").val();
    console.log(bId);
    getTicketInfo(bId);   //获取文章详情
    getTicketReplys(bId); //获取文章评论列表
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
            $("#bContent-div").text(blog.bContent);
            //文章标签
            $("#bLabel").text(blog.lName);
            showPostComments(blog.allowReply);
        }else{
            $("#blog_header").text("无内容！！!");
            $("#bLabel").remove();
        }
    }, "json");
}
/*
 * 显示发表评论区域
 * @param allowReply int  是否允许评论 0/不允许，1/允许评论
 */
function showPostComments(allowReply){
    console.log(allowReply);
    if(allowReply == 1){
        $("#postcomments").show();
    }
}
/*
 * 获取文章评论列表
 * @param bId  文章id
 */
function getTicketReplys(bId){
    if(bId){
        var url = '/home/getBlogReplys';
        $.get(url, {bId:bId}, function(json){
            if(json.code == 200){
                var data = json.data
                var reply_html_str = '';
                for(i in data){
                    var grade = i + 1;
                    reply_html_str += '<li class="comment-content"><span class="comment-f">#'+i+'</span>' +
                        '                  <div class="comment-main">' +
                        '                       <p>' + data[i].uName +
                        '                           <span class="time">(2016/10/2811:41:03)</span><br>' + delHtmlTag(data[i].rContent) +
                        '                       </p>' +
                        '                  </div>' +
                        '               </li>';
                }
                $("#comment_list").append(reply_html_str);
            }else{
                $("#comment_list").append('<span>该文章暂无评论</span>');
            }
        },"json");
    }
}
//js去掉html标签
var delHtmlTag = function (msg) {
    var msg = msg.replace(/<\/?[^>]*>/g, ''); //去除HTML Tag
    msg = msg.replace(/[|]*\n/, '') //去除行尾空格
    msg = msg.replace(/&npsp;/ig, ''); //去掉npsp
    return msg;
}
// $(function (){
//     testEditor = editormd.markdownToHTML("bContent-div", {//注意：这里是上面DIV的id
//         htmlDecode: "style,script,iframe",
//         emoji: true,
//         taskList: true,
//         tex: true, // 默认不解析
//         flowChart: true, // 默认不解析
//         sequenceDiagram: true, // 默认不解析
//         codeFold: true,
//     });
// });