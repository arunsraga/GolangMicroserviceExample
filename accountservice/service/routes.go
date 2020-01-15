package service

import (
	. "github.com/arunsraga/goblog/common/router"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Initialize our routes
var routes = Routes{
	Route{
		"GetAccount", // Name
		"GET",        // HTTP method
		"/accounts/{accountId}", // Route pattern
		GetAccount,
		true,
	},
	Route{
		"Prometheus",
		"GET",
		"/metrics",
		promhttp.Handler().ServeHTTP,
		false,
	},
}