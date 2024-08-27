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
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3" //Using ?
)

const dbFileName = "./db/data.db"

// BACKED DB/REST API STRUCT
type Service struct {
}

func (s *Service) getDatabase() (*sql.DB, error) {

	db, err := sql.Open("sqlite3", dbFileName)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

func (s *Service) testService(c *gin.Context) {
	c.JSON(http.StatusOK, "nil cuz it is what it is")
}
