package routes

import (
	"net/http"
	api "xm-mall/api/v1"
	"xm-mall/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(middleware.Cors())
	r.Use(sessions.Sessions("mysessions", store))
	r.StaticFS("/static", http.Dir("./static"))

	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(ctx *gin.Context) {
			ctx.JSON(200, "success")
		})

		v1.POST("user/register", api.UserRegister)
	}
	return r
}
