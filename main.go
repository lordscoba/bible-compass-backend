package main

import (
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/lordscoba/bible_compass_backend/internal/config"
	"github.com/lordscoba/bible_compass_backend/pkg/repository/storage/mongodb"
	"github.com/lordscoba/bible_compass_backend/pkg/router"
	"github.com/lordscoba/bible_compass_backend/utility"
)

func init() {
	config.Setup()
	// redis.SetupRedis() uncomment when you need redis
	mongodb.ConnectToDB()

	// s3.ConnectAws()

}
func main() {
	//Load config
	logger := utility.NewLogger()
	//  getConfig := config.GetConfig()
	validatorRef := validator.New()
	r := router.Setup(validatorRef, logger)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	//  logger.Info("Server is starting at 127.0.0.1:%s", getConfig.Server.Port)
	//  log.Fatal(r.Run(":" + getConfig.Server.Port))

	logger.Info("Server is starting at 127.0.0.1:%s", port)
	log.Fatal(r.Run(":" + port))
	// fmt.Println(logger)
}
