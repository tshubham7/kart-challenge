package handler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tshubham7/kart-challenge/internal/service"
	"github.com/tshubham7/kart-challenge/pkg/utility"
	"github.com/tshubham7/kart-challenge/pkg/webservice"
)

type httpHandler struct {
	svc service.HttpService
}

func NewHttpHandler(svc service.HttpService) *httpHandler {
	return &httpHandler{svc}
}

func (h *httpHandler) HandleGetProducts(w http.ResponseWriter, r *http.Request) {
	ctx, log := utility.NewLogger(r.Context())
	log.Info("got request for get products")

	products, err := h.svc.GetProducts(ctx)
	if err != nil {
		log.Errorf("operation failed: %v", err)
		webservice.RespondWithStatus(w, statusMap[err.Error()], nil)
		return
	}

	webservice.RespondWithStatus(w, http.StatusOK, products)
}

func (h *httpHandler) HandleGetProductById(w http.ResponseWriter, r *http.Request) {
	ctx, log := utility.NewLogger(r.Context())
	log.Info("got request for get product detail")

	pathVariables := mux.Vars(r)
	id := pathVariables["productId"]
	idInt, err := strconv.Atoi(id)
	if err != nil || idInt < 1 {
		webservice.RespondWithStatus(w, http.StatusBadRequest, nil)
		return
	}

	product, err := h.svc.GetProductById(ctx, id)
	if err != nil {
		log.Errorf("operation failed: %v", err)
		webservice.RespondWithStatus(w, statusMap[err.Error()], nil)
		return
	}
	webservice.RespondWithStatus(w, http.StatusOK, product)
}
