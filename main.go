package main

import (
	"os"
	"unisun/api/course-listenner/src"
	"unisun/api/course-listenner/src/config"
	"unisun/api/course-listenner/src/constants"
)

func main() {
	if os.Getenv(constants.NODE) != constants.PRODUCTION {
		config.ConfigENV()
	}
	r := src.App()
	port := os.Getenv(constants.PORT)
	if port == "" {
		r.Run(":8080")
	} else {
		r.Run(":" + port)
	}
}
