package main

import (
	"os"

	router "github.com/Hirogava/swifty-gasprom/backend/internal/transport/http"
	postgres "github.com/Hirogava/swifty-gasprom/backend/internal/repository/postgres"
)

func main() {
	manager := postgres.NewManager("postgres", os.Getenv("DB_CONNECT_STRING"))

	manager.Migrate()

	r := router.CreateRouter(manager)

	r.Run(os.Getenv("SERVER_PORT"))
}