/**
 *  PostCSS 配置文件
 */

module.exports = {
  // 配置要使用的 PostCSS 插件
  plugins: {
    // 配置使用 autoprefixer 插件
    // 作用：生成浏览器 CSS 样式规则前缀
    // VueCLT 内部已经配置
    // 'autoprefixer': { // autoprefixer 插件的配置
    //   // 配置要兼容到的环境
    //   browsers: ['Android >= 4.0', 'iOS >= 8']
    // },

    // 配置使用 postcss-pxtorem 插件
    // 作用：把 px 转为 rem
    'postcss-pxtorem': {
      // lib-flexible 的 REM 适配方案：把一行分为 10 份，每份十分之一
      // 所所以 rootValue 应该设置为你的设计稿的十分之一
      rootValue({ file }) {
        return file.indexOf('vant') !== -1 ? 37.5 : 75
      },

      // 配置要转换的 CSS 属性
      // * 表示所有
      propList: ['*'],
    },
  },
};
