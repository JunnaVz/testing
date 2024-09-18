package utils

import (
	"lab3/internal/models"
)

func DisplayStatus(statusNum int) string {
	return models.OrderStatuses[statusNum]
}
