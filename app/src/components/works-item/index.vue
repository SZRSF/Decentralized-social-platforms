<template>
  <van-cell-group  class="works-cell-group" inset>
    <van-cell class="works-item">
      <div slot="title" class="title">
        <van-image
          round
          width="1rem"
          height="1rem"
          src="https://img01.yzcdn.cn/vant/cat.jpeg"
        />
      </div>
      <div slot="label">
        <div>
          <span>{{ works.author_name }}</span>
          <span>{{works.create_time}}</span>
        </div>
      </div>
    </van-cell>
  </van-cell-group>

</template>

<script>
export default {
  name: 'WorksItem',
  components: {},
  props: {
    works: {
      type: Object,
      require: true
    }
  }
}
</script>

<style scoped>

</style>
