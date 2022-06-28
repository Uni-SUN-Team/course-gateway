package ports

import "github.com/gin-gonic/gin"

type ConsumerController interface {
	Courses(c *gin.Context)
	CourseById(c *gin.Context)
	CourseBySlug(c *gin.Context)
}

type HealCheckController interface {
	HealthCheckHandler(c *gin.Context)
}
