<!-- 右边内容组件 -->
<template>
  <div class="right-sec">
    <div class="r_top_sec">
      <div id="home_hot" class="home_hot">
        <Top></Top>
        <HomeList></HomeList>
      </div>
    </div>
    <div class="r-left-sec">
      <Left></Left>
    </div>
  </div>
</template>

<script>
import Top from "@/views/Home/Right/Top";
import HomeList from "@/views/Home/Right/HomeList";
import Left from "@/views/Home/Right/Left";

export default {
  name: "Right",
  components: {
    Top,
    HomeList,
    Left,
  }
}
</script>

<style scoped>
.right-sec {
  float: left;
  margin-left: 10px;
  width: 780px;
  *zoom: 1;
  *margin-left: 8px;
}

.home_hot {
  padding: 16px 16px 0;
  width: 780px;
  height: 300px;
  background: #fff;
  border: 1px solid rgba(230,230,230,1);
  box-sizing: border-box;
  font-family: STHeiti,"Microsoft Yahei","Microsoft YaHei",Arial,sans-serif;
}

.r-left-sec {
  float: left;
  width: 550px;
  overflow: hidden;
  min-height: 1000px;
}

</style>
