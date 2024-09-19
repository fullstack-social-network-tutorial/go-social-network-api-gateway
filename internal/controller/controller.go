package controller

import (
	"context"
	"go-service/internal/configs"
	"go-service/internal/controller/handlers"
	"net/http"
)

type apiController struct {
	ctx  context.Context
	mux  *http.ServeMux
	auth handlers.AuthHandler
}

func (controller *apiController) SetUpRoute() {

	auth := "/auth"
	http.HandleFunc(http.MethodPost+" "+auth+"/login", controller.auth.Login)
}

func NewAPIController(ctx context.Context, mux *http.ServeMux, configs configs.Config) *apiController {
	auth := handlers.NewAuthHandler(configs.Inbound.Auth, configs.Key.ApiGateway)
	return &apiController{
		ctx:  ctx,
		mux:  mux,
		auth: auth,
	}
}
