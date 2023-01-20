package main

import (
	"web-service-api/routes"
)

func main() {
	router := routes.NewRouter()

	router.SetMiddleware()
	router.SetProxy()
	router.SetHealthChecker()
	router.NewTaskRouter()

	router.Serve()
}
