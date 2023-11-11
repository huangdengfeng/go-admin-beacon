package interfaces

import (
	"github.com/gin-gonic/gin"
	"go-admin-beacon/internal/application/sys"
	"net/http"
)

type sysParamApi struct {
}

var paramQryExe = sys.NewParamQryExe()

// QryParam 参数查询
func (s *sysParamApi) QryParam(c *gin.Context) {
	response := paramQryExe.Execute(c)
	c.JSON(http.StatusOK, response)
}
