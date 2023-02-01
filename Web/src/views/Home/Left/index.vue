<template>
  <!--  首页账户显式-->
  <div class="left-sec">
    <!--    头部-->
    <div id="my_home_mod" class="region_bright my_home_mod ">
      <h4 class="my_home_header">
        我的账户
      </h4>
      <!--    头像以及信息显式-->
      <div class="region_cnt clearfix">
        <!--头像 -->
        <el-row>
          <el-col :span="12">
            <router-link  to="/user">
              <el-avatar shape="square" :size="50" :src="squareUrl" ></el-avatar>
            </router-link>
          </el-col>
          <el-col :span="12">
                <router-link id="userValue"  class=""  to="/user">{{username}}</router-link>
          </el-col>
        </el-row>
        </div>
      </div>
<!--    资料页面展示-->
    <div>
      <div>爱逛的家</div>
      <div>
        <ul id="example-1">
          <li v-for="item in items" :key="item.message">
            {{ item.message }}
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>


<script>

export default {
  nameL:"Left",
  data() {
    return {
      username:"逝者如斯夫",
      squareUrl: "https://gss0.bdstatic.com/6LZ1dD3d1sgCo2Kml5_Y_D3/sys/portrait/item/tb.1.6c7e99dc.naI5KHZoNvYT7ciplr4dxw?t=1668136071",
      sizeList: ["medium"],
      items: [
        { message: 'Foo' },
        { message: 'Bar' }
      ],
    }
  },
  methods: {
    toUser() {
      this.$router.push("/user")
    }
  }
}
</script>

<style scoped>

.left-sec {
  float: left;
  width: 210px;
  border-top: 1px solid #DCDCDC;
}

.page-container .left-sec .region_bright {
  padding: 16px 12px;
}

.region_bright:first-child {
  border: 0;
}

.region_bright {
  position: relative;
  padding: 16px 19px;
  border-top: 1px solid #e4e6eb;
}

.my_home_mod .media_left, .my_home_mod .media-left {
  width: 80px;
  height: 80px;
  padding: 2px;
  border: 1px solid #ccc;
}
.media_right, .media-right {
  position: relative;
}
.media_left img, .media-left img {
  width: 100%;
  height: 100%;
}

.region_bright .region_cnt {
  margin-top: 10px;
}
.text_overflow, .text-overflow {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.my_home_mod .media_right, .my_home_mod .media-right {
  margin-left: 96px;
}

</style>
