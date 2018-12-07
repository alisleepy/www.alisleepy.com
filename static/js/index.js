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
    // getPageBlogs();
    // //获取统计信息
    // getSiteCounts();
    // //获取站长信息、
    // getUserInfo();
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
                topBlogObj.append(blog_html_str);
            }
        }else{
            topBlogObj.append('<span>暂无推荐内容</span>');
        }
    },'json');
}
//获取文章分类列表
function getBlogsCategorys(){
    var url = "/home/getCategorys";
    $.get(url, function(data){
        if(data.code == 200){
            var categoryObj = $("#categorys");
            var datas = data.data;
            var category_html_str = '';
            for(i in datas){
                console.log(i);
                console.log(datas[i].catName);
                category_html_str += '<a href="/home/getBlogsList?catId='+datas[i].catId+'">'+datas[i].catName+'</a>';
            }
            categoryObj.append(category_html_str);
        }else{
            categoryObj.append('暂无文章分类');
        }
    }, "json");
}