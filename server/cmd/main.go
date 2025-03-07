package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/0-s0g0/TEKUTEKU/server/internal/di"
	env "github.com/0-s0g0/TEKUTEKU/server/pkg/config"
	"github.com/0-s0g0/TEKUTEKU/server/pkg/handler"
	"github.com/0-s0g0/TEKUTEKU/server/pkg/middleware"
	"github.com/rs/cors"
)

func main() {
	env.LoadEnv()
	h := di.InitHandler()
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
		w.WriteHeader(http.StatusOK)
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world")
		w.WriteHeader(http.StatusOK)
	})

	mux.Handle("POST /signup", handler.AppHandler(h.UserHandler.SignUp()))
	mux.Handle("POST /signin", handler.AppHandler(h.UserHandler.SignIn()))

	mux.Handle("GET /messages", handler.AppHandler(h.MessageHandler.GET()))
	mux.Handle("POST /messages", handler.AppHandler(h.MessageHandler.POST()))

	mux.Handle("POST /reply", handler.AppHandler(h.MessageHandler.ReplyPOST()))

	mux.Handle("POST /likes", handler.AppHandler(h.LikeHandler.POST()))

	c := cors.AllowAll()
	handler := middleware.Chain(mux, c.Handler, middleware.Context, middleware.Logger, middleware.Recover)

	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			slog.Error("server error", "error", err)
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("server error", "error", err.Error())
	}
}
