<template>
  <div class="post-con" name="post">
    <van-nav-bar title="发布作品"  @click-left="onClickLeft">
      <template #left>
        <van-icon name="cross" />
      </template>
      <template #right>
        <van-button round block  size="mini" type="primary" color="#7232dd" plain native-type="submit" @click="onSubmit()">发布</van-button>
      </template>
    </van-nav-bar>
    <van-form @submit="onSubmit">
      <!-- 家选择器 -->
      <van-field
        readonly
        clickable
        name="picker"
        :value="familyName"
        label="家选择"
        placeholder="点击选择家"
        @click="showPicker = true"
      />
      <van-popup v-model="showPicker" position="bottom">
        <van-picker
          show-toolbar
          :columns="family"
          @confirm="onConfirm"
          @cancel="showPicker = false"
        />
      </van-popup>
      <!-- /家选择器 -->

      <!-- 标题 -->
      <van-field
        v-model="title"
        name="标题"
        placeholder="请输入完整作品标题（5-31个字）"
        :rules="[{ pattern, message: '请输入完整作品标题（5-31个字' }]"
      />
      <!-- /标题 -->
      <!-- 内容 -->

      <myTinymce  :value="content"  v-model="content"/>

      <!-- /内容 -->
    </van-form>
  </div>
</template>

<script>
import myTinymce from '@/components/myTinymce'
import { getFamilyList } from '@/api/family'
import { postWorks } from '@/api/works'

export default {
  name: 'PostIndex',
  components: {
    // eslint-disable-next-line vue/no-unused-components
    myTinymce
  },
  data () {
    return {
      familyName: '',
      title: '',
      content: '',
      familyList: [],
      family: [],
      showPicker: false,
      pattern: /^.{5,31}$/
    }
  },
  methods: {
    async onSubmit (values) {
      // 获取表单数据将内容传给后端
      try {
        const { data } = await postWorks({
          familyName: this.familyName,
          title: this.title,
          content: this.content
        })
        console.log(data)
      } catch (err) {
        console.log('请求失败', err)
      }
    },
    onClickLeft () {
      // Toast('返回')
      this.$router.back()
    },
    onConfirm (value) {
      this.familyName = value
      this.showPicker = false
    },
    async loadFamily () {
      try {
        const { data } = await getFamilyList()
        this.familyList = data.data
        for (let i = 0; i < this.familyList.length; i++) {
          this.family[i] = this.familyList[i].name
        }
      } catch (err) {
        this.$toast('获取家列表失败')
      }
    }
  },
  created () {
    this.loadFamily()
  }
}
</script>

<style scoped>

</style>
