package server

import (
	apiV1 "go-gravatar/api/v1"
	"go-gravatar/docs"
	"go-gravatar/internal/handler"
	"go-gravatar/internal/middleware"
	"go-gravatar/pkg/jwt"
	"go-gravatar/pkg/log"
	"go-gravatar/pkg/server/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewHTTPServer(
	logger *log.Logger,
	conf *viper.Viper,
	jwt *jwt.JWT,
	userHandler *handler.UserHandler,
	avatarHandler *handler.AvatarHandler,
) *http.Server {
	gin.SetMode(gin.DebugMode)
	s := http.NewServer(
		gin.Default(),
		logger,
		http.WithServerHost(conf.GetString("http.host")),
		http.WithServerPort(conf.GetInt("http.port")),
	)

	// swagger doc
	docs.SwaggerInfo.BasePath = "/v1"
	s.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		//ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", conf.GetInt("app.http.port"))),
		ginSwagger.DefaultModelsExpandDepth(-1),
		ginSwagger.PersistAuthorization(true),
	))

	s.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
		//middleware.SignMiddleware(log),
	)

	// health check
	s.GET("/", func(ctx *gin.Context) {
		logger.WithContext(ctx).Debug("[/] alive")
		apiV1.HandleSuccess(ctx, map[string]interface{}{
			"healthStatus": "alive",
		})
	})
	s.GET("/health", func(ctx *gin.Context) {
		logger.WithContext(ctx).Debug("[/health] alive")
		apiV1.HandleSuccess(ctx, map[string]interface{}{
			"healthStatus": "alive",
		})
	})

	//
	s.GET("/avatar/:hash", avatarHandler.GetAvatar)

	v1 := s.Group("/v1")
	{
		// No route group has permission
		noAuthRouter := v1.Group("/")
		{
			noAuthRouter.POST("/register", userHandler.Register)
			noAuthRouter.POST("/login", userHandler.Login)
			noAuthRouter.POST("/reset", userHandler.Reset)
		}
		// Non-strict permission routing group
		noStrictAuthRouter := v1.Group("/").Use(middleware.NoStrictAuth(jwt, logger))
		{
			noStrictAuthRouter.GET("/user", userHandler.GetProfile)
		}
		// Strict permission routing group
		strictAuthRouter := v1.Group("/").Use(middleware.StrictAuth(jwt, logger))
		{
			strictAuthRouter.PUT("/user", userHandler.UpdateProfile)
			strictAuthRouter.DELETE("/user", userHandler.Delete)
			strictAuthRouter.PUT("/avatar", avatarHandler.UpdateAvatar)
			strictAuthRouter.DELETE("/avatar", avatarHandler.DeleteAvatar)
		}
	}

	return s
}
