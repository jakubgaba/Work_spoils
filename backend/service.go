package backend

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// BACKED DB/REST API STRUCT
type Service struct {
}

func (s *Service) testService(c *gin.Context) {
	c.JSON(http.StatusOK, "nil cuz it is what it is")
}
