package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/golang/glog"
	"net/http"
	"time"
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
	gin.Logger()
	server.engine.Use(myLogger, gin.Recovery())
	server.initRouter()
	return server
}

func (s *Server) Run(addr ...string) error {
	return s.engine.Run(addr...)
}

func (s *Server) initRouter() {
	group := s.engine.Group("/workcalendar")
	group.GET("/", s.home)
	group.GET("/getholidaylist/:key", s.getHolidayList)
	group.GET("/getreplacestringmap/:key", s.getReplaceStringMap)
	group.GET("/getreplacecolormap/:key", s.getReplaceColorMap)
	group.GET("/gettip/:key", s.getTip)
	s.engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})
}

func (s *Server) getTip(c *gin.Context) {
	key := c.Param("key")

	tip, err := s.model.GetTip(key)
	if err != nil {
		log.Errorf("get holiday list error:%v", err)
		tip = ""
	}

	c.JSON(http.StatusOK, tip)
}

func (s *Server) getHolidayList(c *gin.Context) {
	key := c.Param("key")
	holidayList, err := s.model.GetHolidayList(key)
	if err != nil {
		log.Errorf("get holiday list error:%v", err)
		holidayList = nil
	}

	c.JSON(http.StatusOK, holidayList)
}

func (s *Server) home(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(http.StatusOK, `<a href="ftp://138.128.197.157/pub/WorkCalendar.apk" download="" > 下载班表 </a>`)
}

func (s *Server) getReplaceStringMap(c *gin.Context) {
	key := c.Param("key")
	replaceStringMap, err := s.model.GetReplaceStringMap(key)
	if err != nil {
		log.Errorf("get replace string map list error:%v", err)
		replaceStringMap = nil
	}

	c.JSON(http.StatusOK, replaceStringMap)
}

func (s *Server) getReplaceColorMap(c *gin.Context) {
	key := c.Param("key")
	replaceColorMap, err := s.model.GetReplaceColorMap(key)
	if err != nil {
		log.Errorf("get holiday list error:%v", err)
		replaceColorMap = nil
	}

	c.JSON(http.StatusOK, replaceColorMap)
}

func myLogger(c *gin.Context) {
	// Start timer
	start := time.Now()
	path := c.Request.URL.Path
	raw := c.Request.URL.RawQuery
	mothod := c.Request.Method

	// Process request
	c.Next()

	// Stop timer
	end := time.Now()
	latency := end.Sub(start)
	statusCode := c.Writer.Status()
	ecode := c.GetInt("context/err/code")
	clientIP := c.ClientIP()
	bodySize := c.Writer.Size()
	if raw != "" {
		path = path + "?" + raw
	}
	log.Infof("METHOD:%s | PATH:%s | CODE:%d | IP:%s | TIME:%d | ECODE:%d | BODY_SIZE:%d", mothod, path, statusCode, clientIP, latency/time.Millisecond, ecode, bodySize)
	//fmt.Printf("METHOD:%s | PATH:%s | CODE:%d | IP:%s | TIME:%d | ECODE:%d | BODY_SIZE:%d\n", mothod, path, statusCode, clientIP, latency/time.Millisecond, ecode, bodySize)
}
