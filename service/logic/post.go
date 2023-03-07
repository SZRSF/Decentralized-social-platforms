package logic

import (
	"zengzhicheng/Decentralized-social-platforms/dao/ipfs"
	"zengzhicheng/Decentralized-social-platforms/dao/mysql"
	"zengzhicheng/Decentralized-social-platforms/models"
	"zengzhicheng/Decentralized-social-platforms/pkg/snowflake"

	"go.uber.org/zap"
)

func CreatePost(p *models.Works) (err error) {
	// 1.生成 post id
	p.ID = snowflake.GenID()
	// 2.将内容保存到IPFS返回HASH值
	p.Content = postContent(p.Content)
	// 3.保存到数据库
	return mysql.CreatePost(p)
	// 3. 返回
}

// GetPostById 根据作品id查询帖子详情
func GetPostById(pid, followerId int64) (data *models.ApiPostDetail, err error) {
	// 查询并组合我们接口想用的数据
	post, err := mysql.GetPostById(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostById(pid) failed", zap.Int64("pid", pid), zap.Error(err))
		return
	}
	// 根据hash值从ipfs获取帖子内容
	post.Content, err = getContent(post.Content)
	// 根据作者uid查询作者信息
	user, err := mysql.GetUserById(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetPostById(post.AuthorID) failed",
			zap.Int64("author_id", post.AuthorID),
			zap.Error(err))
		return
	}
	// 根据家id查询家详细信息
	family, err := mysql.GetFamilyDetailByID(post.FamilyID)
	if err != nil {
		zap.L().Error("mysql.GetFamilyDetailByID(post.FamilyID) failed",
			zap.Int64("author_id", post.FamilyID),
			zap.Error(err))
		return
	}
	// 根据登录用户Id和作品发布者Id获取登录用户是否关注作品发布者
	followingId := post.AuthorID
	isFollowed, err := mysql.GetIsFollowed(followerId, followingId)
	// 查询登录用户是否关注此文章
	isCollected, err := mysql.IsArticleCollected(followerId, post.ID)
	// 接口数据拼接
	data = &models.ApiPostDetail{
		IsCollected:  isCollected,
		IsFollowed:   isFollowed,
		User:         user,
		Works:        post,
		FamilyDetail: family,
	}
	return
}

// GetPostList 获取作品列表
func GetPostList(page, size int64) (data []*models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		return nil, err
	}
	data = make([]*models.ApiPostDetail, 0, len(posts))
	for _, post := range posts {
		// 根据作者iud查询作者信息
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetPostById(post.AuthorID) failed",
				zap.Int64("author_id", post.AuthorID),
				zap.Error(err))
			continue
		}
		// 根据家id查询家详细信息
		family, err := mysql.GetFamilyDetailByID(post.FamilyID)
		if err != nil {
			zap.L().Error("mysql.GetFamilyDetailByID(post.FamilyID) failed",
				zap.Int64("author_id", post.FamilyID),
				zap.Error(err))
			continue
		}
		// 根据hash值获取文章内容
		post.Content, err = ipfs.CatIPFS(post.Content)
		if err != nil {
			zap.L().Error("ipfs.CatIPFS failed", zap.Error(err))
		}
		postDetail := &models.ApiPostDetail{
			User:         user,
			Works:        post,
			FamilyDetail: family,
		}
		data = append(data, postDetail)
	}
	return
}

// GetFamilyID 根据家名称查家id
func GetFamilyID(familyName string) (familyID int64, err error) {
	// 根据家名称查询家ID
	familyID, err = mysql.GetFamilyID(familyName)
	if err != nil {
		zap.L().Error("mysql.GetFamilyDetailByID(post.FamilyID) failed", zap.Error(err))
		return
	}
	return familyID, err
}

// AddCollect 收藏作品
func AddCollect(postId, followerId int64) error {
	err := mysql.CollectWorks(followerId, postId)
	if err != nil {
		zap.L().Error("mysql.CollectWorks(followerId, postId) failed", zap.Error(err))
		return nil
	}
	return nil
}

// DeleteCollect 取消收藏作品
func DeleteCollect(followerId, postId int64) error {
	err := mysql.CancelCollectWorks(followerId, postId)
	if err != nil {
		zap.L().Error("mysql.CancelCollectWorks(followerId, postId) failed", zap.Error(err))
		return nil
	}
	return nil
}
