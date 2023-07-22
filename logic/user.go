package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	encrypt "bluebell/pkg/encript"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
	"fmt"
	"go.uber.org/zap"
)

func SignUp(p *models.ParamsSignUp) (err error) {
	//判断用户是否存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		zap.L().Error("mysql.CheckUserExist() failed :", zap.Error(err))
		return
	}
	//分配 ID值
	userID := snowflake.GetId()
	//创建用户
	user := models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	fmt.Printf("%#v\n", user)
	//密码加密
	user.Password = encrypt.EncryptPassword(user.Password)
	//插入数据
	err = mysql.InsertUser(user)

	return
}

func Login(p *models.ParamsLogin) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	if err = mysql.Login(user); err != nil {
		return nil, err
	}
	fmt.Printf("%+v\n", user)
	//生成jwt
	token, err := jwt.GenToken(user.UserID, user.Username)
	//token 出错
	if err != nil {
		return nil, err
	}
	user.Token = token
	return
}
