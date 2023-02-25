<template>
  <div class="works-list">
    <van-pull-refresh
      v-model="isFreshLoading"
      :success-text="refreshSuccessText"
      success-duration="1000"
      @refresh="onRefresh">
      <van-list
        v-model="loading"

        :finished="finished"
        finished-text="没有更多了"
        @load="onLoad"
      >
        <works-item
          v-for="(works,index) in list"
          :key="index"
          :works="works"
        />

<!--        <van-cell-->
<!--          v-for="(works,index) in list"-->
<!--          :key="index"-->
<!--          :title="works.title"-->
<!--        />-->
      </van-list>
    </van-pull-refresh>
  </div>
</template>

<script>
import { getWorksList } from '@/api/works'
import WorksItem from '@/components/works-item'

export default {
  name: 'WorksList',
  components: {
    // eslint-disable-next-line vue/no-unused-components
    WorksItem
  },
  props: {
    channel: {
      type: String,
      required: true
    }
  },
  data () {
    return {
      list: [], // 存储列表数组的数组
      loading: false, // 控制加载中 loading 状态
      finished: false, // 控制加载结束的状态
      worksPage: 1, // 页数
      worksSize: 5, // 每页数据
      isFreshLoading: false, // 控制下拉刷新的状态
      refreshSuccessText: '' // 下拉刷新提示文本
    }
  },
  methods: {
    // 初始化或滚动到底部的时候会触发调用 onLoad
    async onLoad () {
      try {
        // 1.异步更新数据
        const { data } = await getWorksList({
          size: this.worksSize,
          page: this.worksPage
        })

        // 2. 把请求结果数据放到 list 数组中
        const results = data.data
        // 数组的展开操作符， 它会把数组元素一个一个拿出来
        this.list.push(...results)

        // 3. 本次数据之和要把加载状态设置为结束
        this.loading = false

        // 4. 判断数据是否全部加载完成
        if (results.length) {
          // 更新获取下一页数据的时间戳
          this.worksPage = this.worksPage + 1
        } else {
          // 没有数据了，将 finished 设置为 true, 不在 load 加载更多
          this.finished = true
        }
      } catch (err) {
        console.log('请求失败', err)
      }
    },

    // 当下拉刷新的时候会触发调用该函数
    async onRefresh () {
      try {
        // 请求获取数据
        // 1.异步更新数据
        const { data } = await getWorksList({
          size: this.worksSize,
          page: this.worksPage
        })

        // 将数据放到列表的顶部
        this.list.unshift(...data.data)
        // 关闭下拉刷新的 loading 状态
        this.isFreshLoading = false
        this.refreshSuccessText = '刷新成功'
      } catch (err) {
        this.refreshSuccessText = '刷新失败'
      }
      this.isFreshLoading = false
    }
  }
}
</script>

<style scoped lang="less">
.works-list {
  height: 79vh;
  overflow-y: auto;
}
</style>
