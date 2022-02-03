package handlers

import (
	"github.com/andhikasamudra/auth-rest-fiber/config"
	"github.com/uptrace/bun"
)

type Handler struct {
	DB     *bun.DB
	Config *config.Config
}

func New(db *bun.DB, cfg *config.Config) *Handler {
	return &Handler{
		DB:     db,
		Config: cfg,
	}
}
