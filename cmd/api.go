package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	repo "github.com/nicnuyster/Effective_Mobile_Junior_GO/internal/adapters/postgresql/sqlc"
	"github.com/nicnuyster/Effective_Mobile_Junior_GO/internal/usecase"
)

// mount
func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// серединное
	r.Use(middleware.RequestID) // ratelimiting
	r.Use(middleware.RealIP)    // ratelimiting + analytics +tracing
	r.Use(middleware.Logger)    // log)
	r.Use(middleware.Recoverer) // recover from crashes

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all good"))
	})

	usecaseService := usecase.NewService(repo.New(app.db))
	usecasehandler := usecase.NewHandler(usecaseService)
	r.Get("/listall", usecasehandler.ListAll)

	return r
}

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 20,
		IdleTimeout:  time.Minute * 1,
	}

	log.Printf("Запустили сервер %s", app.config.addr)

	return srv.ListenAndServe()
}

// app
type application struct {
	config config
	// logger
	db *pgx.Conn
}

// cfg
type config struct {
	addr string
	db   dbConfig
}

// db config
type dbConfig struct {
	dsn string
}
