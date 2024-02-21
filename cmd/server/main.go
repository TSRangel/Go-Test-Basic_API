package main

import (
	"github.com/TSRangel/Go-Test-Basic_API/configs"
	"github.com/TSRangel/Go-Test-Basic_API/internal/infra/database"
	"github.com/TSRangel/Go-Test-Basic_API/internal/infra/webserver/handlers"
	"github.com/TSRangel/Go-Test-Basic_API/pkg/tools"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
)

func main() {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := tools.DBConnectionWithDataSource("test.db")
	if err != nil {
		panic(err)
	}
	// err = tools.CreateProductsTable(db)
	// if err != nil {
	// 	panic(err)
	// }
	// err = tools.CreateUsersTable(db)
	// if err != nil {
	// 	panic(err)
	// }
	productHandler := handlers.NewProductHandler(database.NewProductConnection(db))
	userHandler := handlers.NewUserHandler(database.NewUserConnection(db), cfg.TokenAuth, cfg.JWTExpiresIn)

	r := chi.NewRouter()
	
	r.Use(middleware.Logger)
	
	r.Route("/products",func (r chi.Router) {
		r.Use(jwtauth.Verifier(cfg.TokenAuth))
		r.Use(jwtauth.Authenticator)
		
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.ListAllProducts)
		r.Get("/{id}", productHandler.ListProductByID)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})
	
	r.Route("/users", func (r chi.Router) {
		r.Post("/", userHandler.Create)
		r.Get("/login", userHandler.Login)
	})
	
	http.ListenAndServe(":8000", r)
}