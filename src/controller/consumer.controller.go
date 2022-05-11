package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"unisun/api/course-listenner/src/constants"
	"unisun/api/course-listenner/src/model"
	"unisun/api/course-listenner/src/service"

	"github.com/gin-gonic/gin"
)

func CourseAll(c *gin.Context) {
	var payloadRequest = model.ServiceIncomeRequest{}
	payloadRequest.Path = os.Getenv(constants.PATH_STRAPI_ARTICLE) + "?populate=*"
	payloadRequest.Method = constants.GET
	payloadRequest.Body = nil
	if query := c.Request.URL.RawQuery; query != "" {
		payloadRequest.Path += "&" + strings.Trim(query, "?")
	}
	log.Println(payloadRequest.Path)
	var courses model.Courses
	data := service.GetInformationFormStrapi(payloadRequest)
	log.Println(data)
	err := json.Unmarshal([]byte(data.Payload), &courses)
	if err != nil {
		log.Println("Change byte to json article", err.Error())
	} else {
		err = nil
	}
	c.JSON(http.StatusOK, gin.H{"error": err, "data": courses})
}

func CourseById(c *gin.Context) {
	id := c.Param("id")
	var course model.Course
	var payloadRequest = model.ServiceIncomeRequest{}
	payloadRequest.Path = os.Getenv(constants.PATH_STRAPI_ARTICLE) + "/" + id + "?populate[course_content][populate]=*&populate[thumnail][populate]=*&populate[course_high_light][populate]=*&populate[advisors][populate]=*&populate[categories][populate]=*&populate[localizations][populate]=*"
	payloadRequest.Method = constants.GET
	payloadRequest.Body = nil

	data := service.GetInformationFormStrapi(payloadRequest)
	err := json.Unmarshal([]byte(data.Payload), &course)
	if err != nil {
		log.Println("Change byte to json article", err.Error())
	} else {
		err = nil
	}
	c.JSON(http.StatusOK, gin.H{"error": err, "data": course})
}

func CourseBySlug(c *gin.Context) {
	slug := c.Param("slug")
	populate := c.DefaultQuery("populate", "*")
	var courses model.Courses
	var payloadRequest = model.ServiceIncomeRequest{}
	payloadRequest.Path = os.Getenv(constants.PATH_STRAPI_ARTICLE) + "?filters[slug][$eq]=" + slug + "&populate=" + populate
	payloadRequest.Method = constants.GET
	payloadRequest.Body = nil

	data := service.GetInformationFormStrapi(payloadRequest)
	err := json.Unmarshal([]byte(data.Payload), &courses)
	if err != nil {
		log.Println("Change byte to json article", err.Error())
	} else {
		err = nil
	}
	c.JSON(http.StatusOK, gin.H{"error": err, "data": courses})
}
