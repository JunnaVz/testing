package mongodb

import (
	"context"
	"errors"
	"lab3/internal/models"
	"lab3/internal/repository/repository_errors"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"lab3/internal/repository/repository_interfaces"

	"github.com/google/uuid"
)

type OrderDB struct {
	ID           uuid.UUID `bson:"_id"`
	WorkerID     uuid.UUID `bson:"worker_id"`
	UserID       uuid.UUID `bson:"user_id"`
	Status       int       `bson:"status"`
	Address      string    `bson:"address"`
	CreationDate time.Time `bson:"creation_date"`
	Deadline     time.Time `bson:"deadline"`
	Rate         int       `bson:"rate"`
}

type OrderRepository struct {
	db *mongo.Database
}

func NewOrderRepository(db *mongo.Database) repository_interfaces.IOrderRepository {
	return &OrderRepository{db: db}
}

func copyOrderResultToModel(orderDB *OrderDB) *models.Order {
	return &models.Order{
		ID:           orderDB.ID,
		WorkerID:     orderDB.WorkerID,
		UserID:       orderDB.UserID,
		Status:       orderDB.Status,
		Address:      orderDB.Address,
		CreationDate: orderDB.CreationDate,
		Deadline:     orderDB.Deadline,
		Rate:         orderDB.Rate,
	}
}

func (o OrderRepository) Create(order *models.Order, orderedTasks []models.OrderedTask) (*models.Order, error) {
	var collection = o.db.Collection("orders")
	var m2mCollection = o.db.Collection("order_contains_tasks")

	if order.ID == uuid.Nil {
		order.ID = uuid.New()
	}

	_, err := collection.InsertOne(context.Background(), OrderDB{
		ID:           order.ID,
		WorkerID:     order.WorkerID,
		UserID:       order.UserID,
		Status:       order.Status,
		Address:      order.Address,
		CreationDate: order.CreationDate,
		Deadline:     order.Deadline,
		Rate:         order.Rate,
	})

	if err != nil {
		return nil, repository_errors.InsertError
	}

	var orderedTasksInterface []interface{}
	for _, data := range orderedTasks {
		orderedTasksInterface = append(orderedTasksInterface, data)
	}

	_, err = m2mCollection.InsertMany(context.Background(), orderedTasksInterface)

	return order, nil
}

func (o OrderRepository) Delete(id uuid.UUID) error {
	var ordersCollection = o.db.Collection("orders")
	var m2mCollection = o.db.Collection("order_contains_tasks")
	var filter = map[string]interface{}{"_id": id}

	_, err := m2mCollection.DeleteMany(context.Background(), filter)
	if err != nil {
		return repository_errors.DeleteError
	}
	result, err := ordersCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return repository_errors.DeleteError
	}

	if result.DeletedCount == 0 {
		return repository_errors.DoesNotExist
	}

	return nil
}

func (o OrderRepository) Update(order *models.Order) (*models.Order, error) {
	var collection = o.db.Collection("orders")
	var filter = map[string]interface{}{"_id": order.ID}

	update := map[string]interface{}{
		"$set": map[string]interface{}{
			"worker_id":     order.WorkerID,
			"user_id":       order.UserID,
			"status":        order.Status,
			"address":       order.Address,
			"creation_date": order.CreationDate,
			"deadline":      order.Deadline,
			"rate":          order.Rate,
		},
	}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, repository_errors.UpdateError
	}

	return order, nil
}

func (o OrderRepository) GetOrderByID(id uuid.UUID) (*models.Order, error) {
	var collection = o.db.Collection("orders")
	var filter = map[string]interface{}{"_id": id}

	var order OrderDB
	err := collection.FindOne(context.Background(), filter).Decode(&order)
	if err != nil {
		return nil, repository_errors.DoesNotExist
	}

	return copyOrderResultToModel(&order), nil
}

