package models

import "time"

/*type Community struct {
	ID            int64  `db:"id"`
	CommunityID   string `db:"community_nid"`
	CommunityName string `db:"community_name"`
	Introduction  string `db:"introduction"`
	CreateTime    int64  `db:"create_time"`
	UpdateTime    int64  `db:"update_time"`
}*/

type Community struct {
	ID   int64  `json:"id" db:"community_id"`
	Name string `json:"name" db:"community_name"`
}

type CommunityDetail struct {
	ID           int64     `json:"id" db:"community_id"`
	Name         string    `json:"name" db:"community_name"`
	Introduction string    `json:"introduction,omitempty" db:"introduction"`
	CreateTime   time.Time `json:"create_time" db:"create_time"`
}
