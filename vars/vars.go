package vars

import (
	"github.com/google/uuid"
	"github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/structs"
)

var (
	Cart                                       = make([]structs.Cart, 0, 100) // object to store cart
	CartResponse                               structs.CartResponse
	Product                                    structs.Product
	TotalOrder                                 int16
	Order                                      structs.Order
	DiscountOnFourthOrderMoreThanGivenAmount   float64
	DiscountOnSameThirdProducts                float64
	DiscountOnOrderMoreThanGivenAmountInAMonth float64
	FinalDiscount                              float64
	GivenAmount                                float64
	UserId                                     uuid.UUID
)
