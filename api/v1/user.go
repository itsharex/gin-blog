package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/ginblog/model"
	"github.com/ginblog/utils/errmsg"
	"github.com/ginblog/utils/validator"
	"net/http"
	"strconv"
)

var code int

//// 查询用户是否存在
//func UserExist(c *gin.Context) {
//	userName := c.Query("username")
//	code := model.CheckUser(userName)
//	if code != errmsg.SUCCESS {
//		c.JSON(http.StatusOK, gin.H{
//			"status":  code,
//			"message": errmsg.GetErrMsg(code),
//		})
//	} else {
//		c.JSON(http.StatusOK, gin.H{
//			"status":  code,
//			"message": errmsg.GetErrMsg(code),
//		})
//
//	}
//}

// 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	var msg string
	_ = c.ShouldBindJSON(&data)
	msg, code = validator.Validate(&data)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": msg,
		})
		return
	}

	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询单个用户
func GetUserInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   1,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	username := c.Query("username")

	if pageNum == 0 {
		pageNum = 1
	}
	data, total := model.GetUsers(username, pageSize, pageNum)
	code = errmsg.SUCCESS // 200
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑用户
func EditUser(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code = model.CheckUpUser(id, data.Username)
	if code == errmsg.SUCCESS {
		model.EditUser(id, &data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func UpdatePassword(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	
	code = model.CheckUpUser(id, data.Username)
	if code == errmsg.SUCCESS {
		model.UpdatePassword(id, &data)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
