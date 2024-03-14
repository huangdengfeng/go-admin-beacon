package interfaces

import (
	"github.com/gin-gonic/gin"
	"go-admin-beacon/internal/application/sys"
	"go-admin-beacon/internal/infrastructure/errors"
	"go-admin-beacon/internal/infrastructure/response"
	"net/http"
)

var permissionListQryExe = sys.NewPermissionListQryExe()

type sysPermissionApi struct {
}

func (s *sysPermissionApi) qryPermissionList(c *gin.Context) {
	var qry sys.PermissionListQry
	if err := c.ShouldBindJSON(&qry); err != nil {
		c.JSON(http.StatusOK, response.ErrorWithCodeMsg(errors.BadArgs.Code, err.Error()))
		return
	}
	response, err := permissionListQryExe.Execute(c, &qry)
	packResponse(c, response, err)
}
