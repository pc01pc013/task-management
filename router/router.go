package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pc01pc013/task-management/apis"
	"github.com/pc01pc013/task-management/database"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	dbInstance := database.GetDBInstance()

	testApi := apis.NewTestApi(dbInstance)
	labelsApi := apis.NewLabelsApi(dbInstance)
	tasksApi := apis.NewTasksApi(dbInstance)

	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/db", testApi.DbTest)

	router.GET("/repeat", testApi.Repeat)

	apis := router.Group("/apis")
	{
		labelsRoute := apis.Group("/labels")
		{
			labelsRoute.GET("/list", labelsApi.List)
			labelsRoute.POST("/create", labelsApi.Create)
			labelsRoute.POST("/update", labelsApi.Update)
			labelsRoute.DELETE("/deleteL", labelsApi.Delete)
		}
		tasksRoute := apis.Group("/tasks")
		{
			tasksRoute.GET("/list", tasksApi.List)
			tasksRoute.POST("/create", tasksApi.Create)
			tasksRoute.POST("/update", tasksApi.Update)
		}
	}

	return router
}
