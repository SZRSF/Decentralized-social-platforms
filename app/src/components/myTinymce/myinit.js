// import { postImage } from '@/api/works'
//
// export default {
//   selector: '#textarea', // 容器，可使用css选择器
//   height: 430, // 高度
//   language: 'zh_CN', // 调用放在langs文件夹内的语言包
//   // toolbar: false, // 隐藏工具栏
//   menubar: false, // 隐藏菜单栏
//   // inline: true, // 开启内联模式
//   branding: false, // 隐藏右下角技术支持
//   elementpath: false, // 隐藏底部的元素路径,
//   content_style: 'img {max-width:100%;}', // 限制图片大小
//   toolbar_location: 'bottom', // 工具烂位置
//   plugins: ['lists', 'image', 'media', 'table', 'wordcount', 'lineheight', 'emoticons'], // 选择需加载的插件
//   // 工具栏
//   toolbar: ['lists image media emoticons  formatselect  bold italic forecolor backcolor  table| '],
//   images_upload_handler: (blobInfo, success) => {
//     this.upload(blobInfo, (e) => {
//       success(e)
//     })
//   },
//   upload (blobInfo, fn) {
//     const formData = new FormData()
//     formData.append('file', blobInfo.blob())
//     // 这里为自己的上传接口调用方法
//     const res = postImage({
//       formData
//     })
//     if (res.code === 200) {
//       fn && fn(res.result)
//     } else {
//       this.$message.error('图片上传失败')
//       fn && fn('')
//     }
//   }
// }
