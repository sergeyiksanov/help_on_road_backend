package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sergeyiksanov/help-on-road/internal/service_provider"
)

func main() {
	service_provider := service_provider.ServiceProvider{}
	defer func() {
		service_provider.RedisClient().Close()
		close(service_provider.UsersForModerationChannel())
	}()

	r := chi.NewRouter()
	r.Mount("/api/users", service_provider.UserController().Routes())
	r.Mount("/api/help", service_provider.HelpController().Routes())

	go func() {
		if err := service_provider.TgClient().Start(context.Background()); err != nil {
			panic(err)
		}
	}()

	http.ListenAndServe(":8000", r)
}
