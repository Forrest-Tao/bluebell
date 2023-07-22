package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

func CreatePostHandler(c *gin.Context) {
	//获取参数 参数校验
	p := new(models.Post)
	fmt.Println(p)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("invalid params with c.ShouldBindJSON(p)\n")
		zap.L().Error("c.ShouldBindJSON() failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	authorID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("getCurrentUserID(c) failed", zap.Error(err))
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = authorID

	// 业务处理
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, nil)
}

func GetPostByIDHandler(c *gin.Context) {
	idStr := c.Param("id")
	pid, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	post, err := logic.GetPostById(pid)
	if err != nil {
		zap.L().Error("logic.GetPostById(id) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, post)
}

func GetPostListHandler(c *gin.Context) {
	page, size := getPageSize(c)
	posts, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, posts)
}
