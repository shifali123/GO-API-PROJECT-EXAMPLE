package router

import (
	"github.com/fast-user/fast/pkg/fast-server/handler"
"github.com/globalsign/mgo"
	"github.com/gin-gonic/gin"
)

// AppRouter returns the Gin router that handles API requests...

func FastUserRouter(db *mgo.Session,dbName string) *gin.Engine {
	router := gin.Default()
	createUserHandler := handler.NewCreateFastUserHandler(db,dbName)
		//router.GET("/health", func(ctx *gin.Context) {
		//	ctx.JSON(http.StatusOK, gin.H{"success": true})
	router.GET("/user/:id", 	createUserHandler.GetFastUserByID)
	router.POST("/user", 	createUserHandler.CreateFastUser)

	return router
}
