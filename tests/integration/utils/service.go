package utils

//
//import (
//	"github.com/charmbracelet/log"
//	"github.com/ozontech/allure-go/pkg/framework/runner"
//	"github.com/ozontech/allure-go/pkg/framework/suite"
//	"lab3/internal/models"
//	"lab3/internal/repository/postgres"
//	services "lab3/internal/services"
//	"os"
//	"strconv"
//	"sync"
//	"testing"
//)
//
//func TestRunner(t *testing.T) {
//	db, ctr, ids, _ := NewTestStorage()
//	defer DropTestStorage(db, ctr)
//
//	t.Parallel()
//
//	wg := &sync.WaitGroup{}
//	suits := []runner.TestSuite{
//		&models.CategorySuite{
//			CategoryService: *services.NewCategoryService(postgres.NewCategoryRepository(db), postgres.NewTaskRepository(db), log.New(os.Stdout)),
//			ID:              ids["ID"],
//			Name:            strconv.FormatInt(ids["Name"], 10),
//		},
//	}
//	wg.Add(len(suits))
//
//	for _, s := range suits {
//		go func(s runner.TestSuite) {
//			suite.RunSuite(t, s)
//			wg.Done()
//		}(s)
//	}
//
//	wg.Wait()
//}
