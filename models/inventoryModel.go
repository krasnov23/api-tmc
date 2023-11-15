package models

import "time"

type Inventory struct {
	ID         int       `json:"id"`
	AccountID  int       `json:"account_id"`
	StateID    int       `json:"state_id"`
	Ident      string    `json:"ident"`
	DatePay    string    `json:"date_pay"` // В Go нет типа данных "DATE", поэтому используем здесь строку
	DateCreate time.Time `json:"date_create"`
	CategoryID int       `json:"category_id"`
	NameID     int       `json:"name_id"`
}