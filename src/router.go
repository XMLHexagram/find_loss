package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Service) RouterInit() {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/allLost", s.getAllLost)
	temp := r.Group("opera")
	{
		temp.POST("/add",s.add)
		temp.DELETE("/:id",s.delete)
		temp.PUT("/:id",s.modify)
		temp.GET("/:id", s.getLost)
	}

	s.Router = r
	err :=s.Router.Run(s.Config.Web.Port)
	DealError(err)
}
