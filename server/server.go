package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"anaconda/config"
	"anaconda/product"
	// "xapiens.id/geosurvey/config"
	// "xapiens.id/geosurvey/exploration"
	// geologyModel "xapiens.id/geosurvey/geology-model"
	// "xapiens.id/geosurvey/master"
	// notifService "xapiens.id/geosurvey/notification"
	// "xapiens.id/geosurvey/survey"
	// uploadExcel "xapiens.id/geosurvey/upload-excel"
	// "xapiens.id/shared/content"
	// "xapiens.id/shared/health"
	// notifClient "xapiens.id/shared/notification"
	// "xapiens.id/shared/swaggerui"
)

// Start start the server
func Start() error {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	config := config.Get()

	// contentClient := content.New(config.ContentHost, config.RetryCount, config.DebugMode)
	// notificationClient := notifClient.Init(config.NotificationHost, config.RetryCount, config.DebugMode, config.NotificationType)

	// masterService := master.New(contentClient)
	// uploadExcelService := uploadExcel.New()
	// notificationService := notifService.New(notificationClient)
	productService := product.NewService()
	// explorationService := exploration.New(masterService, contentClient)
	// geologyModelService := geologyModel.NewService(masterService, contentClient)

	// master.Register(router, masterService)
	product.Register(router, productService)
	// exploration.Register(router, explorationService, uploadExcelService, contentClient)
	// geologyModel.Register(router, geologyModelService, uploadExcelService, contentClient)
	// uploadExcel.Register(router, uploadExcelService)

	// health.Register(router)

	// router.Static("/images", "./images")
	// router.StaticFile("/swagger.json", "./swagger.json")

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port),
		Handler: router,
	}

	go func() {
		log.Printf("Server starting on port :%d\n", config.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	return handleShutdown(srv)
}

func handleShutdown(srv *http.Server) error {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error
	if err = srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
	return err
}
