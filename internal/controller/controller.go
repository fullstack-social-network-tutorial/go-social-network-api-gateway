package controller

import (
	"context"
	"go-service/internal/configs"
	"go-service/internal/controller/handlers"
	"go-service/pkg/logger"
	"net/http"
)

type apiController struct {
	ctx    context.Context
	router *http.ServeMux
	auth   handlers.AuthHandler
}

func (controller *apiController) SetUpRoute() {

	auth := "/auth"
	controller.router.HandleFunc(http.MethodPost+" "+auth+"/login", controller.auth.Login)
}

func NewAPIController(ctx context.Context, router *http.ServeMux, configs configs.Config, logger *logger.Logger) *apiController {
	auth := handlers.NewAuthHandler(configs.Outbound.Auth, configs.Key.ApiGateway, logger)
	return &apiController{
		ctx:    ctx,
		router: router,
		auth:   auth,
	}
}
