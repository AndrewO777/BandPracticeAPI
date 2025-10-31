package main

import (
	"os"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/AndrewO777/BandPracticeAPI/routes"
	"github.com/AndrewO777/BandPracticeAPI/db"
)

func main() {
	if err := db.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize database: %v\n", err)
		os.Exit(1)
	}
	defer db.DB.Close()

	router := gin.Default()
	router.SetTrustedProxies(nil)
	routes.RegisterSongRoutes(router)

	if err := router.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
	}
}
