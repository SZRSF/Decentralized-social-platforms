package mysql

import (
	"zengzhicheng/Decentralized-social-platforms/models"
)

// CreatePost 发布作品
func CreatePost(p *models.Works) (err error) {
	sqlStr := `insert into post(post_id, title,content,author_id,family_id) values (?,?,?,?,?)`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.FamilyID)
	return
}

// GetPostById 根据文章id获取作品内容
func GetPostById(pid int64) (post *models.Works, err error) {
	post = new(models.Works)
	sqlStr := `select 
    		post_id,title,content,author_id,family_id,create_time
    		from post
    		where post_id= ?
    		`
	err = db.Get(post, sqlStr, pid)
	return
}

// GetPostList 获取作品列表
func GetPostList(page, size int64) (post []*models.Works, err error) {
	sqlStr := `select 
		post_id,title,content,author_id,family_id,create_time
		from post
		limit ?,?
		`
	post = make([]*models.Works, 0, 2)
	err = db.Select(&post, sqlStr, (page-1)*size, size)
	return
}

// CollectWorks 收藏文章
func CollectWorks(userId, postId int64) error {
	sqlStr := `INSERT INTO user_collection(user_id, post_id, created_at) VALUES (?, ?, NOW())`
	_, err := db.Exec(sqlStr, userId, postId)
	if err != nil {
		return err
	}
	return nil
}

// CancelCollectWorks 取消收藏
func CancelCollectWorks(userId, postId int64) error {
	sqlStr := `DELETE FROM user_collection WHERE user_id = ? AND post_id = ?`
	_, err := db.Exec(sqlStr, userId, postId)
	if err != nil {
		return err
	}
	return nil
}

// IsArticleCollected 检查用户是否收藏该文章
func IsArticleCollected(userId, postId int64) (bool, error) {
	var count int64
	sqlStr := `SELECT COUNT(*) FROM user_collection WHERE user_id = ? AND post_id = ?`
	err := db.Get(&count, sqlStr, userId, postId)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
