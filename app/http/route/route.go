package route

import (
	"github.com/gin-gonic/gin"
	"goinspect.cn/app/http/controller/paper"
	"goinspect.cn/app/http/controller/todolist"
	"goinspect.cn/app/http/controller/user"
	"goinspect.cn/app/http/controller/wx"
	"goinspect.cn/app/http/middleware/authenticate"
)

func SetupRouter(engine *gin.Engine) {

	engine.NoRoute(func(c *gin.Context) {

	})

	engine.GET("/ping", func(c *gin.Context) {

	})

	paperRouter := engine.Group("/paper")
	{
		paperRouter.GET("/list", paper.List)
		paperRouter.GET("/detail", paper.Detail)
		paperRouter.GET("/questions", paper.Questions)
		//paperRouter.POST("/:id/submit", assess.Submit)
	}

	authRouter := engine.Group("/auth")
	{
		authRouter.POST("/login", user.Login)
	}

	myRouter := engine.Group("/my", authenticate.SetUp())
	{
		myRouter.GET("/profile", user.MyProfile)
	}

	// todo list
	tdlRouter := engine.Group("/todolist", authenticate.SetUp())
	{
		tdlRouter.GET("/", todolist.Index)
		tdlRouter.POST("/new", todolist.New)
		tdlRouter.POST("/setdone", todolist.SetDone)
		tdlRouter.POST("/remove", todolist.Remove)
	}

	wxRouter := engine.Group("/wx")
	{
		wxRouter.GET("/verify", wx.ServiceVerify)
		wxRouter.GET("/login", wx.Login)
		wxRouter.GET("/logincb", wx.LoginCallback)
	}

}