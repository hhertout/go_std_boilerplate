package main

import (
	"log"
	"os"
	"std_go_boilerplate/internal/server"
)

func main() {
	server := server.NewServer()

	if os.Getenv("GO_ENV") == "development" {
		log.Println("⚠️ Caution : The server will be running under development mode 🔨🔨")
	}

	log.Printf("🚀 Server running on port %s\n", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		panic("cannot start server")
	}
}
