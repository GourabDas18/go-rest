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

	"github.com/GourabDas18/g-rest/internal/controller"
	"github.com/GourabDas18/g-rest/internal/service"
)

func main() {

	config := config.MustLoad()

	router := http.NewServeMux()

	router.HandleFunc("GET /", controller.Test)

	server := http.Server{
		Addr:    config.HTTPServer.Addr,
		Handler: router,
	}

	go func() {
		fmt.Printf("%s Server is running=====> %s\n", time.Now().Format("dd/MM/yyyy HH:mm a"), config.HTTPServer.Addr)
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
