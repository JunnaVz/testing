package mongodb

import (
	"context"
	"lab3/internal/models"
	"lab3/internal/repository/repository_errors"
	"lab3/internal/repository/repository_interfaces"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"errors"
)

type TaskDB struct {
	ID             uuid.UUID `bson:"_id"`
	Name           string    `bson:"name"`
	PricePerSingle float64   `bson:"price_per_single"`
	Category       int       `bson:"category"`
}

type TaskRepository struct {
	db *mongo.Database
}

func NewTaskRepository(db *mongo.Database) repository_interfaces.ITaskRepository {
	return &TaskRepository{db: db}
}

func copyTaskResultToModel(taskDB *TaskDB) *models.Task {
	return &models.Task{
		ID:             taskDB.ID,
		Name:           taskDB.Name,
		PricePerSingle: taskDB.PricePerSingle,
		Category:       taskDB.Category,
	}
}

func (t TaskRepository) Create(task *models.Task) (*models.Task, error) {
	ctx := context.Background()
	var collection = t.db.Collection("tasks")
	if task.ID == uuid.Nil {
		task.ID = uuid.New()
	}
	_, err := collection.InsertOne(ctx, TaskDB{
		ID:             task.ID,
		Name:           task.Name,
		PricePerSingle: task.PricePerSingle,
		Category:       task.Category,
	})

	if err != nil {
		return nil, repository_errors.InsertError
	}

	return &models.Task{
		ID:             task.ID,
		Name:           task.Name,
		PricePerSingle: task.PricePerSingle,
		Category:       task.Category,
	}, nil
}

func (t TaskRepository) Delete(id uuid.UUID) error {
	var collection = t.db.Collection("tasks")
	var filter = bson.M{"_id": id}

	result, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return repository_errors.DeleteError
	}

	if result.DeletedCount == 0 {
		return errors.New("no task found to delete")
	}

	return nil
}

func (t TaskRepository) Update(task *models.Task) (*models.Task, error) {
	var collection = t.db.Collection("tasks")
	var filter = bson.M{"_id": task.ID}
	update := bson.M{
		"$set": bson.M{
			"name":             task.Name,
			"price_per_single": task.PricePerSingle,
			"category":         task.Category,
		},
	}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, repository_errors.UpdateError
	}

	return task, nil
}

func (t TaskRepository) GetTaskByID(id uuid.UUID) (*models.Task, error) {
	var collection = t.db.Collection("tasks")

	var task TaskDB
	filter := bson.M{"_id": id}

	err := collection.FindOne(context.Background(), filter).Decode(&task)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, repository_errors.DoesNotExist
	} else if err != nil {
		return nil, repository_errors.SelectError
	}

	taskModels := copyTaskResultToModel(&task)

	return taskModels, nil
}

func (t TaskRepository) GetTaskByName(name string) (*models.Task, error) {
	var collection = t.db.Collection("tasks")
	var filter = bson.M{"name": name}

	var task TaskDB
	err := collection.FindOne(context.Background(), filter).Decode(&task)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, repository_errors.DoesNotExist
	} else if err != nil {
		return nil, repository_errors.SelectError
	}

	return copyTaskResultToModel(&task), nil
}

func (t TaskRepository) GetAllTasks() ([]models.Task, error) {
	var collection = t.db.Collection("tasks")

	filter := bson.M{}

	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var tasks []models.Task
	for cur.Next(context.Background()) {
		var task TaskDB
		err := cur.Decode(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, *copyTaskResultToModel(&task))
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t TaskRepository) GetTasksInCategory(category int) ([]models.Task, error) {
	var collection = t.db.Collection("tasks")

	filter := bson.M{"category": category}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	defer cur.Close(context.Background())

	var tasks []models.Task
	for cur.Next(context.Background()) {
		var task TaskDB
		err := cur.Decode(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, *copyTaskResultToModel(&task))
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
