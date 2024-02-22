package server

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"go-gin-server/internal/config"
	"go-gin-server/internal/handler"
	"go-gin-server/internal/logger"
	"go-gin-server/internal/middleware"
)

// Server links handlers to paths via routes.
type Server struct {
	router *gin.Engine
}

// New creates server with gin framework.
func New() (*Server, error) {
	// Dispatch info from config.
	SetLogLevel()

	s := &Server{
		router: gin.New(),
	}

	// Apply required middleware.
	s.router.Use(middleware.RequestLogger())
	s.router.Use(middleware.ResponseLogger())

	// Register defined handlers.
	s.RegisterHandlers()

	return s, nil
}

// Run the gin server on routes.
func (s *Server) Run() {
	port := config.ServerHTTPPort()
	logger.Log.Info("Starting HTTP server on port: ", port)
	s.router.Run(":" + fmt.Sprintf("%d", port))
}

// RegisterHandlers links handlers to API points.
func (s *Server) RegisterHandlers() {
	s.router.GET("/health", handler.Health)
}

// SetLogLevel uses info from config to adjust web server parameters.
func SetLogLevel() {
	switch config.LogLevel() {
	case config.LogLevelNormal:
		gin.SetMode(gin.ReleaseMode)
	case config.LogLevelDebug:
		gin.SetMode(gin.DebugMode)
	}
}
