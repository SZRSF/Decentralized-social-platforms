<template>
  <div id="userinfo_wrap" class="userinfo_wrap clearfix">
    <div class="userinfo_left">
      <div  class="userinfo_left_head" title="点击查看高清头像">
        <a class="userinfo_head" >
          <el-image
              style="width: 100%; height: 100%"
              :src="url"
              :preview-src-list="srcList">
          </el-image>
        </a>
      </div>
    </div>
    <div class="userinfo_middle">
      <div class="userinfo_relation">
        <div class="btn_grounps">
          <a class="btn-sub btn-small">
            编辑资料
          </a>
        </div>
      </div>
      <div class="userinfo_title">
        <span class="userinfo_username ">{{userinfo_userdata.username}}</span>
      </div>
      <div class="userinfo_num">
        <div class="userinfo_userdata">
          <span class="user_name">用户名:{{userinfo_userdata.username}}</span>
          <span class="userinfo_split"></span>
          <span>年龄:{{userinfo_userdata.age}}</span>
          <span class="userinfo_split"></span>
          <span>发帖:{{userinfo_userdata.worksNums}}</span>
          <span class="userinfo_split"></span>
          <span>IP属地:{{userinfo_userdata.IP}}</span>
        </div>
      </div>
      <div class="userinfo_honor clearfix">

      </div>
    </div>
    <div class="userinfo_right">
      <div class="userinfo_scores clearfix">
        <span class="scores-header-hr"></span>
        <span>我的D币</span>
        <a class="userinfo_scores_num j_userinfo_scores_num" href="#">
          <i class="el-icon-coin"></i>
          {{userinfo_userdata.money}}
        </a>
        <div>
          <el-button size="mini">获取D币</el-button>
          <el-button size="mini">社区商城</el-button>
        </div>
      </div>
      <div class="userinfo_shortcut">
        <a href="#">服务中心</a>
        <span class="userinfo_split"></span>
        <a href="#">我的家</a>
      </div>
    </div>
  </div>
</template>

<script>
// import store from "@/store/user";

export default {
  name: "UserInfo",
  data(){
    return {
      userinfo_userdata: {
        username: this.$store.getters.donUsers.username,
        age: '18',
        worksNums: '1',
        IP: '广东',
        money:'100'
      },
      url: 'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg',
      srcList: [
        'https://fuss10.elemecdn.com/8/27/f01c15bb73e1ef3793e64e6b7bbccjpeg.jpeg',
        'https://fuss10.elemecdn.com/1/8e/aeffeb4de74e2fde4bd74fc7b4486jpeg.jpeg'
      ]
    }
  }
}
</script>

<style scoped lang="less">
.userinfo_wrap {
  min-height: 110px;
  width: 978px;
  margin: 0 auto;
  border-style: solid;
  border-color: #8FC0D3;
  border-width: 0 1px;
  background: #F5F7FA;
  position: relative;
}

.userinfo_left_head {
  left: 17px;
  position: absolute;
  top: -65px;
}

.userinfo_middle {
  width: 500px;
  float: left;
}

.userinfo_head {
  display: block;
  background: url(//tb2.bdstatic.com/tb/static-ihome/img/userinfo_head_bg_bbcc388.png) no-repeat;
  height: 150px;
  padding: 5px;
  width: 150px;
  cursor: url(https://gsp0.baidu.com/5aAHeD3nKhI2p27j8IqW0jdnxx1xbK/tb/img/frs/cur_zin.cur?t=1667993245910),pointer;
}

.userinfo_left {
  position: relative;
  width: 190px;
  height: 110px;
  float: left;
}

.userinfo_relation {
  position: relative;
  left: 0;
  top: 0;
  zoom: 1;
}

.btn_grounps {
  position: absolute;
  height: 26px;
  left: 0;
  bottom: 11px;
  width: 780px;
}

.userinfo_title {
  font-family: 'Microsoft YaHei';
  font-size: 18px;
  height: 26px;
  line-height: 26px;
  margin-top: 10px;
  overflow: hidden;
}

.userinfo_username {
  float: left;
}

.userinfo_num {
  margin-top: 4px;
  color: #797C80;
  height: 22px;
  zoom: 1;
}

.userinfo_userdata {
  float: left;
}

.userinfo_split {
  border-color: #CCC;
  border-style: solid;
  border-width: 0 0 0 1px;
  display: inline-block;
  height: 10px;
  line-height: 10px;
  margin: 0 10px 2px 9px;
  vertical-align: middle;
}

.userinfo_num {
  margin-top: 4px;
  color: #797C80;
  height: 22px;
  zoom: 1;
}

.userinfo_honor {
  margin-bottom: 9px;
  position: relative;
}

.userinfo_right {
  position: absolute;
  right: 0;
  height: 100px;
  font-family: 'Microsoft YaHei',serif;
}

.userinfo_scores {
  float: left;
  *display: inline;
  text-align: center;
  color: #797C80;
  margin: 10px;
}

.userinfo_scores span, .userinfo_scores a {
  display: inline-block;
  vertical-align: middle;
}

.userinfo_scores_num {
  color: #FF7F3E;
  height: 28px;
  line-height: 28px;
}

.userinfo_shortcut {
  position: absolute;
  bottom: 0;
  right: 20px;
  width: 206px;
  text-align: right;
}
</style>
