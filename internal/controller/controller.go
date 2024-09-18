package controller

import (
	"context"
	"go-service/app"
	"net/http"
)

type apiController struct {
	ctx context.Context
	mux *http.ServeMux
	app *app.Application
}

func (controller *apiController) SetUpRoute() {
}

func NewAPIController(ctx context.Context, app *app.Application, mux *http.ServeMux) *apiController {
	return &apiController{
		ctx: ctx,
		app: app,
		mux: mux,
	}
}
