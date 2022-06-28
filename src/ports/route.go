package ports

import "github.com/gin-gonic/gin"

type ConsumerRoute interface {
	Consumer(r *gin.RouterGroup)
}
