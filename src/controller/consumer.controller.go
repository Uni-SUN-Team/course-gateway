package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"unisun/api/course-listenner/src/constants"
	"unisun/api/course-listenner/src/model"
	"unisun/api/course-listenner/src/service"

	"github.com/gin-gonic/gin"
)

func CourseAll(c *gin.Context) {
	var payloadRequest = model.ServiceIncomeRequest{}
	payloadRequest.Path = os.Getenv(constants.PATH_STRAPI_COURSE) + "?populate[thumnail]=*&populate[advisors][populate][thumnail]=*"
	payloadRequest.Method = constants.GET
	payloadRequest.Body = nil
	if query := c.Request.URL.RawQuery; query != "" {
		payloadRequest.Path += "&" + strings.Trim(query, "?")
	}
	var courses = model.Courses{}
	data := service.GetInformationFormStrapi(payloadRequest)
	err := json.Unmarshal([]byte(data.Payload), &courses)
	if err != nil {
		log.Panic("Change byte to json article", err.Error())
	} else {
		err = nil
	}
	for index, item := range courses.Data {
		var data_feedback = service.CallSumRateScoreFeedback(strconv.FormatInt(item.Id, 10))
		courses.Data[index].Rate.Star = data_feedback.Score
		courses.Data[index].Rate.Total = data_feedback.Count
	}
	c.JSON(http.StatusOK, gin.H{"error": err, "result": courses.Data, "pagination": courses.Meta.Pagination})
}

func CourseById(c *gin.Context) {
	id := c.Param("id")
	var course model.Course
	var payloadRequest = model.ServiceIncomeRequest{}
	payloadRequest.Path = os.Getenv(constants.PATH_STRAPI_COURSE) + "/" + id + "?populate[course_content][populate]=*&populate[thumnail]=*&populate[course_high_light]=*&populate[advisors][populate][thumnail]=*&populate[categories]=*"
	payloadRequest.Method = constants.GET
	payloadRequest.Body = nil

	data := service.GetInformationFormStrapi(payloadRequest)
	err := json.Unmarshal([]byte(data.Payload), &course)
	if err != nil {
		log.Panic("Change byte to json article", err.Error())
	} else {
		err = nil
	}
	var data_feedback = service.CallSumRateScoreFeedback(strconv.FormatInt(course.Data.Id, 10))
	course.Data.Rate.Star = data_feedback.Score
	course.Data.Rate.Total = data_feedback.Count
	c.JSON(http.StatusOK, gin.H{"error": err, "result": course.Data})
}

func CourseBySlug(c *gin.Context) {
	slug := c.Param("slug")
	var courses model.CourseSlug
	var payloadRequest = model.ServiceIncomeRequest{}
	payloadRequest.Path = os.Getenv(constants.PATH_STRAPI_COURSE) + "?populate[course_content][populate]=*&populate[thumnail]=*&populate[course_high_light]=*&populate[advisors][populate][thumnail]=*&populate[categories]=*&filters[slug][$eq]=" + slug
	payloadRequest.Method = constants.GET
	payloadRequest.Body = nil

	data := service.GetInformationFormStrapi(payloadRequest)
	err := json.Unmarshal([]byte(data.Payload), &courses)
	if err != nil {
		log.Println("Change byte to json article", err.Error())
	} else {
		err = nil
	}
	var course = courses.Data[0]
	var data_feedback = service.CallSumRateScoreFeedback(strconv.FormatInt(course.Id, 10))
	course.Rate.Star = data_feedback.Score
	course.Rate.Total = data_feedback.Count
	c.JSON(http.StatusOK, gin.H{"error": err, "result": course})
}
