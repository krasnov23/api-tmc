package models

import (
)

type InventoryName struct {
	ID         int            `json:"id"`
	Name       string         `json:"name" validate:"required"`
	CategoryId int  		  `json:"categoryId"`
}