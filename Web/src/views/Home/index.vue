<template>
  <div class="main-sec">
    <div class="content-sec clearfix">
      <Left></Left>
      <Right></Right>
    </div>
  </div>
</template>

<script>
import Left from "@/views/Home/Left";
import Right from "@/views/Home/Right";

export default {
  components: {
    Left,
    Right
  }
}
</script>

<style scoped lang="less">
.main-sec {
  width: 1000px;
  margin: 0 auto;
  *zoom: 1;
}

.content-sec {
  width: 1000px;
  height: 1000px;
  margin: 0 auto;
  background: url('../../assets/img/big-bg_3854efe.png') repeat-y -4px 0;
  *zoom: 1;
}
</style>

