package calculator

import "testing"

func TestDiscountCalculator(t *testing.T) {
	type testCase struct {
		minimumPurchaseAmount int
		discount              int
		purchaseAmount        int
		exepectedAmount       int
	}

	testCases := []testCase{
		{minimumPurchaseAmount: 100, discount: 20, purchaseAmount: 150, exepectedAmount: 130},
		{minimumPurchaseAmount: 100, discount: 20, purchaseAmount: 200, exepectedAmount: 160},
		{minimumPurchaseAmount: 100, discount: 20, purchaseAmount: 350, exepectedAmount: 290},
		{minimumPurchaseAmount: 100, discount: 20, purchaseAmount: 50, exepectedAmount: 50},
	}

	for _, tc := range testCases {
		calculator := NewDiscountCalculator(tc.minimumPurchaseAmount, tc.discount)
		amount := calculator.Calculate(tc.purchaseAmount)

		if amount != tc.exepectedAmount {
			t.Errorf("exepected %v, got %v", tc.exepectedAmount, amount)
		}
	}
}
