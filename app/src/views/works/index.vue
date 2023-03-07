<template>
  <div class="works-container">
    <!-- 导航栏 -->
    <van-nav-bar  @click-left="$router.back()" left-arrow>
      <template #title>
        <van-grid direction="horizontal" icon-size="50px" :column-num="1">
          <van-grid-item   :icon="family.headImg"   :text="family.name" />
        </van-grid>
      </template>
      <template #right>
        <van-icon name="ellipsis" size="18" />
      </template>
    </van-nav-bar>
    <!-- /导航栏 -->

    <div class="main-wrap">
      <!-- 加载中 -->
      <div v-if="loading" class="loading-wrap">
        <van-loading
          color="#3296fa"
          vertical
          >加载中</van-loading>
      </div>
      <!-- /加载中 -->

      <!-- 加载完成 文章详情 -->
      <div v-else-if="works.title" class="works-detail">

        <!-- 用户信息 -->
        <van-cell class="user-info" center :border="false">
          <van-image
            class="avatar"
            slot="icon"
            round
            fit="cover"
            :src="user.head_img"
            />
          <div slot="title" class="user-name">{{user.username}}</div>
          <div slot="label" class="publish-data">{{user.create_time | relativeTime}}</div>
          <follow-user
            class="follow-btn"
            :is-followed="works.is_followed"
            :user-id="works.author_id"
            @update-is_followed="works.is_followed = $event"
          />
<!--          <van-button
            v-if="works.is_followed"
            class="follow-btn"
            round
            size="small"
            :loading="followLoading"
            @click="onFollow"
          >已关注</van-button>
          <van-button
            v-else
            class="follow-btn"
            type="info"
            color="#3296fa"
            round
            size="small"
            icon="plus"
            :loading="followLoading"
            @click="onFollow"
          >关注</van-button>-->
        </van-cell>
        <!-- /用户信息 -->

        <!-- 文章标题 -->
        <h1 class="works-title">{{works.title}}</h1>
        <!-- 文章标题 -->

        <!-- 文章内容 -->
        <div
          class="works-content markdown-body"
          v-html="works.content"
          ref="works-content"
        ></div>
        <van-divider>正文结束</van-divider>
        <!-- /文章内容 -->

        <!-- 底部区域 -->
        <div class="works-bottom">
          <van-button
            class="comment-btn"
            type="default"
            round
            size="small"
          >写评论</van-button>
          <van-icon
            name="comment-o"
            info="123"
          />
          <collect-works
            class="btn-item"
            v-model="works.is_collected"
            :works-id="works.works_id"
          />
          <van-button
            class="btn-item"
            icon="good-job-o"
          />
          <van-icon name="share" color="#777777"></van-icon>
        </div>
        <!-- /底部区域 -->
      </div>
      <!-- /加载完成 文章详情 -->

      <!-- 加载失败： 404 -->
      <div v-else-if="errStatus === 404" class="error-wrap">
        <van-icon name="failure" />
        <p class="text">该资源不存在或者已删除！</p>
      </div>
      <!-- /加载失败： 404 -->

      <!-- 加载失败： 其他未知错误（例如网络原因或者服务器原因) -->
      <div  v-else class="error-wrap">
        <van-icon name="failure" />
        <p class="text">内容加载失败！</p>
        <van-button
          class="retry-btn"
          @click="loadWorks"
        >点击重试</van-button>
      </div>
      <!-- /加载失败： 其他未知错误（例如网络原因或者服务器原因) -->
    </div>
  </div>
</template>

<script>
import { getWorksById } from '@/api/works'
import { ImagePreview } from 'vant'
import FollowUser from '@/components/follow-user'
import CollectWorks from '@/components/collect-works'

