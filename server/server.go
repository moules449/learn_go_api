package server

import(
	"github.com/gin-gonic/gin"

	"work/controller"
)

func Init() {
	r := router()
	r.Run()
}

//　関数名(引数) 戻り値の型
func router() *gin.Engine {
	r := gin.Default()

	u := r.Group("/users")
	{
		ctrl := user.Controller{}
		u.GET("", ctrl.Index)
		u.GET("/:id", ctrl.Show)
		u.POST("", ctrl.Create)
		u.PUT("/:id", ctrl.Update)
		u.DELETE("/:id", ctrl.Delete)
	}

	return r
}