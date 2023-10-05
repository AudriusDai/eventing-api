package web

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/audriusdai/eventing-api/config"
	"github.com/audriusdai/eventing-api/web/route"
	"github.com/gin-gonic/gin"
)

func SetupEngine() *gin.Engine {
	engine := gin.New()
	apiGroup := engine.Group("/api")

	route.SetupRoutes(apiGroup)

	return engine
}

func SetupWeb() error {
	server := &http.Server{
		Addr:    config.APP_ADDRESS,
		Handler: SetupEngine(),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// wait for interrupt signal to gracefully shutdown the server with a timeout
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-quit
	log.Println("preparing server for shutdown")
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(config.GRACEFUL_SHUTDOWN_TIMEOUT_SECONDS)*time.Second,
	)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		return err
	}
	<-ctx.Done()
	log.Println("graceful shutdown time has passed - exiting the server")

	return nil
}
