package calculator

import "testing"

func TestDiscountCalculator(t *testing.T) {
	type testCase struct {
		name                  string
		minimumPurchaseAmount int
		discount              int
		purchaseAmount        int
		exepectedAmount       int
	}

	testCases := []testCase{
		{name: "should apply 20", minimumPurchaseAmount: 100, discount: 20, purchaseAmount: 150, exepectedAmount: 130},
		{name: "should apply 40", minimumPurchaseAmount: 100, discount: 20, purchaseAmount: 200, exepectedAmount: 160},
		{name: "should apply 60", minimumPurchaseAmount: 100, discount: 20, purchaseAmount: 350, exepectedAmount: 290},
		{name: "should not apply", minimumPurchaseAmount: 100, discount: 20, purchaseAmount: 50, exepectedAmount: 50},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			calculator, err := NewDiscountCalculator(tc.minimumPurchaseAmount, tc.discount)
			if err != nil {
				// FailNow + log
				t.Fatalf("could not instantiate the calculator %s", err.Error())
			}
			amount := calculator.Calculate(tc.purchaseAmount)

			if amount != tc.exepectedAmount {
				t.Errorf("exepected %v, got %v", tc.exepectedAmount, amount)
			}
		})
	}
}

func TestDiscountCalculatorShouldFailWithZero(t *testing.T) {
	_, err := NewDiscountCalculator(0, 20)
	if err == nil {
		t.Fatalf("should not calcularor with zero %s", err.Error())
	}
}
