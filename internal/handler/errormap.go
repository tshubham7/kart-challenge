package handler

import (
	"net/http"

	"github.com/tshubham7/kart-challenge/pkg/constants"
)

var statusMap = map[string]int{
	constants.ErrorCodeEntityNotFound:    http.StatusNotFound,
	constants.ErrorCodeInternalError:     http.StatusInternalServerError,
	constants.ErrorCodeBadRequest:        http.StatusBadRequest,
	constants.ErrorCodeInvalidCouponCode: http.StatusUnprocessableEntity,
}
