// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"notification-service/internal/config"
	"notification-service/internal/infrastructure/api"
	"notification-service/internal/infrastructure/database"
)

// Injectors from wire.go:

// InitializeServer configura todas las dependencias y retorna una instancia de Server lista para usar
func InitializeServer(cfg config.Config) (*api.Server, error) {
	serverConfig := cfg.Server
	databaseConfig := cfg.Database
	db, err := database.NewGormDB(databaseConfig)
	if err != nil {
		return nil, err
	}
	server := api.NewServer(serverConfig, db)
	return server, nil
}
