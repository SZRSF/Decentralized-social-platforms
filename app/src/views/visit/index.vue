<template>
  <div class="visit-container">
    <van-tabs class="visit-tabs"  animated swipeable  v-model="activeName">
      <van-tab class="visit-tab-my" title="我的家" name="我的家">
        <van-grid
          direction="horizontal"
          :icon-size="80"
          :column-num="2" >
          <van-grid-item
            v-for="family in myFamilyList"
            :key="family.id"
            :icon="family.headImg"
            :text="family.name"/>
        </van-grid>
      </van-tab>
      <van-tab class="visit-tab-family" title="家广场" name="家广场">
        <van-grid
          direction="horizontal"
          :icon-size="80"
          :column-num="2" >
          <van-grid-item
            v-for="family in familyList"
            :key="family.id"
            :icon="family.headImg"
            :text="family.name"/>
        </van-grid>
      </van-tab>
    </van-tabs>
  </div>
</template>

<script>
import { getFamilyList, getMyFamilyList } from '@/api/family'

export default {
  name: 'VisitIndex',
  data () {
    return {
      activeName: '我的家',
      familyList: [],
      myFamilyList: []
    }
  },
  created () {
    this.loadFamily()
    this.myLoadFamily()
  },
  methods: {
    async loadFamily () {
      try {
        const { data } = await getFamilyList()
        this.familyList = data.data
      } catch (err) {
        this.$toast('获取家列表失败')
      }
    },
    async myLoadFamily () {
      try {
        const { data } = await getMyFamilyList(this.$store.state.user.user.user_id)
        this.myFamilyList = data.data
      } catch (err) {
        this.$toast('获取家列表失败')
      }
    }
  }
}
</script>

<style scoped lang="less">
.visit-container {
  .visit-tabs {
  }
}

</style>
