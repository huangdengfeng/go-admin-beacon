package interfaces

import (
	"github.com/gin-gonic/gin"
	"go-admin-beacon/internal/application/sys"
	"go-admin-beacon/internal/infrastructure/errors"
	"go-admin-beacon/internal/infrastructure/response"
	"net/http"
)

type sysRoleApi struct {
}

var roleListQryExe = sys.NewRoleListQryExe()

func (s *sysRoleApi) qryRoleList(c *gin.Context) {
	var qry sys.RoleListQry
	if err := c.ShouldBindJSON(&qry); err != nil {
		c.JSON(http.StatusOK, response.ErrorWithCodeMsg(errors.BadArgs.Code, err.Error()))
		return
	}
	response, err := roleListQryExe.Execute(c, &qry)
	packResponse(c, response, err)
}
