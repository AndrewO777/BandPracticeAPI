package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/AndrewO777/BandPracticeAPI/routes"
	"github.com/AndrewO777/BandPracticeAPI/db"
)

func main() {
	db.Init()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		fmt.Println("\nShutting down...")
		db.DB.Close()
		os.Exit(0)
	}()


	router := gin.Default()
	router.SetTrustedProxies(nil)
	routes.RegisterSongRoutes(router)

	router.Run()
}
