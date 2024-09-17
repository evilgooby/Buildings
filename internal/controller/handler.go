package controller

import (
	"Building/internal/middleware/errorHandling"
	"Building/internal/params"
	"Building/internal/repository/postgres"
	"Building/internal/struct"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateBuilding godoc
// @Summary      Add in database
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Success      200  {integer}  integer 1
// @Failure      400  {integer}  400
// @Failure      418  {integer}  418
// @Failure      500  {integer}  500
// @Router       /CreateBuilding [post]
func CreateBuilding(c *gin.Context) {
	var building entities.Building
	if err := c.ShouldBindJSON(&building); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := postgres.AddBuilding(c, building); err != nil {
		errorHandling.ErrorHandler(c)
	}
	c.JSON(200, gin.H{"message": "Building created successfully"})
}

// GetBuildings godoc
// @Summary      Get from database
// @Description  get accounts
// @Tags         accounts
// @Produce      json
// @Param        input body entities.Building  true  "Buildings"
// @Success      200  {integer}  integer 1
// @Failure      400  {integer}  400
// @Failure      418  {integer}  418
// @Failure      500  {integer}  500
// @Router       /GetBuildings [get]
func GetBuildings(c *gin.Context) {
	param := params.ParseParams(c)
	buildings, err := postgres.GetBuildings(c, param)
	if err != nil {
		errorHandling.ErrorHandler(c)
	}
	c.JSON(200, buildings)
}
