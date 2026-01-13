package sbi

import (
	"net/http"

	"github.com/andy89923/nf-example/internal/logger"
	"github.com/gin-gonic/gin"
)

func (s *Server) HttpChallenge() []Route {
	return []Route{
		{
			Name:    "DRoT HttpChallenge Service!",
			Method:  http.MethodGet,
			Pattern: "/",
			APIFunc: func(c *gin.Context) {
				c.JSON(http.StatusOK, "Welcome to use DRoT HttpChallenge Service!")
			},
			// Use
			// curl -X GET http://127.0.0.163:8000//HttpChallenge/ -w "\n"
		},
		{
			Name:    "HttpChallenge Service!",
			Method:  http.MethodGet,
			Pattern: "/domain",
			APIFunc: s.OracleHTTPChallenge,
			// Use
			// curl -X GET http://127.0.0.163:8000/spyfamily/Anya -w "\n"
			// "Character: Anya Forger"
		},
	}
}

// api_router的职责只是接收外部的请求，然后解析和提取请求中的信息，把这些信息交给Processor去处理
func (s *Server) OracleHTTPChallenge(c *gin.Context) {
	logger.SBILog.Infof("Oracle Http challenge ")
	domain := c.Query("url")
	//domain := c.Param("url")
	if domain == "" {
		c.String(http.StatusBadRequest, "Domain is empty, please provide the domain u want to challenge.")
		return
	}

	s.Processor().ChallengeHttp(c, domain) //实际处理API的函数，写在processor中
}
