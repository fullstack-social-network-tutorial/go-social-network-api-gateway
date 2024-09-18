package main

import (
	"context"
	"go-service/app"
	"go-service/internal/configs"
	"go-service/internal/controller"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	mux := http.NewServeMux()
	ctx := context.Background()
	app := app.NewApplication(ctx)
	ctr := controller.NewAPIController(ctx, app, mux)
	ctr.SetUpRoute()
	http.ListenAndServe(":8081", nil)
}

func getConfig() (configs.Config, error) {
	configs := configs.Config{}
	configFile, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		return configs, err
	}

	err = yaml.Unmarshal(configFile, &configs)
	if err != nil {
		return configs, err
	}
	return configs, nil
}
