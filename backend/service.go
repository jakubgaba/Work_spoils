// TODO: DB filename
//NOTE: Something easy
// TODO: Main /GET/ function for DB
//NOTE: Before this step make sure you have created a DB :)
//TODO: Adding an option of order in a DB
//NOTE: Playing around with descending and ascending blah blah
//TODO: Write reach order line as a row
//TODO: New /GET/ function but this time also /POST/ function

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
