package main

import (
	uuid "github.com/iris-contrib/go.uuid"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/mvc"
	
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	recover2 "github.com/kataras/iris/v12/middleware/recover"
	
	`todolist/controller`
	`todolist/model`
)


func main() {
	app := initWebService()
	initRoute(app)
	_ = app.Run(iris.Addr(":8080"))
}

func initWebService() *iris.Application {
	app := iris.New()
	app.Use(logger.New())
	app.Use(recover2.New())
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowCredentials: true,
	})
	app.Use(crs)
	initWebSource(app)
	return app
}

func initRoute(app *iris.Application) {
	mvc := mvc.Configure(app.Party("/list"))
	mvc.Register(controller.NewService())
	mvc.Handle(new(controller.Controller))
}


func initWebSource(app *iris.Application){
	app.RegisterView(iris.HTML("./webapp/", ".html"))
	app.HandleDir("/static", "./webapp/static")
	// app.RegisterView(iris.HTML("./webapp/build", ".html"))
	// app.HandleDir("/static", "./webapp/build/static")
	app.Get("/", func(c *context.Context) {
		_ = c.View("index.html")
	})
	app.Get("/uuid", func(c *context.Context) {
		uid, err := uuid.NewV1()
		if err != nil {
			_, _ = c.JSON(model.NewFailResponse(err))
		} else {
			_, _ = c.JSON(model.NewSuccessResponse(uid))
		}
	})
}