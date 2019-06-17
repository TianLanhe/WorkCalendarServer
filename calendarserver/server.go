package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	model  *Model
	engine *gin.Engine
}

func NewServer(model *Model) *Server {
	server := &Server{
		model:  model,
		engine: gin.New(),
	}
	server.initRouter()
	return server
}

func (s *Server) Run(addr ...string) error {
	return s.engine.Run(addr...)
}

func (s *Server) initRouter() {
	group := s.engine.Group("/workcalendar")
	group.POST("/getholidaylist", s.getHolidayList)
	group.POST("/getreplacestringmap", s.getReplaceStringMap)
	group.POST("/getreplacecolormap", s.getReplaceColorMap)
	s.engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})
}

func (s *Server) getHolidayList(c *gin.Context) {
}

func (s *Server) getReplaceStringMap(c *gin.Context) {
}

func (s *Server) getReplaceColorMap(c *gin.Context) {
}
