<template>
  <div id="container" class="container_wrap clearfix ihome_body">
    <LeftAside></LeftAside>
  </div>
</template>

<script>
import LeftAside from "@/views/User/Container/LeftAside";

export default {
  components: {
    LeftAside,
  }
}
</script>

<style scoped lang="less">
.ihome_body {
  border-bottom: 1px solid #87BED9;
}

.container_wrap {
  width: 980px;
  margin: 0 auto;
}

</style>
