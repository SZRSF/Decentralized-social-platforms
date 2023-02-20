<template>
  <div>
    <Header v-show="$route.meta.show"></Header>
    <!--     路由组件出口的地方-->
    <router-view></router-view>
    <!-- 在Home、Search显式的，在登录、注册隐藏   -->
    <Footer v-show="$route.meta.show"></Footer>
  </div>
</template>

<script>
//引入组件
import Header from "@/components/Header";
import Footer from "@/components/Footer"

export default {
  components: {
    Header,
    Footer,
  }
}
</script>

<style>

</style>
