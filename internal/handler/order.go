package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/tshubham7/kart-challenge/pkg/model"
	"github.com/tshubham7/kart-challenge/pkg/utility"
	"github.com/tshubham7/kart-challenge/pkg/webservice"
)

func (h *httpHandler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	ctx, log := utility.NewLogger(r.Context())
	log.Info("got request for create order")

	body, err := webservice.ReadFullBody(r)
	if err != nil {
		webservice.RespondWithStatus(w, http.StatusBadRequest, nil)
		return
	}

	var reqData CreateOrderRequest
	if err := json.Unmarshal(body, &reqData); err != nil {
		webservice.RespondWithStatus(w, http.StatusBadRequest, nil)
		return
	}

	validate := validator.New()
	if err := validate.Struct(reqData); err != nil {
		webservice.RespondWithStatus(w, http.StatusBadRequest, nil)
		return
	}

	order, err := h.svc.CreateOrder(ctx, reqData.CouponCode, adaptToOrderItems(reqData.Items))
	if err != nil {
		webservice.RespondWithStatus(w, statusMap[err.Error()], nil)
		return
	}
	webservice.RespondWithStatus(w, http.StatusOK, order)
}

func adaptToOrderItems(reqItems []OrderItem) []model.OrderItem {
	items := make([]model.OrderItem, len(reqItems))
	for i, ri := range reqItems {
		items[i] = model.OrderItem{
			ProductId: ri.ProductId,
			Quantity:  ri.Quantity,
		}
	}
	return items
}
