package integration

import (
	"context"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"github.com/testcontainers/testcontainers-go"
	container "github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"lab3/internal/models"
	"lab3/internal/repository/repository_interfaces"
	services "lab3/internal/services"

	"lab3/internal/repository/postgres"

	"strings"
	"time"

	"os"
	"strconv"
	"sync"
	"testing"
)

type CategoryBuilder struct {
	category models.Category
}

type CategorySuite struct {
	suite.Suite

	categoryService services.CategoryService
	ID              int64
	Name            string
}

const (
	pgImage    = "docker.io/postgres:16-alpine"
	dbName     = "tests"
	dbUsername = "postgres"
	dbPassword = "admin"
)

var ids map[string]int64
var ids2 map[string]uuid.UUID

func NewCategoryBuilder() *CategoryBuilder {
	return &CategoryBuilder{
		category: models.Category{
			ID:   1,
			Name: "DefaultCategory",
		},
	}
}

func (b *CategoryBuilder) WithID(id int) *CategoryBuilder {
	b.category.ID = id
	return b
}

func (b *CategoryBuilder) WithName(name string) *CategoryBuilder {
	b.category.Name = name
	return b
}

func (b *CategoryBuilder) Build() *models.Category {
	return &b.category
}

var CategoryMother = struct {
	Default        func() *models.Category
	WithID         func(id int) *models.Category
	WithName       func(name string) *models.Category
	CustomCategory func(id int, name string) *models.Category
}{
	Default: func() *models.Category {
		return &models.Category{
			ID:   1,
			Name: "DefaultCategory",
		}
	},
	WithID: func(id int) *models.Category {
		return &models.Category{
			ID:   id,
			Name: "CategoryWithSpecificID",
		}
	},
	WithName: func(name string) *models.Category {
		return &models.Category{
			ID:   2,
			Name: name,
		}
	},
	CustomCategory: func(id int, name string) *models.Category {
		return &models.Category{
			ID:   id,
			Name: name,
		}
	},
}

func NewTestStorage() (*sqlx.DB, *container.PostgresContainer, map[string]int64, map[string]uuid.UUID) {
	ctr, err := container.Run(
		context.TODO(),
		pgImage,
		container.WithInitScripts("../../db/sql/init.sql"),
		container.WithDatabase(dbName),
		container.WithUsername(dbUsername),
		container.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(30*time.Second)), // Increased timeout
	)

	if err != nil {
		panic(err)
	}

	ports, err := ctr.Container.MappedPort(context.TODO(), "5432/tcp")
	if err != nil {
		panic(err)
	}
	println(ports)

	host, err := ctr.Host(context.Background())
	if err != nil {
		panic(err)
	}
	println(host)

	connString, err := ctr.ConnectionString(context.TODO(), "sslmode=disable")
	if err != nil {
		panic(err)
	}
	println(connString)

	dsnPGConn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		dbUsername, dbName, dbPassword,
		host, ports.Port())

	println(dsnPGConn)

	conn, err := sqlx.Open("pgx", dsnPGConn)
	if err != nil {
		panic(err)
	}
	err = conn.Ping()
	if err != nil {
		panic(err)
	}

	ids = map[string]int64{}
	ids["categoryID"] = initTestCategoryStorage(postgres.NewCategoryRepository(conn))
	ids2 = map[string]uuid.UUID{}
	ids2["taskID"] = initTestTaskStorage(postgres.NewTaskRepository(conn))

	return conn, ctr, ids, ids2
}

//// Example of NewTestStorage function
//func NewTestStorage() (*sqlx.DB, error) {
//	req := testcontainers.ContainerRequest{
//		Image:        "docker.io/postgres:16-alpine",
//		ExposedPorts: []string{"5432/tcp"},
//		Env: map[string]string{
//			"POSTGRES_USER":     "postgres",
//			"POSTGRES_PASSWORD": "0252",
//			"POSTGRES_DB":       "postgres",
//		},
//	}
//
//	// Check if DB_INIT_PATH is set and add it to the container request
//	dbInitPath := os.Getenv("DB_INIT_PATH")
//	if dbInitPath != "" {
//		req.Files = []testcontainers.ContainerFile{
//			{
//				HostFilePath: dbInitPath,
//				ContainerFilePath: "/docker-entrypoint-initdb.d/init.sql",
//			},
//		}
//	} else {
//		return nil, fmt.Errorf("DB_INIT_PATH environment variable is not set")
//	}
//
//	container, err := testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
//		ContainerRequest: req,
//		Started:          true,
//	})
//	if err != nil {
//		return nil, fmt.Errorf("create container: %w", err)
//	}
//
//	return &TestStorage{container: container}, nil
//}

func DropTestStorage(testDB *sqlx.DB, ctr *container.PostgresContainer) {
	defer func() {
		err := testDB.Close()
		if err != nil {
			return
		}
		err = ctr.Terminate(context.TODO())
		if err != nil {
			return
		}
	}()

	err := postgres.NewCategoryRepository(testDB).Delete(int(ids["categoryID"]))
	if err != nil {
		panic(err)
	}
	err = postgres.NewTaskRepository(testDB).Delete(ids2["taskID"])
	if err != nil {
		panic(err)
	}
}

func initTestCategoryStorage(storage *postgres.CategoryRepository) int64 {
	category, err := storage.Create(&models.Category{
		Name: "TestCategory",
	})
	if err != nil && !strings.Contains(err.Error(), "constraint") {
		panic(err)
	}

	return int64(category.ID)
}

func initTestTaskStorage(storage repository_interfaces.ITaskRepository) uuid.UUID {
	task, err := storage.Create(&models.Task{
		Name:           "TestTask",
		PricePerSingle: 100.0,
		Category:       1,
	})
	if err != nil && !strings.Contains(err.Error(), "constraint") {
		panic(err)
	}

	return task.ID
}

func TestRunner(t *testing.T) {
	db, ctr, ids, _ := NewTestStorage()
	defer DropTestStorage(db, ctr)

	t.Parallel()

	wg := &sync.WaitGroup{}
	suits := []runner.TestSuite{
		&CategorySuite{
			categoryService: *services.NewCategoryService(postgres.NewCategoryRepository(db), postgres.NewTaskRepository(db), log.New(os.Stdout)),
			ID:              ids["ID"],
			Name:            strconv.FormatInt(ids["Name"], 10),
		},
	}
	wg.Add(len(suits))

	for _, s := range suits {
		go func(s runner.TestSuite) {
			suite.RunSuite(t, s)
			wg.Done()
		}(s)
	}

	wg.Wait()
}

func (s *CategorySuite) Test_Category_GetCategory(t provider.T) {
	t.Title("[INT GetCategory] Success")
	t.Tags("category")
	t.Parallel()
	t.WithNewStep("Success", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		request := NewCategoryBuilder().
			WithName("TestCategory").
			Build()

		sCtx.WithNewParameters("ctx", ctx, "request", request)

		category, err := s.categoryService.Create(request.Name)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(category)

		category, err = s.categoryService.GetByID(category.ID)

		sCtx.Assert().NoError(err)
		sCtx.Assert().NotNil(category)
		sCtx.Assert().Equal("TestCategory", category.Name)
	})
}
