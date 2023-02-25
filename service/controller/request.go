package controller

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

const CtxUserIDKey = "userID"

var ErrorUserNotLogin = errors.New("用户未登录")

// getCurrentUserID 获取当前登录用户ID
func getCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

// 获取分页参数
func getPageInfo(c *gin.Context) (int64, int64) {
	pageNumStr := c.Query("page")
	sizeStr := c.Query("size")
	fmt.Println(pageNumStr, sizeStr)

	var (
		page int64
		size int64
		err  error
	)
	page, err = strconv.ParseInt(pageNumStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 20
	}
	return page, size
}
