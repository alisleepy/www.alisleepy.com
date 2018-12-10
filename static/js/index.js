/**
 * Created by Goland.
 * User: wangkaikai
 * Date: 2018/12/06
 * Time: 21:01
 */
$(function(){
    //获取推荐文章
    getTopBlogs();
    //获取文章分类列表
    getBlogsCategorys();
    //获取文章列表（默认第一页）
    getPageBlogs();
    //获取统计信息
    getSiteCounts();
    // //获取站长信息、
    getUserInfo();
    // //获取点击排行前5的文章
    // getTopViewBlogs();
    // //获取友链
    // getFriendlyUrls();
    // //获取底部备案信息
    // getFooterInfo();
})
//获取推荐文章
function getTopBlogs(){
    var url = "/home/getTopBlog";
    $.get(url,function(data){
        if(data.code == 200){
            var topBlogObj = $("#topBlogs");
            var datas = data.data;
            for(i in datas){
                var blog_html_str = '';
                blog_html_str +=
                    '<article class="excerpt-minic excerpt-minic-index">'+
                    '   <h2>'+
                    '       <span class="red">【推荐】</span>'+
                    '       <a href="/home/getBlogInfo?bId='+datas[i].bId+'" title="'+datas[i].bTitle+'" >'+
                                datas[i].bTitle+
                    '       </a>'+
                    '   </h2>'+
                    '   <p class="note"><a href="/home/getBlogInfo?bId='+datas[i].bId+'">'+datas[i].bInfo+'</a></p>'+
                    '</article>';
            }
            topBlogObj.append(blog_html_str);
        }else{
            topBlogObj.append('<span>暂无推荐内容</span>');
        }
    },'json');
}
//获取文章分类列表
function getBlogsCategorys(){
    var url = "/home/getCategorys";
    $.get(url, function(data){
        var categoryObj = $("#categorys");
        if(data.code == 200){
            var datas = data.data;
            var category_html_str = '';
            for(i in datas){
                console.log(i);
                console.log(datas[i].catName);
                category_html_str += '<a href="/home/getBlogsList?catId='+datas[i].catId+'">'+datas[i].catName+'</a>';
            }
            categoryObj.append(category_html_str);
        }else{
            categoryObj.append('<span>暂无文章分类</span>');
        }
    }, "json");
}
//获取首页文章列表
function getPageBlogs(){
    var curpage = $("#curpage").val();
    var catId = $("#catId").val();
    var lId = $("#lId").val();
    var url = "/home/ajaxGetBlogs";
    var param = {curpage:curpage, catId:catId, lId:lId};
    $.get(url, param, function(data){
        console.log(data);
        var blogObj = $("#blogs");
        if(data.code == 200){
            var datas = data.data;
            var blogs_html_str = '';
            for(i in datas){
                console.log(i);
                var add_time = getLocalTime(datas[i].add_time);
                blogs_html_str +=
                    '<article class="excerpt excerpt-1" style="">' +
                    '   <a class="focus" href="#" title="'+datas[i].bTitle+'">' +
                    '       <img class="thumb" src="/static/images/201610181739277776.jpg" alt="'+datas[i].bTitle+'"  style="display: inline;">' +
                    '   </a>' +
                    '   <header>' +
                    '       <a class="cat" href="#" title="'+datas[i].catName+'" >' +
                    '           '+datas[i].catName+'<i></i>' +
                    '       </a>' +
                    '       <h2>' +
                    '           <a href="#" title="'+datas[i].bTitle+'">'+datas[i].bTitle+'</a>' +
                    '       </h2>' +
                    '   </header>' +
                    '   <p class="meta">' +
                    '       <time class="time"><i class="glyphicon glyphicon-time"></i>'+add_time+'</time>' +
                    '       <span class="views"><i class="glyphicon glyphicon-eye-open"></i>'+datas[i].vViews+'</span>' +
                    '       <a class="comment" href="##comment" title="评论"><i class="glyphicon glyphicon-comment"></i>'+datas[i].vReply_num+'</a>' +
                    '   </p>' +
                    '   <p class="note">'+datas[i].bInfo+'</p>' +
                    '</article>';
            }
            blogObj.append(blogs_html_str);
        }else{
            blogObj.append('<span>我可是有底线的！！</span>');
        }
    }, "json");
}
//js把时间戳转化为日期格式
function getLocalTime(nS) {
    return new Date(parseInt(nS) * 1000).toLocaleString().replace(/:\d{1,2}$/,' ');
}
//获取统计信息
function getSiteCounts(){
    var url = "/home/ajaxGetBlogNum";
    $.get(url, function(json){
        if(json.code == 200){
            var blogNum = json.data.BlogNum;
            $("#blogNum").html(blogNum);
        }else{
            console.log('未获取到博客总数')
        }
    },"json");
}
//获取个人信息
function getUserInfo(){
    var url = "/home/GetMyInfo";
    $.get(url, function(json){
        console.log(json);
        if(json.code == 200){

        }else{
            console.log("未获取到个人信息");
        }
    },"json");
}