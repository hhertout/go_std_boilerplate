package router

import (
	"net/http"
	"std_go_boilerplate/internal/controller"
	"std_go_boilerplate/internal/middleware"
)

func ServeRoutes() http.Handler {
	router := http.NewServeMux()
	c := controller.NewBaseController()

	router.HandleFunc("/health", c.HealthCheck)

	Group(router, "/api/v1", func(v1 *http.ServeMux) {
		// other routes here
	}, middleware.ApiKeyMiddleware) // can set middleware here

	return router
}
