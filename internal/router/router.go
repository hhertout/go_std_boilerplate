package router

import (
	"net/http"
	"std_go_boilerplate/internal/controller"
)

func ServeRoutes() http.Handler {
	router := http.NewServeMux()
	c := controller.NewBaseController()

	router.HandleFunc("/health", c.HealthCheck)

	Group(router, "/api/v1", func(v1 *http.ServeMux) {
		// other routes here
	}) // can set middleware here

	return router
}
