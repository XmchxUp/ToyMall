package v1

import (
	"xm-mall/pkg/utils"
	"xm-mall/service"

	"github.com/gin-gonic/gin"
)

func ListCategories(c *gin.Context) {
	s := service.ListCategoriesService{}
	if err := c.ShouldBind(&s); err == nil {
		res := s.List(c.Request.Context())
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}
