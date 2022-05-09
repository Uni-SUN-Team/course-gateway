package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"unisun/api/course-listenner/src/constants"
	"unisun/api/course-listenner/src/model"
	"unisun/api/course-listenner/src/service"

	"github.com/gin-gonic/gin"
)

func CourseAll(c *gin.Context) {
	var payloadRequest = model.ServiceIncomeRequest{}
	payloadRequest.Path = os.Getenv(constants.PATH_STRAPI_ARTICLE)
	payloadRequest.Method = constants.GET
	payloadRequest.Body = nil
	if query := c.Request.URL.RawQuery; query != "" {
		payloadRequest.Path += "?" + query
	}
	var articles model.Courses
	data := service.GetInformationFormStrapi(payloadRequest)
	err := json.Unmarshal([]byte(data.Payload), &articles)
	if err != nil {
		log.Println("Change byte to json article", err.Error())
	} else {
		err = nil
	}
	c.JSON(http.StatusOK, gin.H{"error": err, "data": articles})
}

func CourseById(c *gin.Context) {
	id := c.Param("id")
	populate := c.DefaultQuery("populate", "*")
	var articles model.Course
	var payloadRequest = model.ServiceIncomeRequest{}
	payloadRequest.Path = os.Getenv(constants.PATH_STRAPI_ARTICLE) + "/" + id + "?populate=" + populate
	payloadRequest.Method = constants.GET
	payloadRequest.Body = nil

	data := service.GetInformationFormStrapi(payloadRequest)
	err := json.Unmarshal([]byte(data.Payload), &articles)
	if err != nil {
		log.Println("Change byte to json article", err.Error())
	} else {
		err = nil
	}
	c.JSON(http.StatusOK, gin.H{"error": err, "data": articles})
}

func CourseBySlug(c *gin.Context) {
	slug := c.Param("slug")
	populate := c.DefaultQuery("populate", "*")
	var articles model.Course
	var payloadRequest = model.ServiceIncomeRequest{}
	payloadRequest.Path = os.Getenv(constants.PATH_STRAPI_ARTICLE) + "/?filters[slug][$eq]=" + slug + "&populate=" + populate
	payloadRequest.Method = constants.GET
	payloadRequest.Body = nil

	data := service.GetInformationFormStrapi(payloadRequest)
	err := json.Unmarshal([]byte(data.Payload), &articles)
	if err != nil {
		log.Println("Change byte to json article", err.Error())
	} else {
		err = nil
	}
	c.JSON(http.StatusOK, gin.H{"error": err, "data": articles})
}
