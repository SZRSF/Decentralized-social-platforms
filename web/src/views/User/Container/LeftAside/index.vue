<template>
  <div class="left_aside ihome_left_aside">
    <div class="content_wrap">
      <div id="ihome_nav_wrap" class="ihome_nav_wrap">
        <ul class="ihome_nav_list">
          <li class="focus">
            <div class="tbnav_tab_inner">
              <p class="space">
                <el-link @click="$router.push('/user')" target="_blank">
                  <i class="el-icon-house"></i>
                  我的主页
                </el-link>
              </p>
            </div>
          </li>
          <li>
            <div class="tbnav_tab_inner">
              <p class="space">
                <el-link @click="$router.push('#')" target="_blank">
                  <i class="el-icon-bell"></i>
                  通知
                </el-link>
              </p>
            </div>
          </li>
          <li>
            <div class="tbnav_tab_inner">
              <p class="space">
                <el-link @click="$router.push('#')" target="_blank">
                  <i class="el-icon-star-off"></i>
                  收藏
                </el-link>
              </p>
            </div>
          </li>
          <li>
            <div class="tbnav_tab_inner">
              <p class="space">
                <el-link @click="$router.push('#')" target="_blank">
                  <i class="el-icon-s-flag"></i>
                  我的家庭
                </el-link>
              </p>
            </div>
          </li>
          <li>
            <div class="tbnav_tab_inner">
              <p class="space">
                <el-link @click="$router.push('/works')" target="_blank">
                  <i class="el-icon-edit"></i>
                  作品
                </el-link>
              </p>
            </div>
          </li>
          <li>
            <div class="tbnav_tab_inner">
              <p class="space">
                <el-link @click="$router.push('#')" target="_blank">
                  <i class="el-icon-trophy"></i>
                  成就
                </el-link>
              </p>
            </div>
          </li>
        </ul>
      </div>
      <div class="ihome_forum_group ihome_section clearfix">
        <div id="onekey_sign" >
          <a class="onekey_btn">
            <i class="icon_signed"></i>
          </a>
        </div>
        <h1 class="ihome_title">
          <span class="cut_line">|</span>
          常串的家
        </h1>
        <div id="forum_group_wrap" class="clearfix u-f-wrap" >
          <ul>
            <li v-for="family in familyList" :key="family.id">
              <a class="u-f-item unsign">
                <span>{{family.name}}</span>
              </a>
            </li>
          </ul>
<!--          <span class="more u-f-item j_show_more_forum">-->
<!--          <el-collapse class="more u-f-item j_show_more_forum" v-model="activeNames" @change="handleChange">-->
<!--              <el-collapse-item  title="展开" name="1">-->
<!--                <div>-->
<!--                  <a class="u-f-item unsign">-->
<!--            <span>{{family.name}}</span>-->
<!--          </a>-->
<!--                  <a class="u-f-item unsign">-->
<!--            <span>{{family.name}</span>-->
<!--          </a>-->
<!--                </div>-->
<!--              </el-collapse-item>-->
<!--           </el-collapse>-->
<!--            <span class="triangle-down"></span>-->
<!--          </span>-->
        </div>
      </div>
      <div class="ihome_hot_feed ihome_section">
        <h1 class="threadList_title">
          <span class="cut_line">|</span>
          热门动态
        </h1>
        <HotNewList></HotNewList>
      </div>
    </div>
  </div>
</template>

<script>
import HotNewList from "@/components/HotNewList";
import {dspFamilyList} from "@/api";

export default {
  data() {
    return{
      activeNames: ['1'],
      familyList: []
    }
  },
  methods: {
    handleChange(val) {
      console.log(val);
    },
    // 获取家的列表
    getFamilyList() {
      dspFamilyList().then(res => {
        this.familyList = res.data;
      })
    },
  },
  components: {
    HotNewList,
  },
  mounted(){
    this.getFamilyList();
  },
}
</script>

<style  lang="less">
.ihome_left_aside {
  padding-bottom: 20px;
}

.left_aside {
  width: 100%;
  float: left;
  background: url(//tb2.bdstatic.com/tb/static-ihome/img/content_bg_30a9c8a.jpg) repeat-y;
}

.content_wrap {
  width: 740px;
}

.ihome_nav_wrap {
  height: 46px;
  background-repeat: repeat-x;
}

.ihome_nav_list {
  height: 46px;
  font-family: STHeiti,"Microsoft Yahei",Arial,SimSun,serif;
  font-size: 16px;
  float: left;
  overflow: hidden;
}

.ihome_nav_list .focus {
  background-repeat: repeat-x;
  background-position: 0 -100px;
}

.tbnav_tab_inner {
  background-repeat: no-repeat;
  background-position: left -250px;
}

.space {
  background-position: right -250px;
}

.ihome_nav_list a.nav_icon {
  background-position-x: 0%;
  background-position-y: 0%;
  padding-left: 43px;
}

.ihome_nav_list a.nav_main {
  background-position: 22px -339px;
}

.ihome_nav_list a {
  display: block;
  padding: 0 20px;
  position: relative;
}

li {
  height: 46px;
  line-height: 46px;
  float: left;
  display: inline;
  margin-left: -2px;
}

.ihome_left_aside .ihome_forum_group {
  margin-top: 18px;
}

.ihome_section {
  margin: 20px 20px 0;
}

.ihome_forum_group #onekey_sign {
  width: 72px;
  float: right;
}
.ihome_left_aside .ihome_forum_group {
  margin-top: 18px;
}

.ihome_forum_group{
  height: 136px;
}

.ihome_section {
  margin: 20px 20px 0;
}

#onekey_sign {
  height: 0;
  position: relative;
}

