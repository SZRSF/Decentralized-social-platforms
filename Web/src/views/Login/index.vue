<template>
  <div class="login">
    <img :src="imgSrc" width="100%" height="100%" alt="" />
    <div class="loginPart">
      <h2>用户登录</h2>
  <!--  登录表单区域    -->
      <el-form ref="loginFormRef" :model="loginForm" :rules="loginFormRules">
        <el-form-item prop="username" class="inputElement">
          <el-input v-model="loginForm.username" prefix-icon="el-icon-user-solid"  placeholder="请输入用户名/手机号 "></el-input>
        </el-form-item>
        <el-form-item prop="password" class="inputElement">
          <el-input v-model="loginForm.password" prefix-icon="el-icon-lock"  placeholder="请输入密码 " type="password"></el-input>
        </el-form-item>
        <div>
          <el-button type="primary"  @click.prevent="userLogin">登录</el-button>
        </div>
        <div style="text-align: right;color: white;">
          <el-link type="warning">没有账号？去注册</el-link>
        </div>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      imgSrc: require('../../assets/login2.jpg'),
      //登录表单的数据绑定对象
      loginForm: {
        username:'',
        password:''
      },
      // 这是表单的验证规则
      loginFormRules: {
        //验证用户名是否合法
        username:[
          { required: true, message: '请输入用户名或者手机号', trigger: 'blur' },
          { min: 3, max: 11, message: '长度在 3 到 11 个字符', trigger: 'blur' }
        ],
        //验证密码是否合法
        password:[
          { required: true, message: '请输入登录密码', trigger: 'blur' },
          { min: 6, max: 15, message: '长度在 6 到 15 个字符', trigger: 'blur' }
        ]
      }
    }
  },
  methods: {
    // 登录回调函数
    async userLogin() {
      try {
        //登录成功
        const username = this.loginForm.username;
        const password = this.loginForm.password;
        username && password && (await this.$store.dispatch('userLogin',{username,password}))
        //跳转到home首页
        await this.$router.push('/home')
      } catch (error){
        alert(error.message);
      }
      // this.$refs.loginFormRef.validate(async valid =>{
      //   if(!valid) return;
      //   const { data:res } = await this.$http.post('login',this.loginForm);
      //   if(res.code !== 1000) return this.$message.error('登录失败!');
      //   this.$message.success("登录成功");
      //   //1.将登录成功之后的 token, 保存到客户端的 sessionStorage 中
      //   //  1.1 项目中出了登录之外的其外API接触，必须在登录之后才能访问
      //   //  1.2 token只应在当前网站打开期间生效，所以将token保存在 sessionStorage中
      //   window.sessionStorage.setItem('token',res.data);
      //   //2. 通过编程式导航跳转到后台，路由地址是 /home
      //   await this.$router.push('/home');
      // })
    }
  }
}
</script>

<style scoped lang="less">
.loginPart{
  position:absolute;
  /*定位方式绝对定位absolute*/
  top:50%;
  left:50%;
  /*顶和高同时设置50%实现的是同时水平垂直居中效果*/
  transform:translate(-50%,-50%);
  /*实现块元素百分比下居中*/
  width:450px;
  padding:50px;
  background: rgba(0,0,0,.5);
  /*背景颜色为黑色，透明度为0.8*/
  box-sizing:border-box;
  /*box-sizing设置盒子模型的解析模式为怪异盒模型，
  将border和padding划归到width范围内*/
  box-shadow: 0px 15px 25px rgba(0,0,0,.5);
  /*边框阴影  水平阴影0 垂直阴影15px 模糊25px 颜色黑色透明度0.5*/
  border-radius:15px;
  /*边框圆角，四个角均为15px*/
}
.loginPart h2{
  margin:0 0 30px;
  padding:0;
  color: #fff;
  text-align:center;
  /*文字居中*/
}
.loginPart .inputbox{
  position:relative;
}
.loginPart .inputElement input{
  width: 100%;
  padding:10px 0;
  font-size:16px;
  color:#fff;
  letter-spacing: 1px;
  /*字符间的间距1px*/
  margin-bottom: 30px;
  border:none;
  border-bottom: 1px solid #fff;
  outline:none;
  /*outline用于绘制元素周围的线
  outline：none在这里用途是将输入框的边框的线条使其消失*/
  background: transparent;
  /*背景颜色为透明*/
}

.login{
  width:100%;
  height:100%;
}
</style>

