<template>
  <div class="tinymce-editor">
    <vue-tinymce
      v-model="myValue"
      :setting="setting"
    />
  </div>
</template>

<script>
import { postImage } from '@/api/works'

export default {
  name: 'myTinymce',
  props: {
    value: {
      type: String,
      default: ''
    }
  },
  data () {
    return {
      myValue: this.value,
      setting: {
        height: 430, // 高度
        language: 'zh_CN', // 调用放在langs文件夹内的语言包
        // toolbar: false, // 隐藏工具栏
        menubar: false, // 隐藏菜单栏
        // inline: true, // 开启内联模式
        branding: false, // 隐藏右下角技术支持
        elementpath: false, // 隐藏底部的元素路径,
        content_style: 'img {max-width:100%;}', // 限制图片大小
        toolbar_location: 'bottom', // 工具烂位置
        plugins: ['lists', 'image', 'media', 'table', 'wordcount', 'lineheight', 'emoticons'], // 选择需加载的插件
        // 工具栏
        toolbar: ['lists image media emoticons  formatselect  bold italic forecolor backcolor  table| '],
        images_upload_handler: async function (blobInfo, success, failure) {
          // 文件上传的formData传递，忘记为什么要用这个了
          const formData = new FormData()
          formData.append('file', blobInfo.blob(), blobInfo.filename())
          // 上传的api,和后端配合，返回的是图片的地址，然后加上公共的图片前缀
          const res = await postImage(formData)
          console.log(res)
          if (res.data.code === 1000) {
            const url = res.data.data
            console.log(url)
            success(url)
          } else {
            failure('图片上传失败')
          }
        }
      }
    }
  },
  watch: {
    // 监听内容变化
    value (newValue) {
      this.myValue = newValue
    },
    myValue (newValue) {
      this.$emit('input', newValue)
    }
  }
}
</script>

<style scoped>

</style>
