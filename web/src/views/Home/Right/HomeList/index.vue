<!-- 热门家推荐-->
<template>
  <div class="home_list">
    <ul id="example-1">
      <li class="home_item" v-for="family in familyList" :key="family.id">
        <a class="home_link"  target="_blank">
          <el-image
              class="home_logo"
              :src="family.headImg"
              fit="fill"></el-image>
          <div class="home_icon">
            <div class="home_name">{{family.name}}</div>
          </div>
        </a>
      </li>
    </ul>
  </div>
</template>

<script>
import {dspFamilyList} from "@/api";

export default {
  name: "HomeList",
  data() {
    return {
      familyList: []
    }
  },
  methods: {
    // 获取家的列表
    getFamilyList() {
      dspFamilyList().then(res => {
        this.familyList = res.data;
      })
    },
  },
  mounted(){
    this.getFamilyList();
  },
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
