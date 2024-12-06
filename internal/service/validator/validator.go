package validator

import (
	"context"
	"encoding/json"
	"errors"
	"os"

	"github.com/tshubham7/kart-challenge/pkg/constants"
	"github.com/tshubham7/kart-challenge/pkg/utility"
)

type orderValidator struct{}

func NewOrderValidator() OrderValidator {
	return &orderValidator{}
}

func (s *orderValidator) ValidateCouponCode(_ context.Context, code string) error {
	if code == "" {
		return nil
	}

	if len(code) < 8 || len(code) > 10 {
		return errors.New(constants.ErrorCodeInvalidCouponCode)
	}

	return s.validateCouponCodeFromFile(code)
}

func (s *orderValidator) validateCouponCodeFromFile(code string) error {
	files := []string{constants.CouponFile1, constants.CouponFile2, constants.CouponFile3}
	var presenceCount int
	for _, file := range files {
		codes, err := s.readCodesFromFile(file)
		if err != nil {
			return err
		}
		if utility.Contains(codes, code) {
			presenceCount++
		}
	}
	if presenceCount < 2 {
		return errors.New(constants.ErrorCodeInvalidCouponCode)
	}
	return nil
}

func (s *orderValidator) readCodesFromFile(file string) ([]string, error) {
	b, err := os.ReadFile(constants.BaseDirectoryData + file)
	if err != nil {
		return nil, err
	}

	var codes []string
	err = json.Unmarshal(b, &codes)
	if err != nil {
		return nil, err
	}

	return codes, err
}
