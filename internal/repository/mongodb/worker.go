package mongodb

import (
	"context"
	"errors"
	"lab3/internal/models"
	"lab3/internal/repository/repository_errors"
	"lab3/internal/repository/repository_interfaces"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/google/uuid"
)

type WorkerDB struct {
	ID          uuid.UUID `bson:"_id"`
	Name        string    `bson:"name"`
	Surname     string    `bson:"surname"`
	Address     string    `bson:"address"`
	PhoneNumber string    `bson:"phone_number"`
	Email       string    `bson:"email"`
	Role        int       `bson:"role"`
	Password    string    `bson:"password"`
}

type WorkerRepository struct {
	db *mongo.Database
}

func NewWorkerRepository(db *mongo.Database) repository_interfaces.IWorkerRepository {
	return &WorkerRepository{db: db}
}

func copyWorkerResultToModel(workerDB *WorkerDB) *models.Worker {
	return &models.Worker{
		ID:          workerDB.ID,
		Name:        workerDB.Name,
		Surname:     workerDB.Surname,
		Address:     workerDB.Address,
		PhoneNumber: workerDB.PhoneNumber,
		Email:       workerDB.Email,
		Role:        workerDB.Role,
		Password:    workerDB.Password,
	}
}

func (w WorkerRepository) Create(worker *models.Worker) (*models.Worker, error) {
	var collection = w.db.Collection("workers")
	if worker.ID == uuid.Nil {
		worker.ID = uuid.New()
	}

	_, err := collection.InsertOne(context.Background(), WorkerDB{
		ID:          worker.ID,
		Name:        worker.Name,
		Surname:     worker.Surname,
		Address:     worker.Address,
		PhoneNumber: worker.PhoneNumber,
		Email:       worker.Email,
		Role:        worker.Role,
		Password:    worker.Password,
	})

	if err != nil {
		return nil, repository_errors.InsertError
	}

	return &models.Worker{
		ID:          worker.ID,
		Name:        worker.Name,
		Surname:     worker.Surname,
		Address:     worker.Address,
		PhoneNumber: worker.PhoneNumber,
		Email:       worker.Email,
		Role:        worker.Role,
		Password:    worker.Password,
	}, nil
}

func (w WorkerRepository) Update(worker *models.Worker) (*models.Worker, error) {
	var collection = w.db.Collection("workers")
	var filter = bson.M{"id": worker.ID}
	var update = bson.M{"$set": bson.M{
		"name":         worker.Name,
		"surname":      worker.Surname,
		"address":      worker.Address,
		"phone_number": worker.PhoneNumber,
		"email":        worker.Email,
		"role":         worker.Role,
		"password":     worker.Password,
	}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, repository_errors.UpdateError
	}

	return worker, nil
}

func (w WorkerRepository) Delete(id uuid.UUID) error {
	var collection = w.db.Collection("workers")
	var filter = bson.M{"id": id}
	_, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return repository_errors.DeleteError
	}

	return nil
}

func (w WorkerRepository) GetWorkerByID(id uuid.UUID) (*models.Worker, error) {
	var collection = w.db.Collection("workers")
	var filter = bson.M{"_id": id}
	var worker WorkerDB

	err := collection.FindOne(context.Background(), filter).Decode(&worker)
	if err != nil {
		return nil, repository_errors.DoesNotExist
	}

	return copyWorkerResultToModel(&worker), nil
}

func (w WorkerRepository) GetAllWorkers() ([]models.Worker, error) {
	var collection = w.db.Collection("workers")
	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	var workerModels []models.Worker
	for cur.Next(context.Background()) {
		var worker WorkerDB
		err := cur.Decode(&worker)
		if err != nil {
			return nil, err
		}
		workerModels = append(workerModels, models.Worker{
			ID:          worker.ID,
			Name:        worker.Name,
			Surname:     worker.Surname,
			Address:     worker.Address,
			PhoneNumber: worker.PhoneNumber,
			Email:       worker.Email,
			Password:    worker.Password,
		})
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return workerModels, nil
}

func (w WorkerRepository) GetWorkerByEmail(email string) (*models.Worker, error) {
	var collection = w.db.Collection("workers")
	var filter = bson.M{"email": email}
	var worker WorkerDB

	err := collection.FindOne(context.Background(), filter).Decode(&worker)
	if err != nil {
		return nil, repository_errors.DoesNotExist
	}

	return copyWorkerResultToModel(&worker), nil
}

func (w WorkerRepository) GetWorkersByRole(role int) ([]models.Worker, error) {
	var collection = w.db.Collection("workers")
	var filter = bson.M{"role": role}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	var workerModels []models.Worker
	for cur.Next(context.Background()) {
		var worker WorkerDB
		err := cur.Decode(&worker)
		if err != nil {
			return nil, err
		}
		workerModels = append(workerModels, models.Worker{
			ID:          worker.ID,
			Name:        worker.Name,
			Surname:     worker.Surname,
			Address:     worker.Address,
			PhoneNumber: worker.PhoneNumber,
			Email:       worker.Email,
			Password:    worker.Password,
		})
	}
	return workerModels, nil

}

func (w WorkerRepository) GetAverageOrderRate(worker *models.Worker) (float64, error) {
	var ordersCollection = w.db.Collection("orders")
	ctx := context.Background()

	pipeline := mongo.Pipeline{
		{{"$match", bson.D{
			{"worker_id", worker.ID},
			{"status", 3},
			{"rate", bson.M{"$ne": 0}},
		}}},
		{{"$group", bson.D{
			{"_id", nil},
			{"averageRate", bson.M{"$avg": "$rate"}},
		}}},
	}

	cursor, err := ordersCollection.Aggregate(ctx, pipeline)
	if err != nil {
		return 0, repository_errors.SelectError
	}

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		return 0, repository_errors.SelectError
	}

	if len(results) == 0 {
		return 0, nil
	}

	averageRate, ok := results[0]["averageRate"].(float64)
	if !ok {
		return 0, errors.New("could not convert average rate to float64")
	}

	return averageRate, nil
}
