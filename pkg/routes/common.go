package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// Read the page no. query parameter
		page, _ := strconv.Atoi(c.Query("page"))
		if page == 0 {
			page = 1
		}

		// Read the page_size query parameter
		pageSize, _ := strconv.Atoi(c.Query("page_size"))
		switch {
		// Max size 100
		case pageSize > 100:
			pageSize = 100
			// If -ve value, set to 10 (default)
		case pageSize <= 0:
			pageSize = 10
		}

		// calculate the offset
		offset := (page - 1) * pageSize
		// Return the database object with Offset and Limit
		return db.Offset(offset).Limit(pageSize)
	}
}
