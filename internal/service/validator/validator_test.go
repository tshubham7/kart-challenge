package validator

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tshubham7/kart-challenge/pkg/constants"
)

func TestUnit_ValidateCouponCode(t *testing.T) {
	codeValidator := NewOrderValidator()
	ctx := context.Background()

	err := os.MkdirAll(constants.BaseDirectoryData, os.ModePerm)
	assert.Nil(t, err)
	os.WriteFile(constants.BaseDirectoryData+constants.CouponFile1, []byte(`["HAPPYHRS"]`), os.ModePerm)
	os.WriteFile(constants.BaseDirectoryData+constants.CouponFile2, []byte(`["HAPPYHRS"]`), os.ModePerm)
	os.WriteFile(constants.BaseDirectoryData+constants.CouponFile3, []byte(`["COUPON_CODE2"]`), os.ModePerm)
	defer os.RemoveAll(constants.BaseDirectoryData)

	t.Run("code not present in any file", func(t *testing.T) {
		err := codeValidator.ValidateCouponCode(ctx, "InvalidCode")
		assert.NotNil(t, err)
	})

	t.Run("code present in only 1 file", func(t *testing.T) {
		err := codeValidator.ValidateCouponCode(ctx, "COUPON_CODE2")
		assert.NotNil(t, err)
	})

	t.Run("code length is not 8-10", func(t *testing.T) {
		err := codeValidator.ValidateCouponCode(ctx, "HAPPY")
		assert.NotNil(t, err)
	})

	t.Run("empty coupon code", func(t *testing.T) {
		err := codeValidator.ValidateCouponCode(ctx, "")
		assert.Nil(t, err)
	})

	t.Run("code present in only 2 files", func(t *testing.T) {
		err := codeValidator.ValidateCouponCode(ctx, "HAPPYHRS")
		assert.Nil(t, err)
	})

}
