<template>
  <div class="login-container">
    <van-icon class="login-van-icon" name="arrow-left" @click="$router.back()"/>
<!--    登录方式选择-->
    <van-tabs  class="login-van-tabs"  v-model="activeName">
      <van-tab title="账号登录" name="账号登录">
        <!--  登录表单  -->
        <van-form @submit="onSubmit">
          <van-field
            v-model="user.username"
            name="用户名"
            label="用户名"
            placeholder="用户名"
            :rules="userFormRules.username"
          />
          <van-field
            v-model="user.password"
            type="password"
            name="密码"
            label="密码"
            placeholder="密码"
            :rules="[{ required: true, message: '请填写密码' }]"
          />
          <div style="margin: 16px;">
            <van-button :disabled="btnState === false" :class="{'login-btnBg':btnState}" round block type="info" native-type="submit">
              登录
            </van-button>
          </div>
        </van-form>
        <!--  /登录表单  -->
      </van-tab>
      <van-tab title="手机登录" name="手机登录">手机登录</van-tab>
    </van-tabs>
<!--    /登录方式选择-->
  </div>
</template>

<script>
import { login } from '@/api/user'

export default {
  name: 'LoginIndex',
  data () {
    return {
      activeName: '账号登录',
      user: {
        username: '', // 用户名
        password: '' // 账号密码
      },
      userFormRules: {
        username: [{
          required: true,
          message: '用户名不能为空'
        }],
        password: [{ required: true, message: '密码不能为空' }]
      }
    }
  },
  methods: {
    async onSubmit () {
      // 1. 获取表单数据
      const user = this.user
      // 2. 表单验证
      this.$toast.loading({
        message: '登录中...',
        forbidClick: true, // 禁用背景点击
        duration: 0 // 持续时间， 默认是 2000 , 如果为0 则持续展示
      })
      try {
        // 3. 提交表单请求登录
        const { data } = await login(user)
        if (data.code === 1000) {
          this.$store.commit('setUser', data.data)
          this.$toast.success('登录成功')

          // 登录成功，跳转原来页面
          this.$router.back()
        } else if (data.code === 1004) {
          this.$toast.fail('用户名或密码错误')
        } else {
          this.$toast.fail('登录失败，请稍后重试')
        }
      } catch (err) {
        this.$toast.fail(err)
      }

      // 4. 根据请求响应结果处理后续查操作
    }
  },
  computed: {
    btnState () { // 当用户名和密码框都不为空时btnState==true,利用这个计算属性来动态控制按钮的禁用和颜色
      return this.user.username !== '' && this.user.password !== ''
    }
  }
}
</script>

<style scoped lang="less">

.login-van-tabs{
  top: 150px;
}

/* 登录按钮禁用时的颜色 */
.van-button {
  background: #ee9b86;
}

/* 登录按钮可以使用时的颜色 */
.login-btnBg {
  background: #fc5531;
}

.login-van-icon {
  top: 40px;
  left: 20px;
}
</style>
