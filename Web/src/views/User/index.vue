<template>
  <div class="wrap1" >
    <div class="wrap2">
      <HeadInfo></HeadInfo>
      <UserInfo></UserInfo>
      <Container></Container>
    </div>
  </div>
</template>

<script>
import HeadInfo from "@/views/User/HeadInfo";
import UserInfo from "@/views/User/UserInfo";
import Container from "@/views/User/Container"

export default {
  name: "index",
  components: {
    HeadInfo,
    UserInfo,
    Container,
  }
}
</script>

<style scoped lang="less">
.wrap1 {
  min-width: 980px;
  overflow: hidden;
  margin: 0 auto;
  background-position: center 10px;
  background-repeat: no-repeat;
  position: relative;

  background-image: url('../../assets/img/user/background-image.png');
}

body > .wrap1 > .wrap2 {
  z-index: 2;
  position: relative;
  background: transparent;
}
</style>
