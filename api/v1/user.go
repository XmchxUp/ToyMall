package v1

import (
	"xm-mall/pkg/utils"
	"xm-mall/service"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var s service.UserService
	if err := c.ShouldBind(&s); err == nil {
		res := s.Register(c.Request.Context())
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}

func UserLogin(c *gin.Context) {
	var s service.UserService
	if err := c.ShouldBind(&s); err == nil {
		res := s.Login(c.Request.Context())
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}
