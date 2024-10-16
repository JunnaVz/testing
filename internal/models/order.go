package models

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	ID           uuid.UUID `json:"id"`
	WorkerID     uuid.UUID `json:"worker_id"`
	UserID       uuid.UUID `json:"user_id"`
	Status       int       `json:"status"`
	Address      string    `json:"address"`
	CreationDate time.Time `json:"creation_date"`
	Deadline     time.Time `json:"deadline"`
	Rate         int       `json:"rate"`
}

const NoStatus = 0
const NewOrderStatus = 1
const InProgressOrderStatus = 2
const CompletedOrderStatus = 3
const CancelledOrderStatus = 4

var OrderStatuses = map[int]string{
	NoStatus:              "Не определен",
	NewOrderStatus:        "Новый",
	InProgressOrderStatus: "В процессе",
	CompletedOrderStatus:  "Завершен",
	CancelledOrderStatus:  "Отменен",
}
