<template>
  <div class="works-list">
    <van-list
      v-model="loading"
      :finished="finished"
      finished-text="没有更多了"
      @load="onLoad"
    >
      <van-cell
        v-for="works in list"
        :key="works.id"
        :title="works.title"
      />
    </van-list>
  </div>
</template>

<script>
import { getWorksList } from '@/api/works'

export default {
  name: 'WorksList',
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
      timestamp: null // 请求获取下一页数据的时间戳
    }
  },
  methods: {
    // onLoad () {
    //   // 1.异步更新数据
    //   // setTimeout 仅做示例，真实场景中一般为 ajax 请求
    //   setTimeout(() => {
    //     // 2. 把请求结果数据放到 list 数组中
    //     for (let i = 0; i < 10; i++) {
    //       this.list.push(this.list.length + 1)
    //     }
    //
    //     // 3. 本次数据之和要把加载状态设置为结束
    //     this.loading = false
    //
    //     // 4. 判断数据是否全部加载完成
    //     if (this.list.length >= 40) {
    //       // 如果没有数据了， 把finished 设置为True, 之和不在触发加载更多
    //       this.finished = true
    //     }
    //   }, 1000)
    // }
    async onLoad () {
      try {
        // 1.异步更新数据
        const { data } = await getWorksList()

        // 2. 把请求结果数据放到 list 数组中
        const results = data.data
        // 数组的展开操作符， 它会把数组元素一个一个拿出来
        this.list.push(...results)

        // 3. 本次数据之和要把加载状态设置为结束
        this.loading = false

        // 4. 判断数据是否全部加载完成
        if (results.length) {
          // 更新获取下一页数据的时间戳
        } else {
          // 没有数据了，将 finished 设置为 true, 不在 load 加载更多
          this.finished = true
        }
      } catch (err) {
        console.log('请求失败', err)
      }
    }
  }
}
</script>

<style scoped>

</style>
