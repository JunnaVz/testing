package main

import (
	"fmt"
	"lab3/cmd"
	"lab3/internal/models"
	"lab3/internal/registry"
	"lab3/server"

	"github.com/charmbracelet/log"
)

func main() {
	app := registry.App{}

	err := app.Config.ParseConfig("config.json", "config")
	if err != nil {
		log.Fatal(err)
	}

	err = app.Run()

	if err != nil {
		fmt.Println("Error")
		log.Fatal(err)
	}

	err = initAdmin(app.Services)
	if err != nil {
		log.Fatal(err)
		return
	}

	if app.Config.Mode == "cmd" {
		cmdErr := cmd.RunMenu(app.Services)
		if cmdErr != nil {
			log.Fatal(cmdErr)
			return
		}
	} else if app.Config.Mode == "server" {
		log.Info("Start with server!")
		err = server.RunServer(&app)
		if err != nil {
			log.Fatal(err)
			return
		}
	} else {
		log.Error("Wrong app mode", "mode", app.Config.Mode)
	}
}

func initAdmin(services *registry.Services) error {
	admins, err := services.WorkerService.GetWorkersByRole(models.ManagerRole)
	if err != nil {
		return err
	}

	if len(admins) == 0 {
		defaultAdmin := &models.Worker{
			Email:       "default@admin.com",
			Name:        "admin",
			Surname:     "admin",
			Role:        models.ManagerRole,
			PhoneNumber: "+79999999999",
			Address:     "admin address",
		}
		_, err = services.WorkerService.Create(defaultAdmin, "admin123")
		if err != nil {
			return err
		}

		log.Info("Default admin created")
	}

	return nil
}
