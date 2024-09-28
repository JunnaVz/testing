package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type Order struct {
	ID     string   `json:"id"`
	Status string   `json:"status"`
	Tasks  []string `json:"tasks"`
}

type Rating struct {
	OrderID string `json:"order_id"`
	Score   int    `json:"score"`
}

type Task struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Тест на создание и удаление задачи
func TestE2E_CreateAndDeleteTask(t *testing.T) {
	// Создаем тестовый HTTP-сервер
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Роутинг для задач
		switch r.Method {
		case http.MethodPost:
			// Логика создания задачи
			var task Task
			json.NewDecoder(r.Body).Decode(&task)
			task.ID = "1"
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(task)
		case http.MethodDelete:
			// Логика удаления задачи
			w.WriteHeader(http.StatusNoContent)
		}
	}))
	defer ts.Close()

	// 1. Создание задачи
	task := Task{Name: "Test Task"}
	taskData, _ := json.Marshal(task)
	resp, err := http.Post(ts.URL, "application/json", bytes.NewBuffer(taskData))
	if err != nil || resp.StatusCode != http.StatusCreated {
		t.Fatalf("Expected task to be created, got error: %v", err)
	}

	// Чтение ответа
	var createdTask Task
	json.NewDecoder(resp.Body).Decode(&createdTask)
	if createdTask.Name != "Test Task" {
		t.Fatalf("Expected task name to be 'Test Task', got %v", createdTask.Name)
	}

	// 2. Удаление задачи
	req, _ := http.NewRequest(http.MethodDelete, ts.URL+"/1", nil)
	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil || resp.StatusCode != http.StatusNoContent {
		t.Fatalf("Expected task to be deleted, got error: %v", err)
	}
}

func TestE2E_OrderProcess(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/login":
			var user User
			json.NewDecoder(r.Body).Decode(&user)

			if user.Username == "admin" {
				user.ID = "1"
				user.Role = "admin"
			} else if user.Username == "worker" {
				user.ID = "2"
				user.Role = "worker"
			} else {
				user.ID = "3"
				user.Role = "user"
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(user)

		case "/orders":
			if r.Method == http.MethodPost {
				var order Order
				json.NewDecoder(r.Body).Decode(&order)
				order.ID = "order1"
				order.Status = "created"

				w.WriteHeader(http.StatusCreated)
				json.NewEncoder(w).Encode(order)
			} else if r.Method == http.MethodPut {
				var order Order
				json.NewDecoder(r.Body).Decode(&order)

				if order.Status == "assigned" {
					order.Status = "assigned"
				} else if order.Status == "assigned" {
					order.Status = "completed"
				}

				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(order)
			}

		case "/rate-order":
			var rating Rating
			json.NewDecoder(r.Body).Decode(&rating)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(rating)

		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	// 1. Логин пользователя
	user := login(t, ts, "user")
	if user.Role != "user" {
		t.Fatalf("Expected role 'user', got %s", user.Role)
	}

	// 2. Создание нового заказа пользователем
	order := createOrder(t, ts, []string{"Task 1", "Task 2"})
	if order.Status != "created" {
		t.Fatalf("Expected order status 'created', got %s", order.Status)
	}

	// 3. Логин администратора
	admin := login(t, ts, "admin")
	if admin.Role != "admin" {
		t.Fatalf("Expected role 'admin', got %s", admin.Role)
	}

	// 4. Назначение заказа воркеру админом
	order.Status = "assigned"
	updateOrderStatus(t, ts, order)
	if order.Status != "assigned" {
		t.Fatalf("Expected order status 'assigned', got %s", order.Status)
	}

	// 5. Логин воркера
	worker := login(t, ts, "worker")
	if worker.Role != "worker" {
		t.Fatalf("Expected role 'worker', got %s", worker.Role)
	}

	// 6. Выполнение заказа воркером
	order.Status = "completed"
	updateOrderStatus(t, ts, order)
	if order.Status != "completed" {
		t.Fatalf("Expected order status 'completed', got %s", order.Status)
	}

	// 7. Оценка заказа пользователем
	rateOrder(t, ts, order.ID, 5)
}

func login(t *testing.T, ts *httptest.Server, username string) User {
	user := User{Username: username}
	body, _ := json.Marshal(user)

	resp, err := http.Post(ts.URL+"/login", "application/json", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Error during login: %v", err)
	}
	defer resp.Body.Close()

	var loggedInUser User
	json.NewDecoder(resp.Body).Decode(&loggedInUser)
	return loggedInUser
}

func createOrder(t *testing.T, ts *httptest.Server, tasks []string) Order {
	order := Order{Tasks: tasks}
	body, _ := json.Marshal(order)

	resp, err := http.Post(ts.URL+"/orders", "application/json", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Error creating order: %v", err)
	}
	defer resp.Body.Close()

	var createdOrder Order
	json.NewDecoder(resp.Body).Decode(&createdOrder)
	return createdOrder
}

func updateOrderStatus(t *testing.T, ts *httptest.Server, order Order) {
	body, _ := json.Marshal(order)

	req, err := http.NewRequest(http.MethodPut, ts.URL+"/orders", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Error updating order status: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Error sending PUT request: %v", err)
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&order)
}

func rateOrder(t *testing.T, ts *httptest.Server, orderID string, score int) {
	rating := Rating{OrderID: orderID, Score: score}
	body, _ := json.Marshal(rating)

	resp, err := http.Post(ts.URL+"/rate-order", "application/json", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Error rating order: %v", err)
	}
	defer resp.Body.Close()
}
