package main

import (
	"context"
	"temp/app/control"
	"temp/internal/middleware"

	"github.com/dobyte/due/component/http/v2"
	"github.com/dobyte/due/v2"
	"github.com/dobyte/due/v2/eventbus"
	"github.com/dobyte/due/v2/log"

	eventbusredis "github.com/dobyte/due/eventbus/redis/v2"
)

// @title API文档
// @version 1.0
// @host localhost:8080
// @BasePath /
func main() {
	// 创建容器
	container := due.NewContainer()
	// 创建HTTP组件
	component := http.NewServer()
	// 初始化应用
	initApp(component.Proxy())
	// 添加网格组件
	container.Add(component)
	// 启动容器
	container.Serve()
}

// 初始化应用
func initApp(proxy *http.Proxy) {
	// 路由器
	router := proxy.Router()

	// 注册路由
	router.Post("/login", control.LoginHandler)
	router.Post("/register", control.RegisterHandler)

	router.Post("/createcate", control.CreateCateHandler)
	router.Get("/readcate", control.ReadCateHandler, middleware.AuthHandler)
	router.Post("/updatecate", control.UpdateCateHandler)
	router.Post("/deletecate", control.DeleteCateHandler)

	router.Post("/createarticle", control.CreateArticleHandler)
	router.Get("/readarticle", control.ReadArticleHandler)
	router.Post("/updatearticle", control.UpdateArticleHandler)
	router.Post("/deletearticle", control.DeleteArticleHandler)

	router.Post("/pub", control.PubHandler)
	router.Post("/rpc", control.RpcHandler)

	// 订阅发布
	eb := eventbusredis.NewEventbus()
	eb.Subscribe(context.Background(), "speak", eventHandler)
}

// 处理事件
func eventHandler(event *eventbus.Event) {
	log.Debug(event.ID)
	log.Debug(event.Topic)
	log.Debug(event.Payload)
	log.Debug(event.Timestamp)
}
