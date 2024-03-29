package main

import (
	"github.com/TSRangel/Go-Test-Basic_API/configs"
	"github.com/TSRangel/Go-Test-Basic_API/internal/infra/database"
	"github.com/TSRangel/Go-Test-Basic_API/internal/infra/webserver/handlers"
	"github.com/TSRangel/Go-Test-Basic_API/pkg/tools"
	_ "github.com/TSRangel/Go-Test-Basic_API/docs"

	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title 						Go API Example
// @version 					1.0
// @description 				Product API with JWT authentication
// @termsOfService 				http://swagger.io/terms/

// @contact.name 				Thiago da Silva Rangel
// @contact.email 				exemplo@gmail.com

// @license.name 				Test API license

// @host 						localhost:8000
// @BasePath 					/
// @securityDefinitions.apikey 	ApiKeyAuth
// @in 							header
// @name 						Authorization
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
	userHandler := handlers.NewUserHandler(database.NewUserConnection(db))

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(cfg.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.ListAllProducts)
		r.Get("/{id}", productHandler.ListProductByID)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Route("/users", func(r chi.Router) {
		r.Use(middleware.WithValue("token", cfg.TokenAuth))
		r.Use(middleware.WithValue("expiresIn", cfg.JWTExpiresIn))
		r.Post("/", userHandler.Create)
		r.Post("/login", userHandler.Login)
	})

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	http.ListenAndServe(":8000", r)
}