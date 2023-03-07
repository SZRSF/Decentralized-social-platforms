const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  publicPath: './',
  transpileDependencies: true,
  lintOnSave: true,
  configureWebpack: {
    // 关闭 webpack 的性能提示
    performance: {
      hints: false
    }
  },
  devServer: {
    proxy: {
      '/api': {
        target: 'http://localhost:8082', // 后端接口地址
        changeOrigin: true, // 是否允许跨越
        pathRewrite: {
          '^/api/v1': '/api/v1'// 重写,
        }
      }
    }
  },
  css: {
    loaderOptions: {
      less: {
        // 若 less-loader 版本小于 6.0，请移除 lessOptions 这一级，直接配置选项。
        lessOptions: {
          modifyVars: {
            // 直接覆盖变量
            // 'tabs-default-color': '#CE9FFC', // 标签页下划线颜色
            // 'tabs-bottom-bar-width': '20px', // 标签页宽度
            // 'tabs-line-height': '30px', // 标签页高度
            'cell-border-color': '#fff' // cell下划线颜色
            // 'tabs-nav-background-color': '#f4f7f7'
          }
        }
      }
    }
  }
})
