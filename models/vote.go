package models

type ParamsVoteData struct {
	//UserID int64 `json:"user_id"`
	PostID    int64 `json:"post_id,string" binding:"required"`       // 帖子 id
	Direction int8  `json:"direction,string" binding:"oneof=1 0 -1"` //赞成 反对 取消投票
}
