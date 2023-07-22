package redis

const (
	KeyPrefix          = "bluebell:"
	KeyPostTimeZSet    = "post:time"   //帖子及发帖时间
	KeyPostScoreZSet   = "post:score"  //帖子及投票的分数
	KeyPostVotedPrefix = "post:voted:" //记录用户及投票类型 参数是 post id类型
)

func GetKeys(str string) string {
	return KeyPrefix + str
}
