package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"unisun/api/course-listener/src/constants"
	"unisun/api/course-listener/src/model"
	"unisun/api/course-listener/src/ports"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type ControllerConsumerAdapter struct {
	Service   ports.ConsumerService
	Component ports.MappingAdvisorComponentPort
}

func NewControllerConsumerAdapter(service ports.ConsumerService, component ports.MappingAdvisorComponentPort) *ControllerConsumerAdapter {
	return &ControllerConsumerAdapter{
		Service:   service,
		Component: component,
	}
}

// Course godoc
// @summary      Course Listener
// @description  Course Listener for the service
// @id           CourseListenerHandler
// @tags         course
// @accept       json
// @produce      json
// @success      200    {object} 	model.ResponseCourses	"OK"
// @failure      400    {object} 	model.ResponseFail	"Bad Request"
// @response     500    {object} 	model.ResponseFail	"Internal Server Error"
// @router       /course-listenner/api/courses [get]
func (srv *ControllerConsumerAdapter) Courses(c *gin.Context) {
	var payloadRequest = model.ServiceIncomeRequest{}
	var response = model.ResponseCourses{}
	var responseFail = model.ResponseFail{}
	payloadRequest.Path = strings.Join([]string{viper.GetString("endpoint.strapi-information-gateway.mapping.course.path"), viper.GetString("endpoint.strapi-information-gateway.mapping.course.query.value")}, "")
	payloadRequest.Method = constants.GET
	payloadRequest.Body = nil
	if query := c.Request.URL.RawQuery; query != "" {
		payloadRequest.Path += "&" + strings.Trim(query, "?")
	}
	data, err := srv.Service.GetInformationFormStrapi(payloadRequest)
	if err != nil {
		log.Panic(err)
	}
	err = json.Unmarshal([]byte(data), &response)
	if err != nil {
		responseFail.Error.Message = err.Error()
		responseFail.Error.Status = http.StatusBadRequest
		responseFail.Error.Name = "Unmarshal parses the JSON-encoded"
		c.JSON(http.StatusBadRequest, responseFail)
		log.Panic("Unmarshal parses the JSON-encoded", err)
		return
	}
	for i, c := range response.Data {
		resultAdvisor, err := srv.Component.MappingAdvisor(c.Advisors)
		if err != nil {
			log.Panic(err)
		}
		response.Data[i].Advisors = resultAdvisor
	}
	c.AbortWithStatusJSON(http.StatusOK, response)
}

// Course godoc
// @summary      Course Listener
// @description  Course Listener for the service
// @id           CourseListenerIdHandler
// @tags         course
// @accept       json
// @produce      json
// @success      200    {object}	model.ResponseCourse    	"OK"
// @failure      400    {object}	model.ResponseFail   	"Bad Request"
// @response     500    {object}	model.ResponseFail    	"Internal Server Error"
// @router       /course-listenner/api/courses/:id [get]
func (srv *ControllerConsumerAdapter) CourseById(c *gin.Context) {
	// id := c.Param("id")
	// var course model.Course
	// var payloadRequest = model.ServiceIncomeRequest{}
	// var response = model.ResponseCourse{}
	// var responseFail = model.ResponseFail{}
	// payloadRequest.Path = os.Getenv(constants.PATH_STRAPI_COURSE) + "/" + id + "?populate[Course][populate]=*&populate[thumbnail]=*&populate[course_high_light]=*&populate[advisors][populate][thumbnail]=*&populate[categories]=*"
	// payloadRequest.Method = constants.GET
	// payloadRequest.Body = nil
	// data := service.GetInformationFormStrapi(payloadRequest)
	// err := json.Unmarshal([]byte(data.Payload), &course)
	// if err != nil {
	// 	log.Panic("Change byte to json article", err.Error())
	// 	responseFail.Error.Message = err.Error()
	// 	responseFail.Error.Status = http.StatusBadRequest
	// 	responseFail.Error.Name = "Unmarshal parses the JSON-encoded"
	// 	c.JSON(http.StatusBadRequest, responseFail)
	// 	return
	// } else {
	// 	err = nil
	// }
	// var data_feedback = service.CallSumRateScoreFeedback(strconv.FormatInt(course.Data.Id, 10))
	// course.Data.Rate.Star = data_feedback.Score
	// course.Data.Rate.Total = data_feedback.Count
	// response.Result = course.Data
	// c.JSON(http.StatusOK, response)
}

// Course godoc
// @summary      Course Listener
// @description  Course Listener for the service
// @id           CourseListenerSlugHandler
// @tags         course
// @accept       json
// @produce      json
// @success      200    {object}    model.ResponseCourse    "OK"
// @failure      400    {object}    model.ResponseFail    "Bad Request"
// @response     500    {object}    model.ResponseFail    "Internal Server Error"
// @router       /course-listenner/api/courses/slug/:slug [get]
func (srv *ControllerConsumerAdapter) CourseBySlug(c *gin.Context) {
	// slug := c.Param("slug")
	// var courses model.CourseSlug
	// var payloadRequest = model.ServiceIncomeRequest{}
	// var response = model.ResponseCourse{}
	// var responseFail = model.ResponseFail{}
	// payloadRequest.Path = os.Getenv(constants.PATH_STRAPI_COURSE) + "?populate[Course][populate]=*&populate[thumbnail]=*&populate[course_high_light]=*&populate[advisors][populate][thumbnail]=*&populate[categories]=*&filters[slug][$eq]=" + slug
	// payloadRequest.Method = constants.GET
	// payloadRequest.Body = nil

	// data := service.GetInformationFormStrapi(payloadRequest)
	// err := json.Unmarshal([]byte(data.Payload), &courses)
	// if err != nil {
	// 	log.Println("Change byte to json article", err.Error())
	// 	responseFail.Error.Message = err.Error()
	// 	responseFail.Error.Status = http.StatusBadRequest
	// 	responseFail.Error.Name = "Unmarshal parses the JSON-encoded"
	// 	c.JSON(http.StatusBadRequest, responseFail)
	// 	return
	// } else {
	// 	err = nil
	// }
	// var course = courses.Data[0]
	// // var data_feedback = service.CallSumRateScoreFeedback(strconv.FormatInt(course.Id, 10))
	// // course.Rate.Star = data_feedback.Score
	// // course.Rate.Total = data_feedback.Count
	// response.Result = course
	// c.JSON(http.StatusOK, response)
}
