package main

import (
	"log"
	"unisun/api/course-listener/src"
	"unisun/api/course-listener/src/config"

	"github.com/spf13/viper"
)

func main() {
	envService := config.New("app", "./src/resource")
	if err := envService.ConfigENV(); err != nil {
		log.Panic(err)
	}
	r := src.App()
	port := viper.GetString("app.port")
	if port == "" {
		r.Run(":8080")
	} else {
		r.Run(":" + port)
	}
}
