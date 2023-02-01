<!-- 热门家推荐-->
<template>
  <div class="home_list">
    <ul id="example-1">
      <li class="home_item" v-for="item in items" :key="item.id">
        <a class="home_link" :href="item.homeUrl" target="_blank">
          <el-image
              class="home_logo"
              :src="item.url"
              fit="fill"></el-image>
          <div class="home_icon">
            <div class="home_name">{{item.homeName}}</div>
          </div>
        </a>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  name: "HomeList",
  data() {
    return {
      items:[
        {
          id: '1',
          homeName: '孙家',
          homeUrl:'',
          url:'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg'
        },
        {
          id: '2',
          homeName: '李家',
          homeUrl:'',
          url:'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg'
        },
        {
          id: '3',
          homeName: '李家',
          homeUrl:'',
          url:'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg'
        },
        {
          id: '4',
          homeName: '李家',
          homeUrl:'',
          url:'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg'
        },
        {
          id: '5',
          homeName: '李家',
          homeUrl:'',
          url:'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg'
        },
        {
          id: '6',
          homeName: '李家',
          homeUrl:'',
          url:'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg'
        },
        {
          id: '7',
          homeName: '李家',
          homeUrl:'',
          url:'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg'
        },
        {
          id: '8',
          homeName: '李家',
          homeUrl:'',
          url:'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg'
        },
        {
          id: '9',
          homeName: '李家',
          homeUrl:'',
          url:'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg'
        },
        {
          id: '10',
          homeName: '李家',
          homeUrl:'',
          url:'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg'
        },
        {
          id: '11',
          homeName: '李家',
          homeUrl:'',
          url:'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg'
        },
        {
          id: '12',
          homeName: '李家',
          homeUrl:'',
          url:'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg'
        },
      ]
    }
  }
}
</script>

<style scoped>
.home_item {
  width: 180px;
  height: 60px;
  margin-bottom: 16px;
  float: left;
}

.home_name {
  width: 102px;
  font-size: 15px;
  color: #000;
  font-weight: 400;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
.home_link {
  display: -ms-flexbox;
  display: flex;
  width: 100%;
  height: 100%;
}

.home_logo {
  width: 60px;
  height: 60px;
  border: 1px solid rgba(0,0,0,.1);
  border-radius: 6px;
  background: #f7f7f7;
}

.home_icon {
  margin-left: 10px;
  display: -ms-flexbox;
  display: flex;
  -ms-flex-direction: column;
  flex-direction: column;
  -ms-flex-pack: distribute;
  justify-content: space-around;
}

</style>
