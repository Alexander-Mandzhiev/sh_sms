package handle

import (
	sl "backend/pkg/logger"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"time"
)

type RouteInitializer interface {
	InitRoutes(router *gin.RouterGroup)
}

type ServerAPI struct {
	router   *gin.Engine
	logger   *slog.Logger
	mediaDir string
	env      string
	frontend string
}

func New(logger *slog.Logger, mediaDir string, env string, frontendAddr string) *ServerAPI {
	router := gin.New()
	router.Use(ginLoggerMiddleware(logger), gin.Recovery())

	server := &ServerAPI{
		router:   router,
		logger:   logger,
		mediaDir: mediaDir,
		env:      env,
		frontend: frontendAddr,
	}

	server.setupCORS()

	if mediaDir != "" {
		router.Static("/media", mediaDir)
	}

	return server
}

func (s *ServerAPI) setupCORS() {
	config := cors.Config{
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           12 * time.Hour,
	}

	if s.env == "development" {
		config.AllowAllOrigins = true
		s.logger.Warn("CORS: Allowing all origins in development mode")
	} else {
		config.AllowOrigins = []string{
			fmt.Sprintf("http://%s", s.frontend),
			fmt.Sprintf("https://%s", s.frontend),
		}
	}

	s.router.Use(cors.New(config))
}

func ginLoggerMiddleware(logger *slog.Logger) gin.HandlerFunc {
	return gin.LoggerWithConfig(gin.LoggerConfig{
		Output: sl.NewLoggerWriter(logger, slog.LevelInfo),
		Formatter: func(params gin.LogFormatterParams) string {
			sl.LogGinRequest(logger, params, slog.LevelInfo)
			return ""
		},
	})
}

func (s *ServerAPI) RegisterHandlers(handlers ...RouteInitializer) {
	apiGroup := s.router.Group("/api/v1")
	{
		for _, handler := range handlers {
			handler.InitRoutes(apiGroup)
		}
	}
	s.router.GET("/health", s.healthCheck)
}

func (s *ServerAPI) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (s *ServerAPI) GetHTTPHandler() http.Handler {
	return s.router
}

func (s *ServerAPI) AddMiddleware(middleware ...gin.HandlerFunc) {
	for _, m := range middleware {
		s.router.Use(m)
	}
}

func (s *ServerAPI) GetEnv() string {
	return s.env
}
