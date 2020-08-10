package calculator

type DiscountCalculator struct {
	minimumPurchaseAmount int
	discountAmount        int
}

func NewDiscountCalculator(minimumPurchaseAmount int, discountAmount int) *DiscountCalculator {
	// discountRepository database.Discount
	// if minimumPurchaseAmount == 0 {
	// 	return &DiscountCalculator{}, errors.New("minimum purchase amount could not be zero")
	// }
	return &DiscountCalculator{
		minimumPurchaseAmount: minimumPurchaseAmount,
		discountAmount:        discountAmount,
	}
}

func (c *DiscountCalculator) Calculate(purchaseAmount int) int {
	// if purchaseAmount > c.minimumPurchaseAmount {
	// 	multiplier := purchaseAmount / c.minimumPurchaseAmount
	// 	discount := c.discountRepository.FindCurrentDiscount()

	// 	return purchaseAmount - (discount * multiplier)
	// }

	if purchaseAmount > c.minimumPurchaseAmount {
		multiplier := purchaseAmount / c.minimumPurchaseAmount
		return purchaseAmount - (c.discountAmount * multiplier)
	}

	return purchaseAmount
}
