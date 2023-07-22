package controller

import (
	"bluebell/dao/mysql"
	"bluebell/logic"
	"bluebell/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	//获取参数 参数校验
	p := new(models.ParamsSignUp)
	fmt.Printf("user: %#v\n", *p)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("ShouldBind SignUp with invalid param", zap.Error(err))
		//判断 error 是不是 validator.validatorErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			zap.L().Info("err.(validator.ValidationErrors)", zap.Error(errs))
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInValidPassword, removeTopStruct(errs.Translate(trans))) // 翻译错误
		return
	}
	//fmt.Printf("user: %#v\n", *p)

	//业务处理
	if err := logic.SignUp(p); err != nil {
		// 记录错误日志
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	} // 3. 返回响应
	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	//获取参数 校验参数
	p := new(models.ParamsLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		//请求参数有误
		zap.L().Error("Login with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInValidPassword, removeTopStruct(errs.Translate(trans))) // 翻译错误
		return
	}

	//业务处理
	user, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login() failed ", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		//密码错误
		ResponseError(c, CodeInValidPassword)
		return
	}
	//设置 Authorization 头部
	c.Request.Header.Set("Authorization", user.Token)
	ResponseSuccess(c, gin.H{
		"user_id":  fmt.Sprintf("%d", user.UserID), // 前端 id 值最大为 2^53-1；后端 int64 最大值为 2^63-1；不一致可能会导致失真
		"username": user.Username,
		"token":    user.Token,
	})
}
