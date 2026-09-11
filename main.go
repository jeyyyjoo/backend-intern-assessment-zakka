package main

import (
	"errors"
	"fmt"
)

type CartItem struct {
	ProductID string
	Name      string
	Price     float64
	Quantity  int
}

type Voucher struct {
	Code            string
	DiscountPercent float64
	MaxDiscount     float64
	MinPurchase     float64
}

func CalculateFinalPrice(items []CartItem, voucher *Voucher) (subtotal float64, discount float64, total float64, err error) {
	if len(items) == 0 {
		return 0, 0, 0, errors.New("cart cannot be empty")
	}

	for _, item := range items {
		if item.Quantity <= 0 || item.Price < 0 {
			return 0, 0, 0, errors.New("invalid item price or quantity")
		}

		subtotal += item.Price * float64(item.Quantity)
	}

	if voucher == nil {
		return subtotal, 0, subtotal, nil
	}

	if subtotal < voucher.MinPurchase {
		return subtotal, 0, subtotal, nil
	}

	discount = subtotal * (voucher.DiscountPercent / 100)

	if discount > voucher.MaxDiscount {
		discount = voucher.MaxDiscount
	}

	total = subtotal - discount

	return subtotal, discount, total, nil
}

func main() {
	items := []CartItem{
		{
			ProductID: "P001",
			Name:      "Keyboard",
			Price:     250000,
			Quantity:  2,
		},
		{
			ProductID: "P002",
			Name:      "Mouse",
			Price:     150000,
			Quantity:  1,
		},
	}

	voucher := &Voucher{
		Code:            "FLASH10",
		DiscountPercent: 10,
		MaxDiscount:     500000,
		MinPurchase:     500000,
	}

	subtotal, discount, total, err := CalculateFinalPrice(items, voucher)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Subtotal: %.2f\n", subtotal)
	fmt.Printf("Discount: %.2f\n", discount)
	fmt.Printf("Total: %.2f\n", total)
}
