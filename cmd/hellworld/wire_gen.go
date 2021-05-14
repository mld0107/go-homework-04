// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"hellworld/internal/biz"
	"hellworld/internal/conf"
	"hellworld/internal/data"
	"hellworld/internal/server"
	"hellworld/internal/service"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, error) {
	dataData, err := data.NewData(confData)
	if err != nil {
		return nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	userUseCase := biz.NewUserUseCase(userRepo, logger)
	userService := service.NewUserService(userUseCase, logger)
	httpServer := server.NewHTTPServer(confServer, userService)
	grpcServer := server.NewGRPCServer(confServer, userService)
	app := newApp(logger, httpServer, grpcServer)
	return app, nil
}
