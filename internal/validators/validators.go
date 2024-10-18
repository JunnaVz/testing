package validators

import (
	"github.com/google/uuid"
	"lab3/internal/models"
	"net/mail"
	"regexp"
	"time"
)

func ValidName(name string) bool {
	return len(name) > 0
}

func ValidPrice(price float64) bool {
	return price > 0
}

func ValidCategory(category int) bool {
	return category > 0 && category < 9
}

func ValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func ValidAddress(address string) bool {
	return len(address) > 0
}

func ValidPhoneNumber(phoneNumber string) bool {
	re := regexp.MustCompile(`^\+\d{1,3}\d{10}$`)
	return re.MatchString(phoneNumber)
}

func ValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	reLetter := regexp.MustCompile(`[a-zA-Z]`)
	reNumber := regexp.MustCompile(`[0-9]`)

	return reLetter.MatchString(password) && reNumber.MatchString(password)
}

func ValidRole(role int) bool {
	return role > 0 && role < 3
}

func ValidDeadline(deadline time.Time) bool {
	return deadline.After(time.Now())
}

func ValidTasksNumber(tasks []models.OrderedTask) bool {
	return len(tasks) > 0
}

func ValidStatus(status int) bool {
	return status == models.NewOrderStatus || status == models.InProgressOrderStatus || status == models.CompletedOrderStatus || status == models.CancelledOrderStatus
}

func ValidRate(rate int) bool {
	return rate >= 0 && rate <= 5
}

func TaskIsAttachedToOrder(taskID uuid.UUID, tasks []models.Task) bool {
	for _, task := range tasks {
		if task.ID == taskID {
			return true
		}
	}
	return false
}
