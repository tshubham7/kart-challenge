package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tshubham7/kart-challenge/internal/handler"
	"github.com/tshubham7/kart-challenge/internal/service"
	"github.com/tshubham7/kart-challenge/internal/service/validator"
	"github.com/tshubham7/kart-challenge/pkg/db"
	"github.com/tshubham7/kart-challenge/pkg/middleware"
)

func AddRoutes(r *mux.Router) {
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.AuthMiddleware)

	gets := api.Methods(http.MethodGet).Subrouter()
	posts := api.Methods(http.MethodPost).Subrouter()
	store := db.NewDBStore()
	orderValidation := validator.NewOrderValidator()
	httpService := service.NewHttpService(store, orderValidation)
	httpHandler := handler.NewHttpHandler(httpService)

	gets.HandleFunc("/product", httpHandler.HandleGetProducts)
	gets.HandleFunc("/product/{productId}", httpHandler.HandleGetProductById)

	posts.HandleFunc("/order", httpHandler.HandleCreateOrder)
}
