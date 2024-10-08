package main

import (
	"context"
	"fmt"
	"go-service/internal/configs"
	"go-service/internal/controller"
	"go-service/pkg/handler_fnc"
	"go-service/pkg/logger"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	mux := http.NewServeMux()
	ctx := context.Background()
	configs, err := getConfig()
	if err != nil {
		log.Fatal(err)
		return
	}
	logger := logger.NewLogger()
	ctr := controller.NewAPIController(ctx, mux, configs, logger)
	ctr.SetUpRoute()

	done := make(chan bool)
	go http.ListenAndServe(fmt.Sprintf("%v:%v", configs.Address.Host, configs.Port), handler_fnc.LogRequestHandler(mux, logger))
	log.Printf("Server started at %v:%v", configs.Host, configs.Port)
	<-done
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
