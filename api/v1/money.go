package v1

import (
	"xm-mall/pkg/utils"
	"xm-mall/service"

	"github.com/gin-gonic/gin"
)

func ShowMoney(c *gin.Context) {
	s := service.ShowMoneyService{}
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&s); err == nil {
		res := s.Show(c.Request.Context(), claim.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}
