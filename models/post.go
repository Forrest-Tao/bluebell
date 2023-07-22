package models

import "time"

// 内存对齐
type Post struct {
	Title       string    `db:"title" json:"title" biding:"required"`
	Content     string    `db:"content" json:"content" biding:"required"`
	ID          int64     `db:"post_id" json:"id"`
	AuthorID    int64     `db:"author_id" json:"author_id"`
	CommunityID int64     `db:"community_id" json:"community_id" biding:"required"`
	Status      int8      `db:"status" json:"status"`
	CreateTime  time.Time `db:"create_time" json:"create_time"`
}

type ApiPostDetail struct {
	AuthorName       string `json:"author_name"`
	*Post            `json:"post"`
	*CommunityDetail `json:"community_detail"`
}
