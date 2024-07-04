package cmd

import (
	"sync"
	route "tungit/apps/routes"
	config "tungit/configs"
	"tungit/pkg/common/logger"
	recover "tungit/pkg/common/recover"
)

func InitApi() {
	// Logger
	logger.Init()

	// Load configuration
	config := config.LoadConfig()

	// Create a wait group to wait for both grpc & gin server to finish
	var wg sync.WaitGroup
	wg.Add(1)

	// Gin setup router
	go func() {
		defer recover.Goroutine()
		route.LoadRouter(config)
	}()

	// Wait for both servers to finish
	wg.Wait()
}
