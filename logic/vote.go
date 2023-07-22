package logic

import (
	"bluebell/models"
	"go.uber.org/zap"
)

// 投票功能
// 本项目使用简化版的投票分数
// 投一票就加 432 分 86400/200 —> 需要 200 张赞成票可以给你的帖子续一天

// 1. 判断投票的限制
// 2. 更新帖子的分数
// 3. 记录用户为该帖子投票的分数
func VoteForPost(userID int64, p *models.ParamsVoteData) (err error) {
	zap.L().Debug("VoteForPost",
		zap.Int64("userID", userID),
		zap.Int64("postID", p.PostID),
		zap.Int8("direction", p.Direction))
	// 逻辑判断 是怎样的投票

}
