package user_controller

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sergeyiksanov/help-on-road/internal/services/user_service"
)

type UserController struct {
	userService *user_service.UserService
}

func NewUserController(userService *user_service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Routes() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	r.Post("/sign_up", uc.SignUp)
	r.Post("/sign_in", uc.SignIn)
	r.Get("/me", uc.GetMe)
	r.Post("/update", uc.Update)

	return r
}
