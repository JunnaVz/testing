package models

import "github.com/google/uuid"

type Worker struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Surname     string    `json:"surname"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phoneNumber"`
	Email       string    `json:"email"`
	Role        int       `json:"role"`
	Password    string    `json:"password"`
}

const ManagerRole = 1
const MasterRole = 2

var WorkerRole = map[int]string{
	ManagerRole: "Менеджер",
	MasterRole:  "Мастер",
}

func (w Worker) DisplayRole() string {
	return WorkerRole[w.Role]
}

func (w Worker) FullName() string {
	return w.Name + " " + w.Surname
}
