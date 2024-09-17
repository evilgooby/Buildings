package params

import "github.com/gin-gonic/gin"

// парсинг параметров запроса
func ParseParams(c *gin.Context) map[string]interface{} {
	params := map[string]any{}

	city := c.Query("city")
	if city != "" {
		params["city"] = city
	}
	year := c.Query("year")
	if year != "" {
		params["year"] = year
	}
	floors := c.Query("floors")
	if floors != "" {
		params["floors"] = floors
	}

	return params
}
