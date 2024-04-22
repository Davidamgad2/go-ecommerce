package models

import (
	"gorm.io/gorm"
)
type Cart struct {
	gorm.Model
	cart_user User
	Items     []Product
}

func (c Cart) NewCart() *Cart {
	return &Cart{}
}

func (c *Cart) AddItem(item Product) {
	c.Items = append(c.Items, item)
}

func (c *Cart) RemoveItem(item Product) {
	for i, v := range c.Items {
		if v == item {
			c.Items = append(c.Items[:i], c.Items[i+1:]...)
			break
		}
	}
}

func (c *Cart) GetItems() []Product {
	return c.Items
}
