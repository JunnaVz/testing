package mongodb

import (
	"lab3/config"
	"lab3/internal/repository/repository_errors"
	"lab3/internal/repository/repository_interfaces"

	"github.com/charmbracelet/log"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoConnection struct {
	DB     *mongo.Database
	Config config.Config
}

func NewMongoConnection(Postgres config.DbConnectionFlags, logger *log.Logger) (*MongoConnection, error) {
	fields := new(MongoConnection)
	var err error

	fields.Config.DBFlags = Postgres

	fields.DB, err = fields.Config.DBFlags.InitMongoDB(logger)
	if err != nil {
		logger.Error("POSTGRES! Error parse config for mongoDB")
		return nil, repository_errors.ConnectionError
	}

	logger.Info("Mongo! Successfully create postgres repository fields")

	return fields, nil
}

func CreateTaskRepository(fields *MongoConnection) repository_interfaces.ITaskRepository {
	return NewTaskRepository(fields.DB)
}

func CreateCategoryRepository(fields *MongoConnection) repository_interfaces.ICategoryRepository {
	return NewCategoryRepository(fields.DB)
}

func CreateUserRepository(fields *MongoConnection) repository_interfaces.IUserRepository {
	return NewUserRepository(fields.DB)
}

func CreateWorkerRepository(fields *MongoConnection) repository_interfaces.IWorkerRepository {
	return NewWorkerRepository(fields.DB)
}

func CreateOrderRepository(fields *MongoConnection) repository_interfaces.IOrderRepository {
	return NewOrderRepository(fields.DB)
}
