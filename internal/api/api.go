package api

import (
	"github.com/vidhi-k/expense_tracker_backend/internal/migrator"
	"github.com/vidhi-k/expense_tracker_backend/pkg/auth"
	"github.com/vidhi-k/expense_tracker_backend/pkg/expenses"
	er "github.com/vidhi-k/expense_tracker_backend/pkg/expenses/repository"
	"github.com/vidhi-k/expense_tracker_backend/pkg/user"
	ur "github.com/vidhi-k/expense_tracker_backend/pkg/user/repository"
	"github.com/vidhi-k/expense_tracker_backend/tranport"
	"github.com/vidhi-k/expense_tracker_backend/utl/config"
	"github.com/vidhi-k/expense_tracker_backend/utl/server"
	"github.com/vidhi-k/expense_tracker_backend/utl/storage"
)

func Start() {
	cfg := config.LoadConfig("./utl/config/config.yaml")

	db := storage.GetPostgresDB(cfg.DB.Url)

	migrator.Migrate(db)

	userService := user.InitService(db, ur.NewPostgresRepo())

	authService := auth.InitService(userService)

	expenseService := expenses.InitService(db, er.NewPostgresRepo(), userService)

	ech := server.InitEcho()

	v1Group := ech.Group("/api/v1")

	tranport.InitAuthHTTPHandlers(authService, v1Group)

	tranport.InitUserHTTPHandlers(userService, v1Group)

	tranport.InitExpenseHTTPHandlers(expenseService, v1Group)

	server.StartServer(ech, cfg.Server.Port)
}
