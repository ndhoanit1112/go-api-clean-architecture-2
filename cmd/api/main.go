package main

import (
	"github.com/ndhoanit1112/go-api-clean-architecture-2/internal/adapters/http"
	"github.com/ndhoanit1112/go-api-clean-architecture-2/internal/adapters/mysql"
	"github.com/ndhoanit1112/go-api-clean-architecture-2/internal/configs"
)

func main() {
	// Get all configuration
	cfg, err := configs.GetAllConfig()
	if err != nil {
		return
	}

	// Get HTTP server configuration
	httpCfg, err := cfg.GetHTTPConfig()
	if err != nil {
		return
	}

	// Get database connection configuration
	dbCfg, err := cfg.GetDbConfig()
	if err != nil {
		return
	}

	// Initialize DB instance
	dbInstance, err := mysql.InitDb(dbCfg)
	if err != nil {
		return
	}

	// Initialize HTTP server
	http, err := http.InitServer(httpCfg, dbInstance)
	if err != nil {
		return
	}

	// Start HTTP server
	http.Start()
}
