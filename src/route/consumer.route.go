package route

import (
	"unisun/api/course-listener/src/ports"

	"github.com/gin-gonic/gin"
)

type RouteConsumerAdapter struct {
	Controller ports.ConsumerController
}

func NewRouteConsumerAdapter(controller ports.ConsumerController) *RouteConsumerAdapter {
	return &RouteConsumerAdapter{
		Controller: controller,
	}
}

func (srv *RouteConsumerAdapter) Consumer(r *gin.RouterGroup) {
	r.GET("/courses", srv.Controller.Courses)
	r.GET("/courses/:id", srv.Controller.CourseById)
	r.GET("/courses/slug/:slug", srv.Controller.CourseBySlug)
}
