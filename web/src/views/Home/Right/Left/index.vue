<!-- 热门动态和个性动态 -->
<template>
  <div>
    <div id="sub_nav_wrap" class="sub_nav_wrap">
      <ul class="sub_nav_list">
        <li style="margin-left:8px">
          <a id="j_ramen_nav" class="nav_li nav_li_all" href="#" >热门动态</a>
        </li>
        <li class="sub_nav_line"></li>
        <li id="nav-personal-wrapper">
          <a id="nav-personal" class="nav_li_personal cur" href="#" >个性动态</a>
        </li>
        <li class="nav_hover" style="left: 108px; width: 56px;"></li>
      </ul>
    </div>
    <div id="info_section"></div>
    <ul id="new_list" class="new_list">
      <li class="clearfix j_feed_li  1" v-for="item in  postList" :key="item.id">
        <div class="n_right">
          <div>
            <div class="title-tag-wrapper">
              <a class="n_name feed-forum-link" href="#" title="">
                {{item.family.name}}
              </a>
            </div>
            <div class="thread-name-wrapper">
              <a>
                {{item.time}}
              </a>
            </div>
            <div class="n_txt">{{item.content}}</div>
            <ul class="n_img clearfix" >
              <li>
                <img class="m_pic" :src="item.imgURL" alt="" width="150" height="105">
              </li>
            </ul>
            <div class="n_reply">
              <a class="post_author">{{item.author_name}}</a>
              <span class="time">{{ dateFormat(item.create_time)}}</span>
            </div>
          </div>
        </div>
      </li>
    </ul>
  </div>
</template>

<script>
import {dspWorksList} from "@/api";

export default {
  data() {
    return {
      postList: [],
      name:"",
      items: [
        {
          id : '1',
          title : "IT之家",
          threadName:'平台使用教程',
          txt:'欢迎使用本平台，本平台基于区块链技术开发',
          imgURL:'https://tse3-mm.cn.bing.net/th/id/OIP-C.xo4o5YN7hL-CWdX-IsSGbgHaEo?pid=ImgDet&rs=1',
          author:'admin',
          time:'20:00'
        },
      ]
    }
  },
  methods:{
    // 获取作品内容
    getPostList() {
      dspWorksList().then(res => {
        this.postList = res.data
      })
    },
    //时间格式转换
    dateFormat: function (time) {
      var date = new Date(time);
      var year = date.getFullYear();
      /* 在日期格式中，月份是从0开始的，因此要加0
       * 使用三元表达式在小于10的前面加0，以达到格式统一  如 09:11:05
       * */
      var month =
          date.getMonth() + 1 < 10 ? "0" + (date.getMonth() + 1) : date.getMonth() + 1;
      var day = date.getDate() < 10 ? "0" + date.getDate() : date.getDate();
      var hours = date.getHours() < 10 ? "0" + date.getHours() : date.getHours();
      var minutes = date.getMinutes() < 10 ? "0" + date.getMinutes() : date.getMinutes();
      var seconds = date.getSeconds() < 10 ? "0" + date.getSeconds() : date.getSeconds();
      // 拼接
      return year + "-" + month + "-" + day + " " + hours + ":" + minutes + ":" + seconds;
    },
  },
  mounted() {
    this.getPostList()
  }
}
</script>

<style scoped >

.sub_nav_wrap {
  width: 550px;
  height: 35px;
  font-family: STHeiti,"Microsoft Yahei","Microsoft YaHei",Arial,sans-serif;
  font-size: 14px;
  position: relative;
  z-index: 2;
  background: url(//tb2.bdstatic.com/tb/static-spage/img/css_078cbdc.png) no-repeat -7px -8px;
  box-shadow: 0 1px #eaeaea;
}

.sub_nav_list {
  height: 35px;
  float: left;
  postion: relative;
}

.sub_nav_list a {
  display: block;
  float: left;
  height: 33px;
  line-height: 35px;
  overflow: hidden;
  color: #444;
}

.sub_nav_wrap a {
  text-decoration: none;
}

.sub_nav_list li {
  padding: 0 14px;
  float: left;
  display: inline;
  height: 35px;
}

.sub_nav_list li.sub_nav_line {
  width: 2px;
  padding: 0;
  background: url(//tb2.bdstatic.com/tb/static-spage/img/h_sub_nav_line_1aabe40.jpg) no-repeat center;
}

.new_list {
  margin: 18px 10px 0;
}

.new_list li {
  padding-bottom: 24px;
}

.n_right {
  width: 532px;
}

.title-tag-wrapper {
  margin-bottom: 5px;
  margin-bottom: 7px\9;
}

.n_name, .n_name:hover {
  color: #444;
}

.new_list .title, .n_name {
  position: relative;
  display: inline-block;
  font-family: STHeiti,"Microsoft Yahei","Microsoft YaHei",Arial,sans-serif;
  font-size: 16px;
  color: #005097;
}

.new_list .title, .n_name {
  position: relative;
  display: inline-block;
  font-family: STHeiti,"Microsoft Yahei","Microsoft YaHei",Arial,sans-serif;
  font-size: 16px;
  color: #005097;
}

.n_txt {
  line-height: 24px;
  color: #666;
  font-size: 14px;
  padding: 4px 0 0;
}

.n_img, .n_media {
  width: 532px;
  margin: 10px 0 6px;
}

.n_img img {
  border: 1px solid #efefef;
  cursor: pointer;
}

.n_reply {
  padding-top: 2px;
}

.new_list .post_author {
  color: #999;
  padding-left: 18px;
  background: url(//tb2.bdstatic.com/tb/static-spage/img/css1_44fc180.png) no-repeat -198px -69px;
}

.new_list .time {
  color: #999;
  padding-left: 21px;
  background: url(//tb2.bdstatic.com/tb/static-spage/img/time-bg_d0a7921.png) no-repeat -10px -9px;
}

.n_reply a {
  margin-right: 8px;
}

.j_feed_li{
  height: 250px;
}
</style>
