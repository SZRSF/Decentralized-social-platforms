<template>
  <van-cell-group  class="works-cell-group" inset>
    <van-cell class="works-item">
      <div slot="title" class="title">
        <van-row>
          <van-col span="4">
            <van-image
              round
              width="1rem"
              height="1rem"
              fit="cover"
              :src= "works.family.headImg"
            />
          </van-col>
          <van-col span="8">
            <span> {{works.family.name}} </span>
            <div>
              <span> 关注：  作品： </span>
            </div>
          </van-col>
        </van-row>
      </div>
      <div slot="label">
        <div>
          <span>{{ works.title}}</span>
        </div>
        <div>
          <span>{{ summarize(works.content)}}</span>
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
  },
  methods: {
    // 只显示文章前50个字
    summarize (content) {
      return content.slice(0, 50) + '...'
    }
  }
}
</script>

<style scoped>

</style>
