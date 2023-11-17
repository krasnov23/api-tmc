package models

import "time"

type InventoryTransfer struct {
	ID           int       `json:"id"`
	SenderID     int       `json:"sender_id"`
	ReceiverID   int       `json:"reciever_id"`
	Ident        string    `json:"ident"`
	TransferDate time.Time `json:"transfer_date"`
	Status       string    `json:"status"`
}