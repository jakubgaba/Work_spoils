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

// Handler to retrieve data from the database
func (s *Service) getTestData(c *gin.Context) {
	db, err := s.getDatabase()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, location FROM test")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query database"})
		return
	}
	defer rows.Close()

	// Prepare a slice to hold the data
	var results []map[string]interface{}

	// Loop through the rows
	for rows.Next() {
		var id int
		var name, location string
		if err := rows.Scan(&id, &name, &location); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan row"})
			return
		}

		// Append to the result slice
		results = append(results, map[string]interface{}{
			"id":       id,
			"name":     name,
			"location": location,
		})
	}

	// Send the data as a JSON response
	c.JSON(http.StatusOK, results)
}

func (s *Service) testService(c *gin.Context) {
	c.JSON(http.StatusOK, "nil cuz it is what it is")
}
