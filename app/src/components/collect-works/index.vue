<template>
  <van-button
    :icon="value ? 'star' : 'star-o'"
    :class="{
      collected: value
    }"
    @click="onCollect"
  />
</template>

<script>
import { addCollect, deleteCollect } from '@/api/works'

export default {
  name: 'CollectWorks',
  props: {
    value: {
      type: Boolean,
      required: true
    },
    worksId: {
      type: [Number, String, Object],
      required: true
    }
  },
  data () {
    return {
      loading: false
    }
  },
  methods: {
    async onCollect () {
      this.loading = true
      try {
        if (this.value) {
          // 已收藏, 取消收藏
          await deleteCollect(this.worksId)
        } else {
          // 没有收藏，添加收藏
          await addCollect({
            worksId: this.worksId
          })
        }

        // 更新视图
        this.$emit('input', !this.value)

        this.$toast.success(!this.value ? '收藏成功' : '取消收藏')
      } catch (err) {
        this.$toast.fail('操作失败， 请重试')
      }
      this.loading = false
    }
  }
}
</script>

<style scoped lang="less">
.collected {
  .van-icon {
    color: #ffa500;
  }
}
</style>
