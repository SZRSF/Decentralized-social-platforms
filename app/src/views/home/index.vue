<template>
  <div class="home-container">
    <!-- 选择作品显示类型 -->
    <div class="home-tap">
      <van-tabs class="home-van-tabs" v-model="activeName"  animated swipeable >
        <van-tab title="关注" name="followers"  >
          <!-- 作品列表 -->
          <WorksList channel="followers"/>
          <!-- /作品列表 -->
        </van-tab>
        <van-tab title="推荐" name="recommend" >
          <!-- 作品列表 -->
          <WorksList  channel="recommend"/>
          <!-- /作品列表 -->
        </van-tab>
        <van-tab title="热门" name="hot">
          <!-- 作品列表 -->
          <WorksList  channel="hot"/>
          <!-- /作品列表 -->
        </van-tab>
        <div slot="nav-right" class="search-btn">
          <van-icon name="search" />
        </div>
      </van-tabs>
    </div>
    <!-- /选择作品显示类型 -->
  </div>
</template>

<script>
import WorksList from '@/views/home/components/works-list'

export default {
  name: 'HomeIndex',
  components: {
    // eslint-disable-next-line vue/no-unused-components
    WorksList
  },
  // props: {
  //   channel: {
  //     type: Object,
  //     required: true
  //   }
  // },
  data () {
    return {
      activeName: 'recommend'
    }
  }
}
</script>

<style scoped lang="less">
.home-container{
  padding-top: 100px;
   background-color: #f7f8fa;
  padding-bottom: 100px;
    /deep/ .home-van-tabs {
      .van-tabs__wrap{
        position: fixed;
        top: 3px;
        z-index: 1;
        left: 0;
        right: 0;
        height: 82px;
      }
      .van-tab {
        //min-width: 200px;
        font-size: 30px;
        right: 10px;
        //left: 20px;
      }
      .van-tabs__nav {
        padding-bottom: 0;
      }
      .van-tabs__line {
        bottom: 8px;
        width: 31px;
        height: 6px;
        background-color: #CE9FFC;
      }

      .search-btn {
        position: fixed;
        right: 0;
        display: flex;
        justify-content: center;
        align-items: center;
        width: 66px;
        height: 82px;
      }
  }
}
</style>
