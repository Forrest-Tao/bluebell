package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
	"go.uber.org/zap"
)

func CreatePost(p *models.Post) error {
	//生成post_id
	p.ID = snowflake.GetId()
	//保存到数据库
	return mysql.CreatePost(p)
}

func GetPostById(id int64) (data *models.ApiPostDetail, err error) {
	post, err := mysql.GetPostByID(id)
	if err != nil {
		zap.L().Error("mysql.GetPostByID(id) failed", zap.Int64("pid", id), zap.Error(err))
		return
	}

	//根据用户ID查询用户号信息
	user, err := mysql.GetUserById(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserById(post.AuthorID) failed", zap.Int64("AuthorID", post.AuthorID), zap.Error(err))
		return
	}

	//根据社区ID查询社区信息
	communityDetail, err := mysql.GetCommunityDetail(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetail(post.CommunityID) failed", zap.Int64("CommunityID", post.CommunityID), zap.Error(err))
		return
	}
	data = &models.ApiPostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: communityDetail,
	}
	return
}

// GetPostList 获取帖子详情列表
func GetPostList(page, size int64) (data []*models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList(page, size)
	data = make([]*models.ApiPostDetail, 0, len(posts))
	if err != nil {
		zap.L().Error("mysql.GetPostByID(id) failed", zap.Error(err))
		return
	}
	for _, post := range posts {
		//根据用户ID查询用户号信息
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) failed", zap.Int64("AuthorID", post.AuthorID), zap.Error(err))
			continue
		}

		//根据社区ID查询社区信息
		communityDetail, err := mysql.GetCommunityDetail(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetail(post.CommunityID) failed", zap.Int64("CommunityID", post.CommunityID), zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: communityDetail,
		}
		data = append(data, postDetail)
	}
	return
}
