package models

type OrderedTask struct {
	Task     *Task `json:"task"`
	Quantity int   `json:"quantity"`
}
