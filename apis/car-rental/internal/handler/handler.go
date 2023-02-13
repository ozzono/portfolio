package handler

import (
	"car-rental/controller"
	"car-rental/internal/repository"

	"go.uber.org/zap"
)

type Handler struct {
	Client repository.Client
	log    *zap.SugaredLogger
	ctrl   *controller.Controller
}

func NewHandler(c repository.Client, l *zap.SugaredLogger, debug bool) *Handler {
	return &Handler{c, l, controller.NewController(l, c)}
}