func (o OrderRepository) GetTasksInOrder(id uuid.UUID) ([]models.Task, error) {
	var m2mCollection = o.db.Collection("order_contains_tasks")
	var tasksCollection = o.db.Collection("tasks")

	var order OrderDB
	err := m2mCollection.FindOne(context.Background(), map[string]interface{}{"order_id": id}).Decode(&order)
	if err != nil {
		return nil, repository_errors.DoesNotExist
	}

	var tasks []models.Task
	cursor, err := m2mCollection.Find(context.Background(), map[string]interface{}{"order_id": id})
	if err != nil {
		return nil, repository_errors.SelectError
	}

	for cursor.Next(context.Background()) {
		var orderedTask models.OrderedTask
		err := cursor.Decode(&orderedTask)
		if err != nil {
			return nil, repository_errors.SelectError
		}

		var task models.Task
		err = tasksCollection.FindOne(context.Background(), map[string]interface{}{"_id": orderedTask.Task.ID}).Decode(&task)
		if err != nil {
			return nil, repository_errors.DoesNotExist
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (o OrderRepository) GetCurrentOrderByUserID(id uuid.UUID) (*models.Order, error) {
	var collection = o.db.Collection("orders")
	var filter = map[string]interface{}{"user_id": id}
	var order OrderDB

	opts := options.FindOne().SetSort(bson.D{{"creation_date", -1}})
	err := collection.FindOne(context.Background(), filter, opts).Decode(&order)
	if err != nil {
		return nil, repository_errors.DoesNotExist
	}

	return copyOrderResultToModel(&order), nil
}

func (o OrderRepository) GetAllOrdersByUserID(id uuid.UUID) ([]models.Order, error) {
	var collection = o.db.Collection("orders")
	var filter = map[string]interface{}{"user_id": id}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, repository_errors.SelectError
	}

	var orders []models.Order
	for cursor.Next(context.Background()) {
		var order OrderDB
		err := cursor.Decode(&order)
		if err != nil {
			return nil, repository_errors.SelectError
		}
		orders = append(orders, *copyOrderResultToModel(&order))
	}

	return orders, nil
}

func (o OrderRepository) Filter(params map[string]string) ([]models.Order, error) {
	var collection = o.db.Collection("orders")

	filter := bson.M{}
	for field, value := range params {
		values := strings.Split(value, ",")
		if len(values) > 1 {
			var inValues []interface{}
			for _, v := range values {
				if v == "null" {
					inValues = append(inValues, nil)
				} else if field == "status" {
					status, err := strconv.Atoi(v)
					if err != nil {
						return nil, err
					}
					inValues = append(inValues, status)
				} else {
					inValues = append(inValues, v)
				}
			}
			filter[field] = bson.M{"$in": inValues}
		} else {
			if value == "null" {
				filter[field] = nil
			} else if value == "not null" {
				filter[field] = bson.M{"$ne": nil}
			} else if field == "status" {
				status, err := strconv.Atoi(value)
				if err != nil {
					return nil, err
				}
				filter[field] = status
			} else {
				filter[field] = value
			}
		}
	}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, repository_errors.SelectError
	}

	var orders []models.Order
	for cursor.Next(context.Background()) {
		var order OrderDB
		err := cursor.Decode(&order)
		if err != nil {
			return nil, repository_errors.SelectError
		}
		orders = append(orders, *copyOrderResultToModel(&order))
	}

	return orders, nil
}
func (o OrderRepository) AddTaskToOrder(orderID uuid.UUID, taskID uuid.UUID) error {
	var m2mCollection = o.db.Collection("order_contains_tasks")
	var ordersCollection = o.db.Collection("orders")

	var order OrderDB
	err := ordersCollection.FindOne(context.Background(), map[string]interface{}{"_id": orderID}).Decode(&order)
	if err != nil {
		return repository_errors.DoesNotExist
	}

	_, err = m2mCollection.InsertOne(context.Background(), bson.M{"order_id": orderID, "task_id": taskID})
	if err != nil {
		return repository_errors.InsertError
	}

	return nil
}

func (o OrderRepository) RemoveTaskFromOrder(orderID uuid.UUID, taskID uuid.UUID) error {
	var m2mCollection = o.db.Collection("order_contains_tasks")
	var ordersCollection = o.db.Collection("orders")

	var order OrderDB
	err := ordersCollection.FindOne(context.Background(), map[string]interface{}{"_id": orderID}).Decode(&order)
	if err != nil {
		return repository_errors.DoesNotExist
	}

	_, err = m2mCollection.DeleteOne(context.Background(), bson.M{"order_id": orderID, "task_id": taskID})
	if err != nil {
		return repository_errors.DeleteError
	}

	return nil
}

func (o OrderRepository) UpdateTaskQuantity(orderID uuid.UUID, taskID uuid.UUID, quantity int) error {
	var m2mCollection = o.db.Collection("order_contains_tasks")
	var ordersCollection = o.db.Collection("orders")

	var order OrderDB
	err := ordersCollection.FindOne(context.Background(), map[string]interface{}{"_id": orderID}).Decode(&order)
	if err != nil {
		return repository_errors.DoesNotExist
	}

	_, err = m2mCollection.UpdateOne(context.Background(), bson.M{"order_id": orderID, "task_id": taskID}, bson.M{"$set": bson.M{"quantity": quantity}})
	if err != nil {
		return repository_errors.UpdateError
	}

	return nil
}

func (o OrderRepository) GetTaskQuantity(orderID uuid.UUID, taskID uuid.UUID) (int, error) {
	var m2mCollection = o.db.Collection("order_contains_tasks")

	var orderedTask models.OrderedTask
	err := m2mCollection.FindOne(context.Background(), bson.M{"order_id": orderID, "task_id": taskID}).Decode(&orderedTask)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return 0, repository_errors.DoesNotExist
		}
		return 0, repository_errors.SelectError
	}

	return orderedTask.Quantity, nil
}
