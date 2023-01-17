package routes

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine
}

// Routerを初期化し、そのインスタンスのポインタを返す
func NewRouter() *Router {
	e := gin.New()
	return &Router{
		Engine: e,
	}
}

// Routerのインスタンスにミドルウェアをセットする
func (r *Router) SetMiddleware() {
	r.Engine.Use(gin.Logger())
	r.Engine.Use(gin.Recovery())
}
