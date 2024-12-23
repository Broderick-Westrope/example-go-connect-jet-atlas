package main

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/Broderick-Westrope/example-go-connect-jet-atlas/gen/user/v1/userv1connect"
	"github.com/Broderick-Westrope/example-go-connect-jet-atlas/internal/data"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"

	_ "github.com/lib/pq"
)

const ()

type application struct {
	repo data.Respository
	log  *slog.Logger
}

func main() {
	os.Exit(run())
}

func run() int {
	ctx := context.Background()
	log := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Load environment variables
	addr := "0.0.0.0:" + os.Getenv("PORT")
	dsn := os.Getenv("DSN")

	db, err := setupDatabase(dsn)
	if err != nil {
		log.Error("failed to setup database", slog.Any("err", err))
		return 1
	}
	defer db.Close()

	app := &application{
		repo: data.NewRepository(ctx, db),
		log:  log,
	}

	router := app.setupRouter()
	log.Info("starting server", slog.String("addr", addr))
	err = http.ListenAndServe(
		addr,
		// h2c is used so we can serve HTTP/2 without TLS.
		h2c.NewHandler(router, &http2.Server{}),
	)
	if err != nil {
		log.Error("server error", slog.Any("err", err))
		return 1
	}
	return 0
}

func setupDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("opening: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("pinging: %w", err)
	}
	return db, nil
}

func (app *application) setupRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Mount(userv1connect.NewUserServiceHandler(app, connect.WithInterceptors(
		ServiceVersionInterceptor("UserService", "v1"),
	)))
	return r
}
