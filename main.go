package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:  []string{"Origin"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	// Ping handler
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// Health handler
	router.GET("/healthz", func(c *gin.Context) {
		c.String(200, "healthy")
	})


	// Simple group: v1
	api := router.Group("/crypto")
	{
		encrypt := api.Group("/encrypt")
		{
			encrypt.GET("/", func(c *gin.Context) {
				textToEncrypt := c.DefaultQuery("to_encrypt", "")
				encryptedText := ""
				if textToEncrypt != "" {
					fmt.Printf("Text to be encrypted: %s\n", textToEncrypt)
					encryptedText, _ = Encrypt(textToEncrypt)
					fmt.Printf("Returned encrypted text: %s\n", textToEncrypt)
				}
				c.String(http.StatusOK, encryptedText)
			})
			encrypt.POST("/", func(c *gin.Context) {
				textToEncrypt := c.DefaultPostForm("to_encrypt", "")
				encryptedText := ""
				if textToEncrypt != "" {
					fmt.Printf("Text to be encrypted: %s\n", textToEncrypt)
					encryptedText, _ = Encrypt(textToEncrypt)
					fmt.Printf("Returned encrypted text: %s\n", textToEncrypt)
				}
				c.String(http.StatusOK, encryptedText)
			})
		}
		decrypt := api.Group("/decrypt")
		{
			decrypt.GET("/", func(c *gin.Context) {
				textToDecrypt := c.DefaultQuery("to_decrypt", "")
				decryptedText := ""
				if textToDecrypt != "" {
					fmt.Printf("Text to be decrypted: %s\n", textToDecrypt)
					decryptedText, _ = Decrypt(textToDecrypt)
					fmt.Printf("Returned decrypted text: %s\n", decryptedText)
				}
				c.String(http.StatusOK, decryptedText)
			})
			decrypt.POST("/", func(c *gin.Context) {
				textToDecrypt := c.DefaultPostForm("to_decrypt", "")
				decryptedText := ""
				if textToDecrypt != "" {
					fmt.Printf("Text to be decrypted: %s\n", textToDecrypt)
					decryptedText, _ = Decrypt(textToDecrypt)
					fmt.Printf("Returned decrypted text: %s\n", decryptedText)
				}
				c.String(http.StatusOK, decryptedText)
			})
		}
	}
	return router
}

func startServer() {
	// set config

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	// get and start router
	router := NewRouter()
	router.Run(":" + port)
}

func main() {
	startServer()
}
