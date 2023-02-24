<template>
  <div class="layout-container">
<!--    子路由出口 -->
    <router-view/>
<!--    /子路由出口-->

<!--    标签导航栏-->
    <van-tabbar class="layout-tabbar" route>
      <van-tabbar-item icon="home-o" to="/home">首页</van-tabbar-item>
      <van-tabbar-item icon="guide-o" to="/visit">串门</van-tabbar-item>
      <van-tabbar-item icon="comment-o" to="/message">消息</van-tabbar-item>
      <van-tabbar-item icon="user-o" to="/mine">{{$store.state.user ? '我的' : '未登录'}}</van-tabbar-item>
    </van-tabbar>
<!--    /标签导航栏-->
  </div>
</template>

<script>
export default {
  name: 'LayoutIndex',
  data () {
    return {
    }
  }
}
</script>

<style scoped>

</style>
