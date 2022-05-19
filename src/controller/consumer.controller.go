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

// Course godoc
// @summary      Course Listener
// @description  Course Listener for the service
// @id           CourseListenerHandler
// @tags         course
// @accept       json
// @produce      json
// @success      200    {object} 	model.ResponseCourseAll    "OK"
// @failure      400    {object} 	model.ResponseCourseAll    "Bad Request"
// @response     500    {object} 	model.ResponseCourseAll    "Internal Server Error"
// @router       /course-listenner/api/courses [get]
func CourseAll(c *gin.Context) {
	var payloadRequest = model.ServiceIncomeRequest{}
	var response = model.ResponseCourseAll{}
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
		response.Error = err
		c.JSON(http.StatusBadRequest, response)
		return
	} else {
		err = nil
	}
	for index, item := range courses.Data {
		var data_feedback = service.CallSumRateScoreFeedback(strconv.FormatInt(item.Id, 10))
		courses.Data[index].Rate.Star = data_feedback.Score
		courses.Data[index].Rate.Total = data_feedback.Count
	}
	response.Error = err
	response.Result = courses.Data
	response.Pagination = courses.Meta.Pagination
	c.JSON(http.StatusOK, response)
}

// Course godoc
// @summary      Course Listener
// @description  Course Listener for the service
// @id           CourseListenerIdHandler
// @tags         course
// @accept       json
// @produce      json
// @success      200    {object}	model.ResponseCourse    	"OK"
// @failure      400    {object}	model.ResponseCourse   	"Bad Request"
// @response     500    {object}	model.ResponseCourse    	"Internal Server Error"
// @router       /course-listenner/api/courses/:id [get]
func CourseById(c *gin.Context) {
	id := c.Param("id")
	var course model.Course
	var payloadRequest = model.ServiceIncomeRequest{}
	var response = model.ResponseCourse{}
	payloadRequest.Path = os.Getenv(constants.PATH_STRAPI_COURSE) + "/" + id + "?populate[course_content][populate]=*&populate[thumnail]=*&populate[course_high_light]=*&populate[advisors][populate][thumnail]=*&populate[categories]=*"
	payloadRequest.Method = constants.GET
	payloadRequest.Body = nil

	data := service.GetInformationFormStrapi(payloadRequest)
	err := json.Unmarshal([]byte(data.Payload), &course)
	if err != nil {
		log.Panic("Change byte to json article", err.Error())
		response.Error = err
		c.JSON(http.StatusBadRequest, response)
	} else {
		err = nil
	}
	var data_feedback = service.CallSumRateScoreFeedback(strconv.FormatInt(course.Data.Id, 10))
	course.Data.Rate.Star = data_feedback.Score
	course.Data.Rate.Total = data_feedback.Count
	response.Error = err
	response.Result = course.Data
	c.JSON(http.StatusOK, response)
}

// Course godoc
// @summary      Course Listener
// @description  Course Listener for the service
// @id           CourseListenerSlugHandler
// @tags         course
// @accept       json
// @produce      json
// @success      200    {object}    model.ResponseCourse    "OK"
// @failure      400    {object}    model.ResponseCourse    "Bad Request"
// @response     500    {object}    model.ResponseCourse    "Internal Server Error"
// @router       /course-listenner/api/courses/slug/:slug [get]
func CourseBySlug(c *gin.Context) {
	slug := c.Param("slug")
	var courses model.CourseSlug
	var payloadRequest = model.ServiceIncomeRequest{}
	var response = model.ResponseCourse{}
	payloadRequest.Path = os.Getenv(constants.PATH_STRAPI_COURSE) + "?populate[course_content][populate]=*&populate[thumnail]=*&populate[course_high_light]=*&populate[advisors][populate][thumnail]=*&populate[categories]=*&filters[slug][$eq]=" + slug
	payloadRequest.Method = constants.GET
	payloadRequest.Body = nil

	data := service.GetInformationFormStrapi(payloadRequest)
	err := json.Unmarshal([]byte(data.Payload), &courses)
	if err != nil {
		log.Println("Change byte to json article", err.Error())
		response.Error = err
		c.JSON(http.StatusBadRequest, response)
	} else {
		err = nil
	}
	var course = courses.Data[0]
	var data_feedback = service.CallSumRateScoreFeedback(strconv.FormatInt(course.Id, 10))
	course.Rate.Star = data_feedback.Score
	course.Rate.Total = data_feedback.Count
	response.Error = err
	response.Result = course
	c.JSON(http.StatusOK, response)
}
