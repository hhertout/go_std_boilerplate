package main

import (
	"log"
	"std_go_boilerplate/internal/server"
)

func main() {
	server := server.NewServer()

	log.Printf("🚀 Server running on port %s\n", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		panic("cannot start server")
	}
}
