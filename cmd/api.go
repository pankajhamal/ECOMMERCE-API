package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	repo "github.com/pankajhamal/ECOMMERCE-API/internal/adapters/postgresql/sqlc"
	"github.com/pankajhamal/ECOMMERCE-API/internal/products"
)

//mount
func (app *application) mount() http.Handler{
	r := chi.NewRouter()

	// Middleware 
	r.Use(middleware.RequestID) // important for rate limiting
	r.Use(middleware.RealIP) // import for rate limiting and analytics and tracing 
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer) //recover from crashes

	//60 seconds timeout for the request
	r.Use(middleware.Timeout((60 * time.Second)))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello pankaj"))
	})

	// Product API
	productService := products.NewService(repo.New(app.db))
	productHandler := products.NewHandler(productService)
	r.Get("/products", productHandler.ListProducts)
	r.Get("/product/{id}", productHandler.FindProductByID)


	return r
}

//run
func (app *application) run (h http.Handler) error {
	srv := &http.Server{
		Addr: app.config.addr,
		Handler: h,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second *10,
		IdleTimeout: time.Minute,
	}
	log.Printf("Server has started at add %s", app.config.addr)

	return srv.ListenAndServe()

}


type application struct{
	config config
	//logger
	db *pgx.Conn

}


type config struct{
	addr string //port defination
	db dbConfig //database configuration
}

type dbConfig struct{
	dsn string  //user, password
}