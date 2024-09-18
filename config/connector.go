package config

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/charmbracelet/log"
	_ "github.com/jackc/pgx/v4/stdlib"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

type DbConnectionFlags struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Port     string `mapstructure:"port"`
	DBName   string `mapstructure:"dbname"`
}

func (p *DbConnectionFlags) InitPostgresDB(logger *log.Logger) (*sql.DB, error) {
	logger.Debug("POSTGRES! Start init postgreSQL", "user", p.User, "DBName", p.DBName,
		"host", p.Host, "port", p.Port)

	dsnPGConn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		p.User, p.DBName, p.Password,
		p.Host, p.Port)

	db, err := sql.Open("pgx", dsnPGConn)
	if err != nil {
		logger.Fatal("POSTGRES! Error in method open")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logger.Fatal("POSTGRES! Error in method ping")
		return nil, err
	}

	db.SetMaxOpenConns(10)

	logger.Info("POSTGRES! Successfully init postgreSQL")
	return db, nil
}

func (p *DbConnectionFlags) InitMongoDB(logger *log.Logger) (*mongo.Database, error) {
	logger.Debug("MONGO! Start init mongoDB", "user", p.User, "DBName", p.DBName,
		"host", p.Host, "port", p.Port)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	dsnMongoConn := fmt.Sprintf("mongodb://%s:%s@%s:%s", p.User, p.Password, p.Host, p.Port)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsnMongoConn))
	if err != nil {
		logger.Fatal("MONGO! Error in method connect")
		return nil, err
	}

	// Send a ping to confirm a successful connection
	var result bson.M
	if err = client.Database(p.DBName).RunCommand(ctx, bson.D{{"ping", 1}}).Decode(&result); err != nil {
		return nil, err
	}

	logger.Info("MONGO! Successfully init mongoDB")
	return client.Database(p.DBName), nil
}
