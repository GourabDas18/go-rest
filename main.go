package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	config "github.com/GourabDas18/g-rest/internal"
	"github.com/go-playground/validator/v10"

	"github.com/GourabDas18/g-rest/internal/controller"
	"github.com/GourabDas18/g-rest/internal/middleware"
	"github.com/GourabDas18/g-rest/internal/service"
)

var ValidatorG validator.Validate
var config0 *config.Config

func main() {

	config0 = config.MustLoad(&ValidatorG)

	router := http.NewServeMux()

	router.HandleFunc("GET /", controller.Test)

	router.HandleFunc("POST /auth/create", controller.CreateUser)
	router.HandleFunc("POST /auth/login", controller.LoginUser)

	router.HandleFunc("POST /countries/bulk", controller.CreateBulkCountry)
	router.HandleFunc("POST /country", controller.CountrySave)
	router.HandleFunc("GET /countries", controller.GetCountryList)

	server := http.Server{
		Addr:    config0.HTTPServer.Addr,
		Handler: middleware.Logger(router),
	}

	go func() {
		fmt.Printf("%s Server is running=====> %s\n", time.Now().Format("dd/MM/yyyy HH:mm a"), config0.HTTPServer.Addr)
		service.DBInit()
		service.MigrateModels(service.Db)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server %s", err.Error())
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	fmt.Println("Shutting down....")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Error on shutting down server, %s", err.Error())
	}

	fmt.Println("Server closed.")

}
