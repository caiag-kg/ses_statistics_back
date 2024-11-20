package apps

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"ses_back/internal/config"
	"ses_back/internal/service"
)

// Main function to initialize and run the application.
func App() {
	// Read configuration from file.
	config.ReadConfig()
	var svrCFG = config.Config.Server
	// Setting release mode
	gin.SetMode(gin.ReleaseMode)
	// Initialize the Gin server.
	server := gin.Default()

	// Define API services.
	// Services:
	go server.GET(
		"/events",
		service.GetEvents,
	)

	go server.POST(
		"/events",
		service.GetEvByFilter,
	)

	//---------------------------------------------------------------//
	// Construct server address from configuration.
	var addr = fmt.Sprintf("%s:%s", svrCFG.Host, svrCFG.Port)
	// Print server address.
	fmt.Printf("Server run on https://%s\n", addr)

	// Run the server and handle any errors.
	err := server.Run(addr)
	if err != nil {
		// Log fatal error if server fails to run.
		log.Fatal("Server run error: ", err)
	}
}