package interfaces

import (
	"github.com/gin-gonic/gin"
	"go-admin-beacon/internal/application/sys"
)

type sysParamApi struct {
}

var paramQryExe = sys.NewParamQryExe()

// QryParam 参数查询
func (s *sysParamApi) QryParam(c *gin.Context) {
	response, err := paramQryExe.Execute(c)
	packResponse(c, response, err)
}
