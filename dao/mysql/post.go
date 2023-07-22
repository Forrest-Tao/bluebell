package mysql

import (
	"bluebell/models"
)

func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(
                 post_id,title,content,author_id,community_id
)values (?,?,?,?,?)`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

func GetPostByID(id int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select 
    post_id,title,content,author_id,community_id,create_time
	from post
	where post_id= ?
	`
	err = db.Get(post, sqlStr, id)
	return
}

func GetPostList(page, size int64) (posts []*models.Post, err error) {
	sqlStr := `select 
    post_id,title,content,author_id,community_id,create_time
	from post
	limit ?,?
	`
	posts = make([]*models.Post, 0, size)
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	return
}
