package delivery

import (
	"github.com/artomsopun/clendry/clendry-api/internal/config"
	v1 "github.com/artomsopun/clendry/clendry-api/internal/delivery/http/v1"
	"github.com/artomsopun/clendry/clendry-api/internal/service"
	"github.com/artomsopun/clendry/clendry-api/pkg/auth"
	"github.com/artomsopun/clendry/clendry-api/pkg/files"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type Handler struct {
	services     *service.Services
	tokenManager auth.TokenManager
	filesManager files.Files
}

func NewHandler(services *service.Services, tokenManager auth.TokenManager, filesManager files.Files) *Handler {
	return &Handler{
		services:     services,
		tokenManager: tokenManager,
		filesManager: filesManager,
	}
}

func (h *Handler) Init(cfg *config.Config) *echo.Echo {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStoreWithConfig(
		middleware.RateLimiterMemoryStoreConfig{
			Rate:      cfg.Limiter.RPS,
			Burst:     cfg.Limiter.Burst,
			ExpiresIn: cfg.Limiter.TTL,
		})))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))

	/*docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port)
	if cfg.Environment != config.EnvLocal {
		docs.SwaggerInfo.Host = cfg.HTTP.Host
	}

	if cfg.Environment != config.Prod {
		e.GET("/swagger/*", echoSwagger.WrapHandler)
	}*/

	// Routes
	e.GET("/", healthCheck)
	h.initAPI(e)

	return e
}

func (h *Handler) initAPI(e *echo.Echo) {
	handlerV1 := v1.NewHandler(h.services, h.tokenManager, h.filesManager)
	api := e.Group("/api")
	{
		handlerV1.Init(api)
	}
}

func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Server is up and running",
	})
}
