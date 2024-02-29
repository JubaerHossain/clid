package routes

import (
	"net/http"

	categoryRoute "github.com/JubaerHossain/golang-ddd/internal/category/infrastructure/transport/http"
	"github.com/JubaerHossain/golang-ddd/internal/core/health"
	"github.com/JubaerHossain/golang-ddd/internal/core/middleware"
	"github.com/JubaerHossain/golang-ddd/internal/core/monitor"
	"github.com/JubaerHossain/golang-ddd/internal/core/pkg/utils"
	tableRoute "github.com/JubaerHossain/golang-ddd/internal/table/infrastructure/transport/http"
	userRoute "github.com/JubaerHossain/golang-ddd/internal/user/infrastructure/transport/http"
)

// SetupRoutes initializes and returns the HTTP router with all routes.
func SetupRoutes() *http.ServeMux {
	router := http.NewServeMux()

	// Register health check endpoint
	router.Handle("GET /health", middleware.LoggingMiddleware(http.HandlerFunc(health.HealthCheckHandler())))

	// Register monitoring endpoint
	router.Handle("GET /metrics", monitor.MetricsHandler())

	// swagger endpoint
	router.Handle("GET /swagger", http.StripPrefix("/swagger/", http.FileServer(http.Dir("swagger"))))

	// register user routes
	userRoute.SetupUserRoutes(router)

	// register category routes
	categoryRoute.SetupCategoryRoutes(router)

	// register table routes
	tableRoute.SetupTableRoutes(router)

	// Register a welcome message
	router.Handle("/", middleware.LimiterMiddleware(middleware.LoggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.WriteJSONResponse(w, http.StatusOK, map[string]interface{}{"message": "Welcome to the API"})
	}))))

	return router
}
