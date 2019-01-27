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
            //$("#bContent-div").text(blog.bContent);
            $("#bContent").text(blog.bContent);
            showArticleContent();
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
//把markdown格式的内容转化为html显示
function showArticleContent(){
    testEditor = editormd.markdownToHTML("bContent-div", {//注意：这里是上面DIV的id
        htmlDecode: "style,script,iframe",
        emoji: true,
        taskList: true,
        tex: true, // 默认不解析
        flowChart: true, // 默认不解析
        sequenceDiagram: true, // 默认不解析
        codeFold: true,
    });
}
//根据关键词搜索
$('button[name="search"]').click(function(){
    var keywords = $('input[name="keywords"]').val();
    $("#blogs").empty();
    removeInfoAndAddDiv();
    getPageBlogs(1, 0, 0, keywords);
});
//获取文章内容
function getPageBlogs(curpage, catId, lId, keywords){
    var url = "/home/ajaxGetBlogs";
    var param = {curpage:curpage, catId:catId, lId:lId, keywords:keywords};
    $.get(url, param, function(data){
        resetGetmoreblog_a();
        var blogObj = $("#blogs");
        if(data.code == 200){
            var datas = data.data;
            console.log(datas);
            var blogs_html_str = '';
            for(i in datas){
                //console.log(i);
                var add_time = getLocalTime(datas[i].add_time);
                blogs_html_str +=
                    '<article class="excerpt excerpt-1" style="">' +
                    '   <a class="focus" href="/home/blogInfo?bId='+datas[i].bId+'" title="'+datas[i].bTitle+'">' +
                    '       <img class="thumb" src="/static/images/201610181739277776.jpg" alt="'+datas[i].bTitle+'"  style="display: inline;">' +
                    '   </a>' +
                    '   <header style="padding-top:12px;">' +
                    '       <a class="cat" href="'+datas[i].bId+'" title="'+datas[i].catName+'" >' +
                    '           '+datas[i].catName+'<i></i>' +
                    '       </a>' +
                    '       <h2>' +
                    '           <a href="/home/blogInfo?bId='+datas[i].bId+'" title="'+datas[i].bTitle+'">'+datas[i].bTitle+'</a>' +
                    '       </h2>' +
                    '   </header>' +
                    '   <p class="meta">' +
                    '       <time class="time"><i class="glyphicon glyphicon-time"></i>'+add_time+'</time>' +
                    '       <span class="views"><i class="glyphicon glyphicon-eye-open"></i>'+datas[i].vViews+'</span>' +
                    '       <a class="comment" href="/home/blogInfo?bId='+datas[i].bId+'" title="评论"><i class="glyphicon glyphicon-comment"></i>'+datas[i].vReply_num+'</a>' +
                    '   </p>' +
                    '   <p class="note"><a href="/home/blogInfo?bId='+datas[i].bId+'">'+datas[i].bInfo+'</a></p>' +
                    '</article>';
            }
            blogObj.append(blogs_html_str);
        }else{
            removeGetmoreblog_a(); //修改a标签
        }
        //更新当前页码和标签id和分类id
        setPageAndCatIDAndlId(data.page, data.catId, data.lId, data.keywords)
    }, "json");
}
function removeInfoAndAddDiv(){
    //移除现在的元素
    var obj = $('div[class="content"]');
    obj.empty();
    //新增加元素
    var divstr = '<div id="blogs"></div>' +
        '<nav class="pagination" style="display: block;">' +
        '<ul>' +
        '<li><a id="getmoreblog_a" href="javascript:void(0);" onclick="getMoreBlogs()">点击加载更多</a></li>' +
        '</ul>' +
        '</nav>';
    obj.append(divstr);
}
//去掉底部a标签
function removeGetmoreblog_a(){
    $("#getmoreblog_a").text("我可是有底线的！！!");
    $("#getmoreblog_a").css("color", "red");
    $("#getmoreblog_a").removeAttr("onclick");
}
//复原底部a标签
function resetGetmoreblog_a(){
    $("#getmoreblog_a").text("点击加载更多");
    $("#getmoreblog_a").css("color", "#666");
    $("#getmoreblog_a").attr("onclick","getMoreBlogs()");
}
//更新当前页码和标签id和分类id
function setPageAndCatIDAndlId(page, catId, lId, keywords){
    $("#catId").val(catId);
    $("#lId").val(lId);
    $("#keywords").val(keywords);
    var nextPage = parseInt(page) + 1;
    $("#curpage").val(nextPage);
    console.log("当前页码是："+page);
    console.log("下一页页码是："+nextPage);
}
//加载更多
function getMoreBlogs(){
    var nextPage = $("#curpage").val();
    var catId = $("#catId").val();
    var lId = $("#lId").val();
    var keywords = $('input[name="keywords"]').val();
    getPageBlogs(nextPage, catId, lId, keywords);
}
//发表评论
$('#comment-submit').click(function(){
    var url = '/home/postReply';
    var bId = $('#bId').val();
    var uName = $('input[name="uName"]').val();
    var uEmail = $('input[name="uEmail"]').val();
    var rContent = $('textarea[name="rContent"]').val();;
    if(uName == undefined || uName.length == 0){
        alert('请填写下您的称呼吧！');
    }
    if(uName.length > 30){
        alert('用户名长度不对啊！');
    }
    if(rContent == undefined || rContent.length == 0){
        alert('请输入点评论内容吧！');
    }
    if(rContent.length > 200){
        alert('评论内容不能超过200个字哦！');
    }
    var data = {
        bId : bId,
        uName : uName,
        uEmail : uEmail,
        rContent : rContent
    };
    $.post(
        url, data,
        function(json){
            if(json.code == 0){
                alert('感谢您的评论');
                window,location.reload();
            }else{
                alert(json.msg);
            }
        },"json"
    );
});