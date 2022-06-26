package httpserver

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/saratchandra13/sampleProject/pkg/domain/services"
	"github.com/saratchandra13/sampleProject/third_party/platlogger"
	"log"
	"net/http"
)

const (
	port = ":8000"
)

type interactor struct {
	appLogic services.AppInterface
	logger   *platlogger.Client
}

func (i *interactor) populate(appLogic services.AppInterface, logger *platlogger.Client) {
	i.appLogic = appLogic
	i.logger = logger
}

var appInteractor interactor
var srv *http.Server

func NewServer(appLogic services.AppInterface, logger *platlogger.Client) {
	appInteractor.populate(appLogic, logger)

	router := gin.Default()

	// health check route
	router.GET("/health", func(c *gin.Context) {
		c.String(200, "Health Check")
	})

	// inject routes
	v1 := router.Group("/v1")
	{
		v1.GET("/listVegetable", listVegetable)
		v1.POST("/addVegetable", addVegetable)
	}

	srv = &http.Server{
		Addr:    port,
		Handler: router,
	}
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
}

func Shutdown(ctx context.Context) {
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Error in Shutdown", err)
	}
}
