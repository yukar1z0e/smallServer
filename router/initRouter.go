package router


import (
	"github.com/gin-gonic/gin"
	"smallServer/handler"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Static("/download","./smallweb/")
	//router.Static("/static","./static")
	//router.StaticFile("/favicon.ico","./favicon.ico")
	index:=router.Group("/")
	{
		index.POST("/ddd77eddd/dfdfdsf/dfsfd/jjjdfdf9933/kk", handler.Index)
	}
	return router
}