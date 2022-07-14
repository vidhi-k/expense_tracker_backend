package api

import (
	"github.com/vidhi-k/expense_tracker_backend/cmd/migrator"
	"github.com/vidhi-k/expense_tracker_backend/pkg/user"
	"github.com/vidhi-k/expense_tracker_backend/pkg/user/repository"
	"github.com/vidhi-k/expense_tracker_backend/tranport"
	"github.com/vidhi-k/expense_tracker_backend/utl/config"
	"github.com/vidhi-k/expense_tracker_backend/utl/server"
	"github.com/vidhi-k/expense_tracker_backend/utl/storage"
)

func Start() {
	cfg := config.LoadConfig("./utl/config/config.yaml")

	db := storage.GetPostgresDB(cfg.DB.Url)

	migrator.Migrate(db)

	userService := user.InitService(db, repository.NewPostgresRepo())

	ech := server.InitEcho()

	v1Group := ech.Group("/api/v1")

	tranport.InitHTTPUserHandlers(userService, v1Group)

	server.StartServer(ech, cfg.Server.Port)
}
