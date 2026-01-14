package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sergeyiksanov/help-on-road/internal/service_provider"
)

func RecoverWithAlert(alert func(string)) func() {
	return func() {
		if r := recover(); r != nil {
			alert(fmt.Sprintf("PANIC: %v", r))
		}
	}
}

func main() {
	service_provider := service_provider.ServiceProvider{}
	defer func() {
		service_provider.RedisClient().Close()
		close(service_provider.UsersForModerationChannel())
	}()

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Mount("/api/users", service_provider.UserController().Routes())
	r.Mount("/api/help", service_provider.HelpController().Routes())

	go func() {
		defer RecoverWithAlert(service_provider.TgClient().Alert)

		if err := service_provider.TgClient().Start(context.Background()); err != nil {
			panic(err)
		}
	}()

	log.Fatal(http.ListenAndServe(":8000", r))
}
