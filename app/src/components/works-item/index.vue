<template>
  <van-cell-group  class="works-cell-group" inset>
    <van-cell
      class="works-item"
      :to="{
      // 根据路由名称跳转
          name: 'works',
          // 传递路由动态参数
          params: {
            // 属性名： 路由路径中设计的动态参数名称
            worksId: works.works_id
          }
      }"
    ><div slot="title" class="title">
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
            <span > {{works.family.name}} </span>
            <div>
              <span class="works-item-text"> 关注： 作品： </span>
            </div>
          </van-col>
        </van-row>
      </div>
      <div slot="label">
        <div>
          <span class="works-item-title">{{ works.title}}</span>
        </div>
        <div v-html="summarize(works.content)">
<!--          <span v-html="{{ summarize(works.content)}}">{{ summarize(works.content)}}</span>-->
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
      return content.slice(0, 100) + '...'
    }
  }
}
</script>

<style scoped lang="less">
.works-cell-group{
  .works-item {
    .works-item-title{
      font-size: 30px;
      font-weight: 550;
      color: #2a3746;
    }
    .works-item-text {
      font-size: 23px;
      color: #a3a1a1;
    }
  }
}

</style>
