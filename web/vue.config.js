const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  lintOnSave:false,
  // 配置代理跨域
  devServer: {
    proxy: {
      '/api/v1': {
        target: 'http://127.0.0.1:8082',
        changeOrigin: true,
      }
    }
  }
})
