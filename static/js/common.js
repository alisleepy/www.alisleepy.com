/**
 * Created by Goland.
 * User: wangkaikai
 * Date: 2018/12/14
 * Time: 22:01
 */
$(function(){
    getTopBlogs();
    //获取统计信息
    getSiteCounts();
    //获取站长信息、
    getUserInfo();
    //获取点击排行前5的文章
    getTopViewBlogs();
    //获取友链
    getFriendlyUrls();
})
//获取推荐文章
function getTopBlogs(){
    var url = "/home/getTopBlog";
    $.get(url,function(data){
        var topBlogObj = $("#topBlogs");
        if(data.code == 200){
            var datas = data.data;
            for(i in datas){
                var blog_html_str = '';
                blog_html_str +=
                    '<article class="excerpt-minic excerpt-minic-index">'+
                    '   <h2>'+
                    '       <span class="red">【推荐】</span>'+
                    '       <a href="/home/blogInfo?bId='+datas[i].bId+'" title="'+datas[i].bTitle+'" >'+
                    datas[i].bTitle+
                    '       </a>'+
                    '   </h2>'+
                    '   <p class="note"><a href="/home/blogInfo?bId='+datas[i].bId+'">'+datas[i].bInfo+'</a></p>'+
                    '</article>';
            }
            topBlogObj.append(blog_html_str);
        }else{
            topBlogObj.append('<span>暂无推荐内容</span>');
        }
    },'json');
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
    var url = "/home/getMyInfo";
    $.get(url, function(json){
        console.log(json);
        if(json.code == 200){
            var data = json.data;
            for(i in data){
                switch(data[i].key){
                    case 'qq':
                        $("#user_qq").text(data[i].value);
                        break;
                    case 'email':
                        $("#user_email").text(data[i].value);
                        $("#user_email").parent().attr("href", 'mailto:'+data[i].value);
                        break;
                    case 'viewNum':
                        $("#viewNum").text(data[i].value);
                        break;
                }
            }
        }else{
            console.log("未获取到个人信息");
        }
    },"json");
}
//获取点击前5的文章
function getTopViewBlogs(){
    var url = "/home/getTopViewBlog";
    $.get(url, function(json){
        if(json.code == 200){
            var datas = json.data;
            var topViewBlogs_html_str = '';
            for(i in datas){
                topViewBlogs_html_str +=
                    '<li>' +
                    '   <a title="" href="/home/blogInfo?bId='+datas[i].bId+'" >' +
                    '       <span class="thumbnail">' +
                    '           <img class="thumb" src="/static/images/201610181739277776.jpg" alt="'+datas[i].bTitle+'"  style="display: block;">' +
                    '       </span>' +
                    '       <span class="text">'+datas[i].bTitle+'</span>' +
                    '       <span class="muted"><i class="glyphicon glyphicon-time"></i>'+getLocalTime(datas[i].add_time)+'</span>' +
                    '       <span class="muted"><i class="glyphicon glyphicon-eye-open"></i>'+datas[i].vViews+'</span>' +
                    '   </a>' +
                    '</li>';
            }
            $("#topViewBlos").append(topViewBlogs_html_str);
        }else{
            $("#topViewBlos").append('<span>暂无文章</span>');
        }
    }, "json");
}
//获取友联
function getFriendlyUrls(){
    var url = "/home/getFriendluUrl";
    $.get(url, function(urls){
        //console.log(urls);
        var jsonValue = urls.data.value;
        if(jsonValue.length > 0){
            var dataArr = JSON.parse(jsonValue);
            var friendlyurl_html_str = '';
            for(i in dataArr){
                friendlyurl_html_str +=
                    '<div class="widget-sentence-link">' +
                    '   <a href="'+dataArr[i].url+'" title="'+dataArr[i].name+'" target="_blank" >'+dataArr[i].name+'</a>' +
                    '</div>';
            }
            $("#friendlyUrl").append(friendlyurl_html_str)
        }
    },"json");
}
//js把时间戳转化为日期格式
function getLocalTime(nS) {
    return new Date(parseInt(nS) * 1000).toLocaleString().replace(/:\d{1,2}$/,' ');
}
//回到顶部
function gotop(){
    $("html,body").animate({scrollTop:0},200);
}