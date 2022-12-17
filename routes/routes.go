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
		v1.POST("user/login", api.UserLogin)

		v1.GET("carousels", api.ListCarousels)
		v1.GET("products", api.ListProducts)
		v1.POST("products", api.SearchProducts)
		v1.GET("product/:id", api.ShowProduct)
		v1.GET("imgs/:id", api.ListProductImg)

		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.PUT("user", api.UserUpdate)
			authed.POST("avatar", api.UploadAvatar)
			authed.POST("user/sending-email", api.SendEmail)
			authed.POST("user/valid-email", api.ValidEmail)

			authed.POST("money", api.ShowMoney)

			authed.POST("product", api.CreateProduct)
		}
	}
	return r
}
