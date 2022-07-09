package api

import (
	"github.com/vidhi-k/expense_tracker_backend/utl/config"
	"github.com/vidhi-k/expense_tracker_backend/utl/storage"
)

func Start() {
	cfg := config.LoadConfig("./utl/config/config.yaml")

	_ = storage.GetPostgresDB(cfg.DB.Url)
}
