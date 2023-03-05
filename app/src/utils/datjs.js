import Vue from 'vue'
import dayjs from 'dayjs'

// 加载中文语言包
import 'dayjs/locale/zh-cn'

import relativeTime from 'dayjs/plugin/relativeTime'

// 配置使用相对时间的插件
dayjs.extend(relativeTime)

// dayjs 默认语言是英文，我们这里配置为中文
dayjs.locale('zh-cn') // 全局使用

// 定义应该全局过滤器，然后就可以在任何组件模板中使用了
// 其实过滤器相当于应该全局可用的方法（仅模板使用）
// 参数1： 过滤器名称
// 参数2： 过滤器函数
// 使用方法： {{ 表达式 | 过滤器名称}}
// 管道符前面的表达式的结果作为参数传递到过滤器函数中
// 过滤器的返回值会渲染到使用过滤器的模板位置
Vue.filter('relativeTime', value => {
  return dayjs().to(dayjs(value))
})

// dayjs() 获取当前最新时间
// dayjs().format('YYYY-MM-DD')
