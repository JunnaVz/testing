package utils

//
//import (
//	"context"
//	"github.com/google/uuid"
//	"github.com/jmoiron/sqlx"
//	"github.com/testcontainers/testcontainers-go"
//	container "github.com/testcontainers/testcontainers-go/modules/postgres"
//	"github.com/testcontainers/testcontainers-go/wait"
//	"lab3/internal/models"
//	"lab3/internal/repository/postgres"
//	"lab3/internal/repository/repository_interfaces"
//	"os"
//	"strings"
//	"time"
//)
//
//const (
//	pgImage    = "docker.io/postgres:16-alpine"
//	dbName     = "tests"
//	dbUsername = "postgres"
//	dbPassword = "admin"
//)
//
//var ids map[string]int64
//var ids2 map[string]uuid.UUID
//
//func NewTestStorage() (*sqlx.DB, *container.PostgresContainer, map[string]int64, map[string]uuid.UUID) {
//	ctr, err := container.Run(
//		context.TODO(),
//		pgImage,
//		container.WithInitScripts(os.Getenv("DB_INIT_PATH")),
//		container.WithDatabase(dbName),
//		container.WithUsername(dbUsername),
//		container.WithPassword(dbPassword),
//		testcontainers.WithWaitStrategy(
//			wait.ForLog("database system is ready to accept connections").
//				WithOccurrence(2).
//				WithStartupTimeout(5*time.Second)),
//	)
//	if err != nil {
//		panic(err)
//	}
//
//	connString, err := ctr.ConnectionString(context.TODO(), "sslmode=disable")
//	if err != nil {
//		panic(err)
//	}
//
//	conn, err := sqlx.Open("pgx", connString)
//	if err != nil {
//		panic(err)
//	}
//
//	ids = map[string]int64{}
//	ids["categoryID"] = initTestCategoryStorage(postgres.NewCategoryRepository(conn))
//	ids2 = map[string]uuid.UUID{}
//	ids2["taskID"] = initTestTaskStorage(postgres.NewTaskRepository(conn))
//
//	return conn, ctr, ids, ids2
//}
//
//func DropTestStorage(testDB *sqlx.DB, ctr *container.PostgresContainer) {
//	defer func() {
//		testDB.Close()
//		ctr.Terminate(context.TODO())
//	}()
//
//	err := postgres.NewCategoryRepository(testDB).Delete(int(ids["categoryID"]))
//	if err != nil {
//		panic(err)
//	}
//	err = postgres.NewTaskRepository(testDB).Delete(ids2["taskID"])
//	if err != nil {
//		panic(err)
//	}
//}
//
//func initTestCategoryStorage(storage *postgres.CategoryRepository) int64 {
//	category, err := storage.Create(&models.Category{
//		Name: "TestCategory",
//	})
//	if err != nil && !strings.Contains(err.Error(), "constraint") {
//		panic(err)
//	}
//
//	return int64(category.ID)
//}
//
//func initTestTaskStorage(storage repository_interfaces.ITaskRepository) uuid.UUID {
//	task, err := storage.Create(&models.Task{
//		Name:           "TestTask",
//		PricePerSingle: 100.0,
//		Category:       1,
//	})
//	if err != nil && !strings.Contains(err.Error(), "constraint") {
//		panic(err)
//	}
//
//	return task.ID
//}
