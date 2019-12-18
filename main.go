package main

import (
	"flag"
	"log"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	port := flag.String("port", "16003", "port for this server")
	debug := flag.Bool("debug", false, "turn on debug mode")
	var dllFolder string
	flag.StringVar(&dllFolder, "dll", "", "location to invoke dll")
	
	flag.Parse()

	if dllFolder == "" {
		os.Exit(1)
	}



	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H {
			"message": "pong",
		})
	})
}