#onekey_sign .onekey_btn {
  position: absolute;
  right: 6px;
  top: -5px;
  width: 101px;
  height: 24px;
  background: url(//tb2.bdstatic.com/tb/static-ihome/widget/onekey_sign/images/btn-nonmember_a67b540.png) no-repeat -280px 0;
}

.ihome_title {
  font-weight: 400;
  font-family: STHeiti,"Microsoft Yahei",Arial,SimSun,sans-serif;
  font-size: 16px;
  color: #494a4c;
  line-height: 18px;
}

.cut_line {
  color: #4585E6;
  font-size: 14px;
  vertical-align: top;
  font-weight: 700;
  line-height: 16px;
  margin-right: 4px;
}

.u-f-wrap {
  height: 118px;
  padding: 3px 0 25px;
  margin: 0 -5px 0 5px;
  position: relative;
}

.u-f-wrap .u-f-item, .always-forum-item .u-f-item {
  line-height: 35px;
  padding-left: 11px;
  padding-right: 30px;
  color: #444;
  float: left;
  height: 35px;
  overflow: hidden;
  width: 124px;
  position: relative;
  margin-right: 10px;
  margin-top: 10px;
  text-decoration: none;
  white-space: nowrap;
}

.u-f-wrap .unsign, .always-forum-item .unsign {
  background: url(//tb2.bdstatic.com/tb/static-ihome/img/btn_groups_ff520ed.png) no-repeat -0px -238px;
}

.u-f-wrap .more {
  text-align: center;
  background: url(//tb2.bdstatic.com/tb/static-ihome/img/btn_groups_ff520ed.png) no-repeat 0 -238px;
  padding: 0;
  width: 165px;
  position: absolute;
  top: 48px;
  right: 0;
  cursor: pointer;
}

.triangle-down {
  width: 10px;
  height: 10px;
  position: absolute;
  right: 58px;
  top: 15px;
  background: url(//tb2.bdstatic.com/tb/static-ihome/img/icon_groups_efa01ec.png) no-repeat 0 -687px;
}

.ihome_section {
  margin: 20px 20px 0;
}

.threadList_title {
  color: #494A4C;
  font-family: STHeiti,"Microsoft Yahei",Arial,SimSun,serif;
  font-size: 16px;
  font-weight: 400;
  line-height: 18px;
}

.cut_line {
  color: #4585E6;
  font-size: 14px;
  font-weight: 700;
  line-height: 16px;
  vertical-align: top;
  margin-right: 4px;
}

</style>
