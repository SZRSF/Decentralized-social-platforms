<template>
  <header class="header">
  <!-- 头部的第一行  -->
    <div class="top">
      <div class="container">
        <div class="loginList">
          <p>
            <span>请</span>
            <!-- 声明式导航：务必要有to属性 -->
            <router-link to="/login">登录</router-link>
            <router-link class="register" to="/register">注册</router-link>
<!--            <el-button type="info" @click="logout">退出</el-button>-->
            <el-link @click="logout" target="_blank">退出</el-link>
          </p>
        </div>
      </div>
    </div>
  <!-- 头部第二行 搜素区域-->
      <div class="bottom">
        <h1 class="logoArea">
          <router-link class="logo" to="/home"></router-link>
            <img src="###" alt="">
        </h1>
        <div class="searchArea">
          <form action="###" class="searchFrom">
            <input
              type="text"
              id="autocomplete"
              class="input-error"
              v-model="keyword"
            >
            <button class="sui-btn" type="button" @click="goSearch">
              搜素
            </button>
          </form>
        </div>
      </div>
  </header>
</template>

<script>
export default {
  name: "",
  data(){
    return{
      keyword:''
    }
  },
  methods: {
    //搜索按钮的回调函数：需要向search路由进行跳转
    goSearch(){
      //路由传递参数:
      this.$router.push({name:"search",params:{keyword:this.keyword},query:{k:this.keyword.toUpperCase()}})
    },
    logout() {
      window.sessionStorage.clear()
      this.$router.push('/login')
    }
  }
}
</script>

<style scoped lang="less">
.header {
  & > .top {
    background-color: #eaeaea;
    height: 30px;
    line-height: 30px;

    .container {
      width: 1200px;
      margin: 0 auto;
      overflow: hidden;

      .loginList {
        float: right;

        p {
          float: left;
          margin-left: 10px;

          .register {
            border-left: 1px solid #b3aeae;
            padding: 0 5px;
            margin-left: 5px;
          }
        }
      }
    }
  }
    & > .bottom {
      width: 1200px;
      margin: 0 auto;
      overflow: hidden;

      .logoArea {
        float: left;

        .logo {
          img {
            width: 175px;
            margin:25px 45px;
          }
        }
      }

      .searchArea {
        float:right;
        margin-top: 35px;

        .searchFrom {
          overflow: hidden;

          .input {
            box-sizing: border-box;
            width: 490px;
            height: 32px;
            padding: 4px;
            border: 2px solid #ea4a36;
            float: left;

            &:focus {
              outline: none;
            }
          }

          .button {
            height: 32px;
            width: 68px;
            background-color: #ea4a36;
            border:none;
            color: #fff;
            float: left;
            cursor: pointer;

            &:focus {
              outline: none;
            }
          }
        }
      }
    }
}
</style>
