package routes

import (
	"time"

	"github.com/gin-contrib/cors"
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

// serverを起動する
func (r *Router) Serve() {
	r.Engine.Run("localhost:8080")
}

// Routerのインスタンスにミドルウェアをセットする
func (r *Router) SetMiddleware() {
	r.Engine.Use(gin.Logger())
	r.Engine.Use(gin.Recovery())
}

// TODO: corsライブラリを理解する
func (r *Router) SetProxy() {
	r.Engine.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders: []string{"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
			"Origin",
			"X-CSRF-Token",
			"accept",
			"origin",
			"Cache-Control",
			"X-Requested-With",
		},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))
}