export default {
  name: 'worksIndex',
  components: {
    FollowUser,
    CollectWorks
  },
  props: {
    worksId: {
      type: [Number, String, Object],
      required: true
    }
  },
  data () {
    return {
      user: {}, // 作者信息
      works: {}, // 文章详情
      family: {}, // 家信息
      loading: true, // 加载中的 loading 状态
      errStatus: 0, // 失败的状态码
      followLoading: false
    }
  },
  computed: {},
  watch: {},
  created () {
    this.loadWorks()
  },
  mounted () {},
  methods: {
    async loadWorks () {
      // 展示 loading 加载中
      this.loading = true
      try {
        const { data } = await getWorksById(this.worksId)

        // 数据驱动视图不是立即的
        this.works = data.data
        this.user = data.data.user
        this.family = data.data.family

        // 初始化图片点击预览
        setTimeout(() => {
          this.previewImage()
        }, 0)

        // 请求成功， 关闭 loading
        // this.loading = false
      } catch (err) {
        if (err.data && err.data.code === 1004) {
          this.errStatus = 404
        }
      }
      // 无论成功还是失败，都需要关闭 loading
      this.loading = false
    },

    previewImage () {
      // 得到所有的img 节点
      const worksContent = this.$refs['works-content']
      const imgs = worksContent.querySelectorAll('img')
      const images = []
      imgs.forEach((img, index) => {
        images.push(img.src)
        img.onclick = () => {
          ImagePreview({
            images,
            // 起始位置， 从 0 开始
            startPosition: index
          })
        }
      })
    }

  }
}
</script>

<style scoped lang="less">
@import "./github-markdown.css";

.works-container {
  .main-wrap {
    position: fixed;
    left: 0;
    right: 0;
    top: 0;
    bottom: 0;
    background-color: #fff;
  }
  .works-detail {
    position: fixed;
    left: 0;
    right: 0;
    top: 92px;
    bottom: 88px;
    overflow-y: scroll;
    background-color: #fff;
    .works-title {
      font-size: 40px;
      padding: 50px 32px;
      margin: 0;
      color: #3a3a3a;
    }

    .user-info {
      padding: 0 32px;
      .avatar {
        width: 70px;
        height: 70px;
        margin-right: 17px;
      }
      .van-cell__label {
        margin-top: 0;
      }
      .user-name {
        font-size: 24px;
        color: #3a3a3a;
      }
      .publish-date {
        font-size: 23px;
        color: #b7b7b7;
      }
      .follow-btn {
        width: 170px;
        height: 58px;
      }
    }

    .works-content {
      padding: 55px 32px;
      /deep/ p {
        text-align: justify;
      }
    }
  }

  .loading-wrap {
    padding: 200px 32px;
    display: flex;
    align-items: center;
    justify-content: center;
    background-color: #fff;
  }

  .error-wrap {
    padding: 200px 32px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    background-color: #fff;
    .van-icon {
      font-size: 122px;
      color: #b4b4b4;
    }
    .text {
      font-size: 30px;
      color: #666666;
      margin: 33px 0 46px;
    }
    .retry-btn {
      width: 280px;
      height: 70px;
      line-height: 70px;
      border: 1px solid #c3c3c3;
      font-size: 30px;
      color: #666666;
    }
  }

  .works-bottom {
    position: fixed;
    left: 0;
    right: 0;
    bottom: 0;
    display: flex;
    justify-content: space-around;
    align-items: center;
    box-sizing: border-box;
    height: 88px;
    border-top: 1px solid #d8d8d8;
    background-color: #fff;
    .comment-btn {
      width: 282px;
      height: 46px;
      border: 2px solid #eeeeee;
      font-size: 30px;
      line-height: 46px;
      color: #a7a7a7;
    }
    /deep/ .van-icon {
      font-size: 40px;
    }
    .comment-icon {
      top: 2px;
      color: #777;
      .van-info {
        font-size: 16px;
        background-color: #e22829;
      }
    }
    .btn-item {
      border: none;
      padding: 0;
      height: 40px;
      line-height: 40px;
      color: #777777
    }
    .collect-btn--collected {
      color: #ffa500;
    }
    .like-btn--liked {
      color: #e5645f;
    }
  }
}
</style>
