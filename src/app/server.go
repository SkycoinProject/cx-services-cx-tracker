package app

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/SkycoinPro/cx-services-cx-tracker/src/api"
)

// Server model
type Server struct {
	Engine *gin.Engine
}

// NewServer - server initialization
func NewServer(ctrls ...api.Controller) *Server {
	if viper.GetBool("server.release-mode") {
		gin.SetMode(gin.ReleaseMode)
	}

	server := &Server{
		Engine: gin.Default(),
	}
	server.initCors()
	server.initRoutes(ctrls...)
	return server
}

// Run - run server
func (s *Server) Run() {
	if err := s.Engine.Run(serverAddress()); err != nil {
		panic(err.Error())
	}
}

func (s *Server) initCors() {
	s.Engine.Use(cors.New(cors.Config{
		AllowHeaders:    viper.GetStringSlice("c0rs.allowed-headers"),
		AllowMethods:    viper.GetStringSlice("c0rs.allowed-methods"),
		AllowAllOrigins: true,
		MaxAge:          viper.GetDuration("c0rs.max-age"),
	}))
}

func (s *Server) initRoutes(ctrls ...api.Controller) {
	publicAPIGroup := s.Engine.Group("/api/v1")
	closedAPIGroup := publicAPIGroup.Group("")

	// use ginSwagger middleware to
	//publicAPIGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	for _, controller := range ctrls {
		controller.RegisterAPIs(publicAPIGroup, closedAPIGroup)
	}
}

func serverAddress() string {
	return fmt.Sprintf("%s:%s", viper.GetString("server.ip"), viper.GetString("server.port"))
}
