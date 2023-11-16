package interfaces

import (
	"github.com/gin-gonic/gin"
	"go-admin-beacon/internal/application/sys"
	"go-admin-beacon/internal/application/sys/auth"
	"go-admin-beacon/internal/infrastructure/errors"
	"go-admin-beacon/internal/infrastructure/response"
	"net/http"
)

type sysUserApi struct {
}

var userPageQryExe = sys.NewUserPageQryExe()
var myInfoExe = sys.NewMyInfoExe()
var addUserCmdExe = sys.NewAddUserCmdExe()

// UserPageQry 用户分页查询
func (s *sysUserApi) UserPageQry(c *gin.Context) {
	var qry sys.UserPageQry
	if err := c.ShouldBindJSON(&qry); err != nil {
		c.JSON(http.StatusOK, response.ErrorWithCodeMsg(errors.BadArgs.Code, err.Error()))
		return
	}
	response, err := userPageQryExe.Execute(c, &qry)
	packResponse(c, response, err)
}

// My 用户个人信息
func (s *sysUserApi) My(c *gin.Context) {
	context := auth.SetUserToContext(c)
	response, err := myInfoExe.Execute(context)
	packResponse(c, response, err)
}

// AddUser 添加用户
func (s *sysUserApi) AddUser(c *gin.Context) {
	var cmd sys.AddUserCmd
	if err := c.ShouldBindJSON(&cmd); err != nil {
		c.JSON(http.StatusOK, response.ErrorWithCodeMsg(errors.BadArgs.Code, err.Error()))
		return
	}
	response, err := addUserCmdExe.Execute(c, &cmd)
	packResponse(c, response, err)
}
