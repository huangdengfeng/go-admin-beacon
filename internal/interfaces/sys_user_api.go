package interfaces

import (
	"github.com/gin-gonic/gin"
	"go-admin-beacon/internal/application/sys"
	"go-admin-beacon/internal/infrastructure/errors"
	"go-admin-beacon/internal/infrastructure/response"
	"net/http"
)

type sysUserApi struct {
}

var userPageQryExe = sys.NewUserPageQryExe()

func (s *sysUserApi) UserPageQry(c *gin.Context) {
	qry := &sys.UserPageQry{}
	if err := c.ShouldBindJSON(qry); err != nil {
		c.JSON(http.StatusOK, response.ErrorWithCodeMsg(errors.BadArgs.Code, err.Error()))
		return
	}
	response := userPageQryExe.Execute(c, qry)
	c.JSON(http.StatusOK, response)
}
