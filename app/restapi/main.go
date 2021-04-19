package main

import (
	"zipcode/data/repositories"
	"zipcode/domain/usecase"
	"zipcode/infrastructure"

	"zipcode/presentation/api"
)

var apiController api.ApiController

func init() {
	repository := repositories.NewFileRepository()
	useCase := usecase.NewAddressFinder(repository)
	apiController = api.NewApiController(useCase)
}

func main() {
	apiServer := infrastructure.ApiServer(apiController)
	apiServer.Run()
}