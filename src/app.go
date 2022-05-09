package src

import (
	"github.com/gin-gonic/gin"
)

func App() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
	// g := r.Group(os.Getenv("CONTEXT_PATH") + "/api")
	// {

	// }
	return r
}
