package server

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"go-gin-server/internal/config"
	"go-gin-server/internal/handler"
	"go-gin-server/internal/handler/task"
	"go-gin-server/internal/log"
	"go-gin-server/internal/middleware"
)

// Server links handlers to paths via routes.
type Server struct {
	router *gin.Engine
}

// New creates server with gin framework.
func New(version string) (*Server, error) {
	handler.SetVersion(version)

	gin.SetMode(gin.ReleaseMode)

	s := &Server{
		router: gin.New(),
	}

	// Apply required middleware. The order matters as request log should
	// go before response log.
	s.router.Use(middleware.ResponseLogger())
	s.router.Use(middleware.RequestLogger())

	// Register defined handlers.
	s.RegisterHandlers()

	return s, nil
}

// Run the gin server on routes.
func (s *Server) Run() {
	port := config.ServerHTTPPort()
	log.Logger.Info("Starting HTTP server", slog.String("port", port))
	s.router.Run(":" + port)
}

// RegisterHandlers links handlers to API points.
func (s *Server) RegisterHandlers() {
	// System operations
	s.router.GET("/health", handler.HealthHandler)
	s.router.GET("/version", handler.VersionHandler)
	s.router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// CRUD task item operations
	s.router.POST("/task", task.CreateHandler)
	s.router.GET("/task/:id", task.ReadHandler)
	s.router.PUT("/task/:id", task.UpdateHandler)
	s.router.DELETE("/task/:id", task.DeleteHandler)
}
