package help_controller

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sergeyiksanov/help-on-road/internal/services/help_service"
	"github.com/sergeyiksanov/help-on-road/internal/services/user_service"
)

type HelpController struct {
	helpService *help_service.HelpService
	userService *user_service.UserService
}

func NewHelpController(helpService *help_service.HelpService) *HelpController {
	return &HelpController{
		helpService: helpService,
	}
}

func (hc *HelpController) Routes() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	r.Post("/help", hc.Help)
	r.Get("/get", hc.Get)

	return r
}
