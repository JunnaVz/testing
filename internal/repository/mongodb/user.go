package mongodb

import (
	"context"
	"lab3/internal/models"
	"lab3/internal/repository/repository_errors"
	"lab3/internal/repository/repository_interfaces"

	"errors"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDB struct {
	ID          uuid.UUID `bson:"_id"`
	Name        string    `bson:"name"`
	Surname     string    `bson:"surname"`
	Address     string    `bson:"address"`
	PhoneNumber string    `bson:"phone_number"`
	Email       string    `bson:"email"`
	Password    string    `bson:"password"`
}

type UserRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) repository_interfaces.IUserRepository {
	return &UserRepository{db: db}
}

func copyUserResultToModel(userDB *UserDB) *models.User {
	return &models.User{
		ID:          userDB.ID,
		Name:        userDB.Name,
		Surname:     userDB.Surname,
		Address:     userDB.Address,
		PhoneNumber: userDB.PhoneNumber,
		Email:       userDB.Email,
		Password:    userDB.Password,
	}
}

func (u UserRepository) Create(user *models.User) (*models.User, error) {
	var collection = u.db.Collection("users")
	ctx := context.Background()
	if user.ID == uuid.Nil {
		user.ID = uuid.New()
	}

	_, err := collection.InsertOne(ctx, UserDB{
		ID:          user.ID,
		Name:        user.Name,
		Surname:     user.Surname,
		Address:     user.Address,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Password:    user.Password,
	})

	if err != nil {
		return nil, repository_errors.InsertError
	}

	return &models.User{
		ID:          user.ID,
		Name:        user.Name,
		Surname:     user.Surname,
		Address:     user.Address,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Password:    user.Password,
	}, nil
}

func (u UserRepository) Delete(id uuid.UUID) error {
	var usersCollection = u.db.Collection("users")
	var ordersCollection = u.db.Collection("orders")
	ctx := context.Background()
	_, err := ordersCollection.DeleteMany(ctx, bson.M{"user_id": id})
	if err != nil {
		return err
	}

	_, err = usersCollection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}
func (u UserRepository) Update(user *models.User) (*models.User, error) {
	var collection = u.db.Collection("users")
	ctx := context.Background()

	filter := bson.M{"_id": user.ID}
	update := bson.M{
		"$set": bson.M{
			"name":         user.Name,
			"surname":      user.Surname,
			"address":      user.Address,
			"phone_number": user.PhoneNumber,
			"email":        user.Email,
			"password":     user.Password,
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, repository_errors.UpdateError
	}

	return &models.User{
		ID:          user.ID,
		Name:        user.Name,
		Surname:     user.Surname,
		Address:     user.Address,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Password:    user.Password,
	}, nil
}

func (u UserRepository) GetUserByID(id uuid.UUID) (*models.User, error) {
	var usersCollection = u.db.Collection("users")
	ctx := context.Background()
	filter := bson.M{"_id": id}

	var user UserDB
	err := usersCollection.FindOne(ctx, filter).Decode(&user)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, repository_errors.DoesNotExist
	} else if err != nil {
		return nil, repository_errors.SelectError
	}

	return &models.User{
		ID:          user.ID,
		Name:        user.Name,
		Surname:     user.Surname,
		Address:     user.Address,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Password:    user.Password,
	}, nil
}

func (u UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var collection = u.db.Collection("users")
	ctx := context.Background()

	filter := bson.M{"email": email}
	var user UserDB
	err := collection.FindOne(ctx, filter).Decode(&user)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return nil, repository_errors.DoesNotExist
	} else if err != nil {
		return nil, repository_errors.SelectError
	}

	return &models.User{
		ID:          user.ID,
		Name:        user.Name,
		Surname:     user.Surname,
		Address:     user.Address,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Password:    user.Password,
	}, nil
}

func (u UserRepository) GetAllUsers() ([]models.User, error) {
	var usersCollection = u.db.Collection("users")
	ctx := context.Background()

	cur, err := usersCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var userModels []models.User
	for cur.Next(context.Background()) {
		var user UserDB
		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		userModels = append(userModels, models.User{
			ID:          user.ID,
			Name:        user.Name,
			Surname:     user.Surname,
			Address:     user.Address,
			PhoneNumber: user.PhoneNumber,
			Email:       user.Email,
			Password:    user.Password,
		})
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return userModels, nil
}
