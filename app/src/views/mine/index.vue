<template>
  <div class="mine-container">
    <!-- 已登录头部 -->
    <div v-if="user" class="header user-info">
      <div class="base-info">
        <div class="left">
          <van-image
            class="avatar"
            :src="userInfo.head_img"
            round
            fit="cover"
          />
          <div class="userinfo-middle">
            <div class="userinfo-title">
              <span class="userinfo_username">{{userInfo.username}}</span>
            </div>
            <div class="userinfo-nums">
              <span class="userinfo-num">关注 {{userInfo.follow_count}}</span>
              <span class="userinfo-num">粉丝 {{userInfo.fans_count}}</span>
            </div>
          </div>
        </div>
        <div class="right">
          <van-button size="mini" round>编辑资料</van-button>
        </div>
      </div>
     <!-- 导航 -->
      <div class="data-stats">
        <div class="data-item">
          <span class="count">{{userInfo.works_count}}</span>
          <span class="text">我的作品</span>
        </div>
        <div class="data-item">
          <span class="count">{{userInfo.joined_family}}</span>
          <span class="text">加入的家</span>
        </div>
        <div class="data-item">
          <span class="count">{{userInfo.collect_count}}</span>
          <span class="text">收藏</span>
        </div>
        <div class="data-item">
          <span class="count">{{userInfo.browsing_history}}</span>
          <span class="text">浏览记录</span>
        </div>
      </div>
      <!-- /导航 -->
    </div>
    <!-- /已登录头部 -->

    <!-- 未登录头部 -->
    <div v-else class="header not-login">
      <div class="login-btn" @click="$router.push('/login')">
        <img class="mobile-img" src="@/assets/mobile.png" alt="">
        <span class="text">登录 / 注册</span>
      </div>
    </div>
    <!-- /未登录头部 -->

    <!-- 功能 -->
    <van-cell-group class="gn-van-cell-group" title="辅助功能"  inset>
      <van-cell class="gn-van-cell" title="服务中心"  is-link/>
      <van-cell  class="gn-van-cell" title="隐私设置" is-link/>
      <van-cell  class="gn-van-cell" title="社区公约"  is-link/>
      <van-cell  v-if="user" class="gn-van-cell" title="更多" is-link/>
    </van-cell-group>
    <!-- /功能 -->

    <!-- 退出 -->
    <van-cell-group  v-if="user" class="logout-cell-group" inset>
      <van-cell class="logout-cell" title="退出登录"  clickable  @click="onLogout" />
    </van-cell-group>
    <!-- /退出 -->
  </div>
</template>

<script>
import { mapState } from 'vuex'
import { getUserInfo } from '@/api/user'

export default {
  name: 'MineIndex',
  data () {
    return {
      userInfo: {} // 用户信息
    }
  },
  computed: {
    ...mapState(['user'])
  },
  watch: {},
  created () {
    // 如果用户登录了,则请求加载用户数据
    if (this.user) {
      this.loadUserInfo()
    }
  },
  methods: {
    onLogout () {
      // 退出提示
      this.$dialog.confirm({
        title: '确定要退出吗？'
      })
        .then(() => {
          // on confirm
          // 确认退出： 清除登录状态(‘容器中的user + 本地存储中的user)
          this.$store.commit('setUser', null)
        })
        .catch(() => {
          // on cancel
        })
    },
    async loadUserInfo () {
      try {
        const { data } = await getUserInfo(this.user.user.user_id)
        this.userInfo = data.data
      } catch (err) {
        this.$toast('获取数据失败， 请稍后重试')
      }
    }
  }
}
</script>

<style scoped lang="less">
.mine-container {
  background-color: #f7f8fa;
  .header {
    height: 361px;
    background-image: linear-gradient(to top, #ebc0fd 0%, #d9ded8 100%);
    background-size: cover;
  }

  .not-login {
    display: flex;
    justify-content: center;
    align-items: center;
    .login-btn {
      display: flex;
      flex-direction: column;
      justify-content: center;
      align-items: center;
      .mobile-img {
        width: 132px;
        height: 132px;
        margin-bottom: 15px;
      }
      .text {
        font-size: 28px;
        color: #fff;
      }
    }
  }

  .user-info {
    .base-info {
      height: 231px;
      padding: 76px 32px 23px;
      box-sizing: border-box;
      display: flex;
      justify-content: space-between;
      align-items: center;
      background-image: linear-gradient(-225deg, #E3FDF5 0%, #FFE6FA 100%);
      .left {
        display: flex;
        align-items: center;
        .avatar {
          width: 132px;
          height: 132px;
          margin-right: 23px;
        }
        .userinfo-middle {
          .userinfo-title{
            .userinfo_username{
              font-size: 30px;
              color: #252c41;
              font-weight: bold;
            }
          }
          .userinfo-nums{
            .userinfo-num{
              font-size: 20px;
              color: #8b8687;
              padding-left: 0;
              padding-right: 8px;
            }
          }
        }
      }
    }
    .data-stats {
      height: 130px;
      background-color: #ffffff;
      display: flex;
      .data-item {
        height: 130px;
        flex: 1;
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        .count{
          font-size: 36px;
        }
        .text{
          font-size: 23px;
          color: #a3a1a1;
        }
      }
    }
  }

  .gn-van-cell-group {
    .gn-van-cell {
      margin-bottom: 9px;
      //position: static;
    }
  }

  .logout-cell-group{
    top: 20px;
    .logout-cell {
      text-align: center;
      color: #d86262;
    }
  }
}
</style>
