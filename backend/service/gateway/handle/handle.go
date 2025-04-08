package handle

import (
	sl "backend/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"log/slog"
	"net/http"
)

type RouteInitializer interface {
	InitRoutes(router *gin.RouterGroup)
}

type ServerAPI struct {
	router   *gin.Engine
	logger   *slog.Logger
	mediaDir string
}

func New(logger *slog.Logger, mediaDir string) *ServerAPI {
	router := gin.New()
	router.Use(
		gin.LoggerWithConfig(gin.LoggerConfig{
			Output: sl.NewLoggerWriter(logger, slog.LevelInfo),
			Formatter: func(params gin.LogFormatterParams) string {
				sl.LogGinRequest(logger, params, slog.LevelInfo)
				return ""
			},
		}),
		gin.Recovery(),
	)
	if mediaDir != "" {
		router.Static("/media", mediaDir)
	}

	return &ServerAPI{
		router:   router,
		logger:   logger,
		mediaDir: mediaDir,
	}
}

func (s *ServerAPI) RegisterHandlers(handlers ...RouteInitializer) {
	apiGroup := s.router.Group("/api/v1")
	{
		for _, handler := range handlers {
			handler.InitRoutes(apiGroup)
		}
	}
	s.router.GET("/healthcheck", s.healthCheck)
}

func (s *ServerAPI) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (s *ServerAPI) GetHTTPHandler() http.Handler {
	return s.router
}

func (s *ServerAPI) EnableCORS() {
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	s.router.Use(func(c *gin.Context) {
		corsHandler.HandlerFunc(c.Writer, c.Request)
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	})
}

func (s *ServerAPI) AddMiddleware(middleware ...gin.HandlerFunc) {
	for _, m := range middleware {
		s.router.Use(m)
	}
}
