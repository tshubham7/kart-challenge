package validator

import "context"

type OrderValidator interface {
	ValidateCouponCode(ctx context.Context, code string) error
}
