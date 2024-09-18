package mongodb

import (
	"context"
	"errors"
	"lab3/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryDB struct {
	ID   int    `bson:"_id"`
	Name string `bson:"name"`
}

type CategoryRepository struct {
	db *mongo.Database
}

func NewCategoryRepository(db *mongo.Database) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func getNextSequence(db *mongo.Database, sequenceName string) (int, error) {
	collection := db.Collection("counters")

	filter := bson.M{"_id": sequenceName}
	update := bson.M{"$inc": bson.M{"seq": 1}}

	var result struct {
		Seq int `bson:"seq"`
	}

	err := collection.FindOneAndUpdate(context.Background(), filter, update).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Если документа не существует, создаем его
			_, err := collection.InsertOne(context.Background(), bson.M{"_id": sequenceName, "seq": 1})
			if err != nil {
				return 0, err
			}
			return 1, nil
		}
		return 0, err
	}

	return result.Seq, nil
}

func (c CategoryRepository) GetAll() ([]models.Category, error) {
	var collection = c.db.Collection("categories")
	ctx := context.Background()

	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var categories []models.Category
	for cur.Next(context.Background()) {
		var category CategoryDB
		err := cur.Decode(&category)
		if err != nil {
			return nil, err
		}
		categories = append(categories, models.Category{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func (c CategoryRepository) GetByID(id int) (*models.Category, error) {
	var collection = c.db.Collection("categories")
	ctx := context.Background()

	var category CategoryDB
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&category)
	if err != nil {
		return nil, err
	}

	return &models.Category{
		ID:   category.ID,
		Name: category.Name,
	}, nil
}

func (c CategoryRepository) Create(category *models.Category) (*models.Category, error) {
	var collection = c.db.Collection("categories")
	ctx := context.Background()

	id, err := getNextSequence(c.db, "categoryid")
	if err != nil {
		return nil, err
	}

	_, err = collection.InsertOne(ctx, bson.M{"_id": id, "name": category.Name})
	if err != nil {
		return nil, err
	}

	return &models.Category{
		ID:   id,
		Name: category.Name,
	}, nil
}

func (c CategoryRepository) Update(category *models.Category) (*models.Category, error) {
	var collection = c.db.Collection("categories")
	ctx := context.Background()

	filter := bson.M{"_id": category.ID}
	update := bson.M{"$set": bson.M{"name": category.Name}}
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return &models.Category{
		ID:   category.ID,
		Name: category.Name,
	}, nil
}

func (c CategoryRepository) Delete(id int) error {
	var collection = c.db.Collection("categories")
	ctx := context.Background()

	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("no task found to delete")
	}

	return nil
}
