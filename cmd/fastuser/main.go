package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/fast-user/fast/pkg/fast-server/router"
"github.com/fast-user/fast/pkg/fast-server/store"
	"github.com/gin-gonic/gin"
)

func main() {

session, dbName := store.GetSessionForMongo()
	approuter := router.FastUserRouter(session,dbName)
	// Start up the server.
	//logger.Info("Starting Server")

	//	go startJobHandler(sns, sqs, db)

	startServer(approuter)
}
func startServer(router *gin.Engine) {

	var err error

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	//logger.Info("Server listening")
	go func() {
		err = server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			//		logger.Fatal(errors.Trace(err))
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	//logger.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = server.Shutdown(ctx)
	if err != nil {
		//logger.Fatalf("Server Shutdown: %s", errors.Details(err))
	}
	//logger.Info("Server exiting")
}
