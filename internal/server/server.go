package server

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"go-gin-server/internal/cache"
	"go-gin-server/internal/config"
	"go-gin-server/internal/handler/system"
	"go-gin-server/internal/handler/entity/tag"
	"go-gin-server/internal/logging"
	"go-gin-server/internal/middleware"
	"go-gin-server/internal/repository"
)

// Server links handlers to paths via routes.
type Server struct {
	router *gin.Engine
}

// New creates server with gin framework.
func New(version string) (*Server, error) {
	system.SetVersion(version)
	repository.InitWithMongo(config.RepositoryDBName(), config.RepositoryURL())
	cache.InitWithRedis(config.CacheRedisAddress(), config.CacheRedisPassword(), config.CacheRedisDB())
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
func (s *Server) Run() error {
	port := config.ServerHTTPPort()
	logging.Logger.Info("Starting HTTP server", slog.String("port", port))
	return s.router.Run(":" + port)
}

// RegisterHandlers links handlers to API points.
func (s *Server) RegisterHandlers() {
	s.router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// System operations
	s.router.GET("/system/health", system.HealthHandler)
	s.router.GET("/system/version", system.VersionHandler)

	// CRUD tag item operations

	// by primary key access
	s.router.POST("/entity/tag", tag.CreateHandler)
	s.router.GET("/entity/tag/:id", tag.ReadOneHandler)
	s.router.PUT("/entity/tag/:id", tag.UpdateHandler)
	s.router.DELETE("/entity/tag/:id", tag.DeleteHandler)

	// by bulk access
	s.router.GET("/entity/tags", tag.ReadHandler)
	s.router.DELETE("/entity/tags", tag.DropHandler)
}
