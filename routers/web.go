package routers

import (
	"gitminhq/gitmin/controllers"
	"github.com/ansrivas/fiberprometheus/v2"
	"gitminhq/gitmin/app"
)

// RegisterRoutes registers all the routes associated with this application.
func RegisterRoutes(application *app.App) {

	ctrl := controllers.New()

	// This here will appear as a label, one can also use
	// fiberprometheus.NewWith(servicename, namespace, subsystem )
	prometheus := fiberprometheus.New("template-app")
	prometheus.RegisterAt(application.Server, "/metrics")
	application.Server.Use(prometheus.Middleware)

	application.Server.Get("/", ctrl.GetIndexHTML)
	application.Server.Get("/healthz", ctrl.GetHandlerHealthz)

	apiGroup := application.Server.Group("/api")
	{
		// Note we are not passing prepareRoutes as we prefixed the proxy earlier
		apiGroup.Get("/index.html", ctrl.GetIndexHTML)
	}
}
