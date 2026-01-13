package sbi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 自定义网络功能中处理数据请求的API接口
//
//	func (s *Server) getDefaultRoute() []Route {
//		return []Route{
//			{
//				Name:    "Hello, DRoT on Free5GC!",
//				Method:  http.MethodGet,
//				Pattern: "/",
//				APIFunc: func(c *gin.Context) {
//					c.JSON(http.StatusOK, "Hello, DRoT on free5GC!")
//				},
//				// Use
//				// curl -X GET http://127.0.0.163:8000/default/ -w "\n"
//			},
//		}
//	}
func (s *Server) GetRandomNumber() []Route {
	return []Route{
		{
			Name:    "Oracle_GetRandomNumber",
			Method:  http.MethodGet,
			Pattern: "/getRandomNumber",
			APIFunc: func(c *gin.Context) {

				c.JSON(http.StatusOK, "6617")
			},
			// Use
			// curl -X GET http://127.0.0.163:8000/default/ -w "\n"
		},
	}
}
