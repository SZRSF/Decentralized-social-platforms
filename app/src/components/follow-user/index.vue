<template>
  <van-button
    v-if="isFollowed"
    class="follow-btn"
    round
    size="small"
    :loading="loading"
    @click="onFollow"
  >已关注</van-button>
  <van-button
    v-else
    class="follow-btn"
    type="info"
    color="#3296fa"
    round
    size="small"
    icon="plus"
    :loading="loading"
    @click="onFollow"
  >关注</van-button>
</template>

<script>
import { addFollow, deleteFollow } from '@/api/user'

export default {
  name: 'FollowUser',
  components: {},
  props: {
    isFollowed: {
      type: [Number, Boolean],
      required: true
    },
    userId: {
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
    async onFollow () {
      this.loading = true // 展示关注按钮的loading 状态
      try {
        if (this.isFollowed) {
          // 已关注, 取消关注
          await deleteFollow(this.userId)
          // this.works.is_followed = 0
        } else {
          // 没有关注， 添加关注
          await addFollow({
            target: this.userId
          })
          // this.works.is_followed = 1
        }
        // this.works.is_followed = !this.works.is_followed
        this.$emit('update-is_followed', !this.isFollowed)
      } catch (err) {
        let message = '操作失败，请重试'
        if (err.response && err.response.status === 400) {
          message = '你不能关注你自己!'
        }
        this.$toast(message)
      }
      this.loading = false // 关闭关注按钮的loading 状态
    }
  }
}
</script>

<style scoped>

</style>
