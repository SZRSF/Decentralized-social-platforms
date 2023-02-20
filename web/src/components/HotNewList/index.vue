<template>
  <div>
    <ul class="new_list clearfix">
      <li class="new_list_item"  v-for="n in hotNewList" :key="n.n_id">
        <div class="n_right n_right_first clearfix">
          <div class="n_post_time">{{n.n_PostTime}}</div>
          <div class="n_avatar">
            <a href="#">
              <img :src="n.n_logo" alt="">
            </a>
          </div>
          <div class="n_contain">
            <div>
              <div class="title_wraper">
                <a  class="n_name" href="#" >{{n.n_name}}</a>
              </div>
              <div class="thread_name">
                <a class="title" href="#">{{n.n_title}}</a>
              </div>
            </div>
            <div class="n_txt">{{n.n_txt}}</div>
            <ul class="n_media clearfix">
              <li v-for="img in n.n_img" :key="img.id">
                <img :src="img.url" alt="">
              </li>
            </ul>
          </div>
        </div>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  name: "HotNewList",
  data () {
    return {
      hotNewList: [
        {
          n_id: '1',
          n_PostTime: '20:00',
          n_logo: 'http://imgsrc.baidu.com/forum/pic/item/bd3eb13533fa828b52da1e31fb1f4134960a5a63.jpg',
          n_name: '笔记本之家',
          n_title: '笔记本避坑指南',
          n_txt:'intel、Nvidia、AMD这些上游芯片厂商在1月初公布了2023年的产品.....',
          n_img:[
            {
              id: '1',
              url: 'http://imgsrc.baidu.com/forum/pic/item/bd3eb13533fa828b52da1e31fb1f4134960a5a63.jpg'
            },
            {
              id: '2',
              url: 'http://imgsrc.baidu.com/forum/pic/item/bd3eb13533fa828b52da1e31fb1f4134960a5a63.jpg'
            },
            {
              id: '3',
              url: 'http://imgsrc.baidu.com/forum/pic/item/bd3eb13533fa828b52da1e31fb1f4134960a5a63.jpg'
            }
          ],
          n_author: "逝者如斯夫",
          n_reply: "20",
        },
        {
          n_id: '2',
          n_PostTime: '20:00',
          n_logo: 'http://imgsrc.baidu.com/forum/pic/item/bd3eb13533fa828b52da1e31fb1f4134960a5a63.jpg',
          n_name: '笔记本之家',
          n_title: '笔记本避坑指南',
          n_txt:'intel、Nvidia、AMD这些上游芯片厂商在1月初公布了2023年的产品.....',
          n_img:[
            {
              id: '1',
              url: 'http://imgsrc.baidu.com/forum/pic/item/bd3eb13533fa828b52da1e31fb1f4134960a5a63.jpg'
            },
            {
              id: '2',
              url: 'http://imgsrc.baidu.com/forum/pic/item/bd3eb13533fa828b52da1e31fb1f4134960a5a63.jpg'
            },
            {
              id: '3',
              url: 'http://imgsrc.baidu.com/forum/pic/item/bd3eb13533fa828b52da1e31fb1f4134960a5a63.jpg'
            }
          ],
          n_author: "逝者如斯夫",
          n_reply: "20",
        },
      ]
    }
  },
  methods: {

  }
}
</script>

<style scoped lang="less">
.new_list {
  margin-left: 5px;
}

.n_right_first {
  border-top: 0;
}

.new_list_item{
  height:300px;
}

.n_right {
  width: 690px;
  border-top: 1px dotted #E1E3E6;
  position: relative;
  padding-top: 15px;
  margin-bottom: 12px;
}

.new_list .n_post_time {
  position: absolute;
  font-family: Arial SimSun,serif;
  font-size: 12px;
  line-height: 12px;
  color: #919499;
  top: 16px;
  height: 20px;
  right: 0;
}

.n_avatar {
  float: left;
  margin: 2px 10px 0 0;
  position: relative;
}

.n_avatar a {
  display: block;
  border: 1px solid #ebebeb;
  width: 43px;
  height: 43px;
}

.n_avatar img {
  width: 43px;
  height: 43px;
}

.n_contain {
  float: left;
  width: 630px;
  position: relative;
}

.title_wraper {
  margin-bottom: 5px;
}

.n_name {
  font-size: 12px;
  color: #666;
}

.thread_name {
  position: relative;
}

.new_list .title, .new_list .title:hover {
  font-size: 14px;
  color: #2962B6;
}

.n_txt {
  line-height: 12px;
  color: #919499;
  font-size: 12px;
  margin: 7px 0 10px;
}

.n_media {
  width: 532px;
}
</style>
