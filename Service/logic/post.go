package logic

import (
	"zengzhicheng/Decentralized-social-platforms/dao/mysql"
	"zengzhicheng/Decentralized-social-platforms/models"
	"zengzhicheng/Decentralized-social-platforms/pkg/snowflake"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	// 1.生成 post id
	p.ID = snowflake.GenID()
	// 2.保存到数据库
	return mysql.CreatePost(p)
	// 3. 返回
}

// GetPostById 根据作品id查询帖子详情
func GetPostById(pid int64) (data *models.ApiPostDetail, err error) {
	// 查询并组合我们接口想用的数据
	post, err := mysql.GetPostById(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostById(pid) failed", zap.Int64("pid", pid), zap.Error(err))
		return
	}
	// 根据作者iud查询作者信息
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
	// 接口数据拼接
	data = &models.ApiPostDetail{
		AuthorName:   user.Username,
		Post:         post,
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
		postDetail := &models.ApiPostDetail{
			AuthorName:   user.Username,
			Post:         post,
			FamilyDetail: family,
		}
		data = append(data, postDetail)
	}
	return
}
