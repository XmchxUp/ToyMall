package v1

import (
	"xm-mall/pkg/utils"
	"xm-mall/service"

	"github.com/gin-gonic/gin"
)

func ListProductImg(c *gin.Context) {
	s := service.ListProductImgService{}
	if err := c.ShouldBind(&s); err == nil {
		res := s.List(c.Request.Context(), c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}

func ShowProduct(c *gin.Context) {
	s := service.ProductService{}
	res := s.Show(c.Request.Context(), c.Param("id"))
	c.JSON(200, res)
}

func SearchProducts(c *gin.Context) {
	s := service.SearchProductService{}
	if err := c.ShouldBind(&s); err == nil {
		res := s.Search(c.Request.Context())
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}

func ListProducts(c *gin.Context) {
	s := service.ListProductService{}
	if err := c.ShouldBind(&s); err == nil {
		res := s.List(c.Request.Context())
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}

func CreateProduct(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	s := service.ProductService{}
	if err := c.ShouldBind(&s); err == nil {
		res := s.Create(c.Request.Context(), claim.ID, files)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}
