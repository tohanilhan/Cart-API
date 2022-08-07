package controllers_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/app/controllers"
	"github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/structs"
	"github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/vars"
)

var (
	cart1 = []structs.Cart{}
	cart2 = []structs.Cart{}
	cart3 = []structs.Cart{}
)

func TestCalculateDiscount(t *testing.T) {

	cart1 = append(cart1, structs.Cart{
		ProductId:   uuid.New(),
		Quantity:    5,
		Price:       400,
		ProductName: "Product 4",
		Vat:         1,
	})

	cart2 = append(cart2, structs.Cart{
		ProductId:   uuid.New(),
		Quantity:    1,
		Price:       400,
		ProductName: "Product 1",
		Vat:         1,
	})
	cart2 = append(cart2, structs.Cart{
		ProductId:   uuid.New(),
		Quantity:    1,
		Price:       400,
		ProductName: "Product 2",
		Vat:         8,
	})
	cart2 = append(cart2, structs.Cart{
		ProductId:   uuid.New(),
		Quantity:    1,
		Price:       400,
		ProductName: "Product 3",
		Vat:         18,
	})
	// cart2 = append(cart2, structs.Cart{
	// 	ProductId:   uuid.New(),
	// 	Quantity:    1,
	// 	Price:       400,
	// 	ProductName: "Product 3",
	// 	Vat:         18,
	// })

	cart3 = append(cart3, structs.Cart{
		ProductId:   uuid.New(),
		Quantity:    1,
		Price:       100,
		ProductName: "Product 2",
	})

	cart3 = append(cart3, structs.Cart{
		ProductId:   uuid.New(),
		Quantity:    1,
		Price:       200,
		ProductName: "Product 3",
	})

	type args struct {
		totalOrder       int16
		orderTotalAmount float64
		givenAmount      float64
		cart             []structs.Cart
		month            string
	}

	tests := []struct {
		name  string
		args  args
		want  float64
		want1 string
		want2 float64
	}{
		// Add test cases.

		// Discount discount on same third products and fourth order but not more than given amount in a month
		{
			name:  "Test case 1",
			args:  args{totalOrder: 3, givenAmount: 4000, cart: cart1, month: "February", orderTotalAmount: 2000},
			want:  2000,
			want1: "If there are more than 3 items of the same product, then fourth and subsequent ones would have %8 off.",
			want2: 1936,
		},

		// Discount discount on same third products and fourth order but also more purchases made more than given amount in a month
		{
			name:  "Test case 2",
			args:  args{totalOrder: 3, givenAmount: 4000, cart: cart1, month: "August", orderTotalAmount: 4001},
			want:  2000,
			want1: "If the customer made purchase which is more than given amount in a month then all subsequent purchases should have %10 off.",
			want2: 1800,
		},

		// Discount on same third products because its not fourth order but purchases made more than given amount in a month
		{
			name:  "Test case 3",
			args:  args{totalOrder: 2, givenAmount: 3000, cart: cart1, month: "August", orderTotalAmount: 3000},
			want:  2000,
			want1: "If the customer made purchase which is more than given amount in a month then all subsequent purchases should have %10 off.",
			want2: 1800,
		},

		// Discount on fourth order more than given amount
		{
			name:  "Test case 4",
			args:  args{totalOrder: 3, givenAmount: 1200, cart: cart2, month: "February", orderTotalAmount: 0},
			want:  1200,
			want1: "Every fourth order whose total is more than given amount may have discount depending on products. Products whose VAT is %1 don’t have any discount but products whose VAT is %8 and %18 have discount of %10 and %15 respectively.",
			want2: 1100,
		},

		// Discount on fourth order less than given amount in a month but purchases made more than given amount in a month
		{
			name:  "Test case 5",
			args:  args{totalOrder: 3, givenAmount: 3000, cart: cart2, month: "August", orderTotalAmount: 3000},
			want:  1200,
			want1: "If the customer made purchase which is more than given amount in a month then all subsequent purchases should have %10 off.",
			want2: 1080,
		},

		// No discount because not fourth order or purchases made less than given amount in a month
		{
			name:  "Test case 6",
			args:  args{totalOrder: 2, givenAmount: 3000, cart: cart2, month: "February", orderTotalAmount: 0},
			want:  1200,
			want1: "No discount",
			want2: 1200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vars.TotalOrder = tt.args.totalOrder
			vars.GivenAmount = tt.args.givenAmount
			vars.Cart = tt.args.cart
			controllers.Month = tt.args.month
			controllers.OrderTotalAmount = tt.args.orderTotalAmount

			totalPriceWithoutDiscount, reason, totalPriceWithDiscount := controllers.CalculateDiscount()

			if totalPriceWithoutDiscount != tt.want {
				t.Errorf("CalculateDiscount() totalPriceWithoutDiscount = %v, want %v", totalPriceWithoutDiscount, tt.want)
			}
			if reason != tt.want1 {
				t.Errorf("CalculateDiscount() reason = %v, want %v", reason, tt.want1)
			}
			if totalPriceWithDiscount != tt.want2 {
				t.Errorf("CalculateDiscount() totalPriceWithDiscount = %v, want %v", totalPriceWithDiscount, tt.want2)
			}
		})
	}
}
