//[x] Test fetch

package backend

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3" //Using ?
)

// Testing
type Testing struct {
	service *Service
}

// Handler to retrieve data from the database
func (s *Testing) getTestData(c *gin.Context) {
	db, err := s.service.getDatabase()
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
