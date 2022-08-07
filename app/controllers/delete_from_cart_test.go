package controllers_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/app/controllers"
	"github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/structs"
	"github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/vars"
)

func TestDeleteProductFromBasket(t *testing.T) {

	id1 := uuid.New()
	id2 := uuid.New()

	vars.Cart = append(vars.Cart, structs.Cart{
		ProductId:   id1,
		Quantity:    1,
		Price:       100,
		ProductName: "Product 1",
	})
	vars.Cart = append(vars.Cart, structs.Cart{
		ProductId:   id2,
		Quantity:    1,
		Price:       100,
		ProductName: "Product 2",
	})
	type args struct {
		productId string
		isInCart  bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// Add test cases.
		{
			name: "Delete product from cart",
			args: args{
				productId: id1.String(),
				isInCart:  true,
			},
			want: true,
		},
		{
			name: "Delete product from cart",
			args: args{
				productId: id2.String(),
				isInCart:  true,
			},
			want: true,
		},
		{
			name: "Delete product from cart",
			args: args{
				productId: uuid.New().String(),
				isInCart:  false,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := controllers.DeleteProductFromBasket(tt.args.productId); got != tt.want {
				t.Errorf("DeleteProductFromBasket() = %v, want %v", got, tt.want)
			}
		})
	}
}
