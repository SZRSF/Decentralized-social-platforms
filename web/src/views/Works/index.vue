<template>
  <div class="works_home_1">
    <div class="works_home_2">
        <el-container>
          <el-tabs v-model="activeName"  type="card" @tab-click="handleClick">
            <el-tab-pane label="首页" name="home">首页</el-tab-pane>
            <el-tab-pane label="作品列表" name="works">作品列表</el-tab-pane>
          </el-tabs>
          <el-row class="btn_works_fb">
            <el-button type="primary" icon="el-icon-edit" @click="$router.push('/publish')">发布</el-button>
          </el-row>
        </el-container>
    </div>
  </div>
</template>

<script>
export default {
  name: "Works",
  data() {
    return {
      activeName: 'works',
      tabPosition: 'left'
    };
  },
  methods: {
    handleClick(tab, event) {
      console.log(tab, event);
    }
  }
};
</script>

<style scoped lang="less">
.works_home_1 {

}

.works_home_2 {
  width: 978px;
  margin: 0 auto;
}

.btn_works_fb {
  left: 800px;
}
</style>
