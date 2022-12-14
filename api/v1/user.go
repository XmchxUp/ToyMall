package v1

import (
	"xm-mall/pkg/utils"
	"xm-mall/service"

	"github.com/gin-gonic/gin"
)

func SendEmail(c *gin.Context) {
	var s service.SendEmailService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&s); err == nil {
		res := s.Send(c.Request.Context(), claims.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}

func UploadAvatar(c *gin.Context) {
	var s service.UserService
	file, fileHeader, _ := c.Request.FormFile("file")
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&s); err == nil {
		res := s.Post(c.Request.Context(), claims.ID, file, fileHeader.Size)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}

func UserUpdate(c *gin.Context) {
	var s service.UserService
	claims, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&s); err == nil {
		res := s.Update(c.Request.Context(), claims.ID)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
		utils.LogrusObj.Infoln(err)
	}
}

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
