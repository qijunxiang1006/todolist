package main

import (
	`fmt`
	
	`github.com/kataras/iris`
	`github.com/kataras/iris/core/router`
	`github.com/kataras/iris/middleware/logger`
	recover2 `github.com/kataras/iris/middleware/recover`
)

func main(){
	app:=initApp()
	fmt.Print(app)
	// mvc.New(app).Handle(nil)
}
// func registerView(application *iris.Application){
// 	application.RegisterView()
// }

func initApp()  *iris.Application{
	app:=iris.New()
	app.Use(logger.New())
	app.Use(recover2.New())
	// cor:=cors.New(cors.Options{
	// 	AllowedOrigins: []string{"*"},
	// 	AllowedMethods:         []string{"GET","POST","PUT","DELETE"},
	// 	MaxAge:                 0,
	// })
	// app.Use(func(context iris.Context) {
	// 	cor.HandlerFunc(context.ResponseWriter(),context.Request())
	// })
	return app
}

func creatMVC(app *iris.Application){
	r:=router.NewAPIBuilder()
	r.Handle("GET","/",)
	// mvc.New(r).Handle()
}