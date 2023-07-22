package mysql

import (
	"bluebell/models"
	encrypt "bluebell/pkg/encript"
	"database/sql"
)

// CheckUserExist 查看用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

func InsertUser(user models.User) (err error) {
	sqlStr := `insert into user(user_id,username,password)values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

func Login(user *models.User) (err error) {
	//查看用户是否存在  查看用户的密码是否正确
	opassword := user.Password
	sqlStr := `select user_id,username,password from user where username= ?`
	err = db.Get(user, sqlStr, user.Username)
	//用户不存在
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		return err
	}
	//先取得明文 将明文转为密文 免去了密码解密的func
	password := encrypt.EncryptPassword(opassword)
	if password != user.Password {
		return ErrorInvalidPassword
	}
	return
}

func GetUserById(uid int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select user_id,username from user where user_id = ?`
	if err = db.Get(user, sqlStr, uid); err != nil {
		return
	}
	return
}
