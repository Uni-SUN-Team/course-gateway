package route

import (
	"unisun/api/course-listenner/src/controller"

	"github.com/gin-gonic/gin"
)

func Consumer(r *gin.RouterGroup) {
	r.GET("/courses", controller.CourseAll)
	r.GET("/courses/:id", controller.CourseById)
	r.GET("/courses/slug/:slug", controller.CourseBySlug)
}
