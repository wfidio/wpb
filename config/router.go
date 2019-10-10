package router

import(
	"github.com/gin-gonic/gin"
)


type Router struct {
	routeList map[string]string
	r *gin.Engine
}

func (r *Router)setRouter()  {
	
}