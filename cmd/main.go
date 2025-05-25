package main

import (
	"canchitas-libres-transaction/internal/configuration"
	database2 "canchitas-libres-transaction/internal/database"
	"canchitas-libres-transaction/internal/pgk/domain"
	"canchitas-libres-transaction/internal/pgk/infrastructure/repository/storage"
	"canchitas-libres-transaction/internal/pgk/infrastructure/web"
	"context"
	"fmt"
)

func main() {
	config, err := configuration.Load("../.env") //Inicializamos el archivo config y le cargamos la data que esta en el .env
	if err != nil {
		panic(err)
	}
	fmt.Println("Conexión a la base de datos exitosa")
	//database connection
	database, err := database2.NewDBConnection(context.Background(), config)
	if err != nil {
		panic(err)
	}

	// storage layer
	postgresStorage := storage.NewPostgresStorage(config, database)

	// application layer (services layer)
	service := domain.NewService(config, postgresStorage)

	// infrastructure layer
	handler := web.NewHandler(service)
	serv, err := web.NewServer(config, handler)
	if err != nil {
		fmt.Printf("error starting the server: %s\n", err)
		panic(err)
	}
	// Start application
	serv.Start()
}
