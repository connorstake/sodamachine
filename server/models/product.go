package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	AmountAvailable int    `json:"amountAvailable"`
	Cost            int    `json:"cost"`
	ProductName     string `json:"productName"`
	SellerID        uint   `json:"sellerID"`
}

func (p *Product) DecreaseStock(quantity int) error {
	if quantity > p.AmountAvailable {
		return fmt.Errorf("there is not enough product in stock to fulfill this order")
	}
	p.AmountAvailable = p.AmountAvailable - quantity
	return nil
}

func (p *Product) AddStock(quantity int) error {
	p.AmountAvailable = p.AmountAvailable + quantity
	return nil
}

func (p *Product) GetName() string {
	return p.ProductName
}

func (p *Product) ChangeName(newName string) error {
	p.ProductName = newName
	return nil
}
