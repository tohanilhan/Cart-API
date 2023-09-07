package models

import (
	"github.com/tohanilhan/Cart-API/vars"
)

type Discount interface {
	ApplyDiscount(order *Order) (float64, string)
}

func (o *Order) ApplyBestDiscounts(discounts []Discount) {

	maxDiscount := 0.0
	var discountReason string

	for _, d := range discounts {
		discountAmount, reason := d.ApplyDiscount(o)
		if discountAmount > maxDiscount {
			maxDiscount = discountAmount
			discountReason = reason
		}
	}

	o.Discount = maxDiscount
	o.DiscountReason = discountReason
	o.TotalPriceWithoutDiscount = o.Cart.TotalPrice
	o.TotalPriceWithDiscount = o.TotalPriceWithoutDiscount - maxDiscount

}

type DiscountOnFourthOrderMoreThanGivenAmount struct {
}

func (d DiscountOnFourthOrderMoreThanGivenAmount) ApplyDiscount(order *Order) (float64, string) {
	discount := 0.0
	reason := "Every fourth order whose total is more than given amount have discount if products have %8 and %18 VAT"

	isFourthOrder := (vars.TotalOrder+1)%4 == 0

	// check if this is the fourth order. Every fourth order whose total is more than given amount may have discount depending on products.
	if order.TotalPriceWithoutDiscount >= vars.AppConfigs.GivenAmount && isFourthOrder {

		// Products whose VAT is %1 don’t have any discount but products whose VAT is %8 and %18 have discount of %10 and %15 respectively.
		for _, item := range order.Cart.Products {

			// if Vat is 8, discount is %10
			if item.Vat == 8 {
				discount += item.Price * float64(item.Quantity) * 0.1

			}

			// if Vat is 18, discount is %15
			if item.Vat == 18 {
				discount += item.Price * float64(item.Quantity) * 0.15
			}
		}

	}

	return discount, reason
}

type DiscountOnSameThirdProducts struct {
}

func (d DiscountOnSameThirdProducts) ApplyDiscount(order *Order) (float64, string) {
	discount := 0.0
	reason := "%8 off because cart contains more than 3 items of the same product"

	// If there are more than 3 items of the same product, then fourth and subsequent ones would have %8 off.
	for _, item := range order.Cart.Products {
		if item.Quantity > 3 {
			discount += item.Price * float64(item.Quantity-3) * 0.08
		}
	}

	return discount, reason
}

type DiscountOnOrderMoreThanGivenAmountInAMonth struct {
}

func (d DiscountOnOrderMoreThanGivenAmountInAMonth) ApplyDiscount(order *Order) (float64, string) {

	discount := 0.0
	reason := "%10 off because customer made purchase which is more than given amount in a month "

	// If the customer made purchase which is more than given amount in a month then all subsequent purchases should have %10 off.
	if order.TotalPriceWithoutDiscount >= vars.AppConfigs.GivenAmount {
		discount = order.TotalPriceWithoutDiscount * 0.1
	}

	return discount, reason
}
