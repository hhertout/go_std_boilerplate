package main

import (
	"os"
	"std_go_boilerplate/internal/router"
	"std_go_boilerplate/internal/server"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	if os.Getenv("GO_ENV") == "development" {
		logger, _ = zap.NewDevelopment()
	}
	defer logger.Sync()

	if os.Getenv("GO_ENV") == "development" {
		logger.Sugar().Info("⚠️ Caution : The server will be running under development mode 🔨🔨")
	}

	ctx := server.Context{
		Logger: logger,
	}

	r := router.ServeRoutes(ctx)
	server := server.NewServer(r, ctx)

	logger.Sugar().Infof("🚀 Server running on port %s\n", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		logger.Sugar().Fatalf("Error starting server: %v", err)
	}
}
