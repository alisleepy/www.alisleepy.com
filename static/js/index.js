/**
 * Created by Goland.
 * User: wangkaikai
 * Date: 2018/12/06
 * Time: 21:01
 */
$(function(){
    var curpage = $("#curpage").val();
    var catId = $("#catId").val();
    var lId = $("#lId").val();
    var keywords = $("#keywords").val();
    //获取文章分类列表
    getBlogsCategorys();
    //获取文章列表（默认第一页）
    getPageBlogs(curpage, catId, lId, keywords);
})

//获取文章分类列表
function getBlogsCategorys(){
    var url = "/home/getCategorys";
    $.get(url, function(data){
        var categoryObj = $("#categorys");
        if(data.code == 200){
            var datas = data.data;
            var category_html_str = '';
            for(i in datas){
                category_html_str += '<a href="javascript:void(0);" onclick=getCatBlogs('+datas[i].catId+');>'+datas[i].catName+'</a>';
            }
            categoryObj.append(category_html_str);
        }else{
            categoryObj.append('<span>暂无文章分类</span>');
        }
    }, "json");
}
//获取首页文章列表
function getPageBlogs(curpage, catId, lId, keywords){
    var url = "/home/ajaxGetBlogs";
    var param = {curpage:curpage, catId:catId, lId:lId, keywords:keywords};
    $.get(url, param, function(data){
        //console.log(data);
        resetGetmoreblog_a();
        var blogObj = $("#blogs");
        if(data.code == 200){
            var datas = data.data;
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
//加载更多文章
function getMoreBlogs(){
    var nextPage = $("#curpage").val();
    var catId = $("#catId").val();
    var lId = $("#lId").val();
    var keywords = $("#keywords").val();
    getPageBlogs(nextPage, catId, lId, keywords);
}
//获取某个分类下的文章
function getCatBlogs(catId){
    //先把原来的数据清掉
    $("#blogs").empty();
    getPageBlogs(1, catId, 0, "");
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