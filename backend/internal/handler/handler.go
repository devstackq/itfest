package handler

import (
	"bimbo/internal/config"
	"bimbo/internal/service"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	services service.Services
	logger   *logrus.Logger
	cfg      *config.Config
}

func NewHandler(srv service.Services, logger *logrus.Logger, cfg *config.Config) *Handler {
	return &Handler{services: srv, logger: logger, cfg: cfg}
}
