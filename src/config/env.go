package config

import (
	"os"
	"unisun/api/course-listenner/src/constants"
)

func ConfigENV() {
	/**
	* Highlights: Endpoint
	 */
	os.Setenv(constants.PORT, "8080")
	os.Setenv(constants.CONTEXT_PATH, "/course-listenner")
	/**
	* Highlights: Strapi information gateway
	 */
	os.Setenv(constants.HOST_STRAPI_SERVICE, "https://api.unisun.dynu.com")
	os.Setenv(constants.PATH_STRAPI_INFORMATION_GATEWAY, "/strapi-information-gateway/api/strapi")
	os.Setenv(constants.PATH_STRAPI_COURSE, "/api/courses")
	/**
	* Highlights: Feedback Processor API
	 */
	os.Setenv(constants.HOST_FEEDBACK_PROCESSOR_API, "http://localhost:8081")
	/**
	* @Path : /feedback-processor-api/api/feedbacks/:id
	 */
	os.Setenv(constants.PATH_FEEDBACK_PROCESSOR_API, "/feedback-processor-api/api/feedbacks/")
}
