package postgres

import (
	"Building/config"
	"Building/internal/middleware/errorHandling"
	"Building/internal/struct"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"strings"
)

// Константы для добавления/получения данных в/из базы данных
const (
	addDB        = `INSERT INTO buildings (name, city, year, floors)VALUES ($1, $2, $3, $4)`
	getBuildings = `SELECT * FROM buildings `
)

// Добавление в базу данных
func AddBuilding(c *gin.Context, buildings entities.Building) error {
	_, err := config.StorePostgres.PostgresDB.Exec(addDB, buildings.Name, buildings.City, buildings.Year, buildings.Floors)
	if err != nil {
		return c.Error(errorHandling.ErrDB)
	}
	return nil
}

// Получение из базы данных с параметрами и без
func GetBuildings(c *gin.Context, option map[string]any) ([]entities.Building, error) {
	var buildings []entities.Building
	var values []any
	var where []string
	var i = 1
	for k, v := range option {
		values = append(values, v)
		where = append(where, fmt.Sprintf("%s = $%d", k, i))
		i++
	}
	var query = getBuildings
	if len(where) != 0 {
		query += "WHERE "
	}
	query = query + strings.Join(where, " AND ")
	r, err := config.StorePostgres.PostgresDB.Query(query, values...)
	if err != nil {
		return nil, c.Error(errorHandling.ErrDB)
	}
	for r.Next() {
		var id int
		var name string
		var city string
		var year int
		var floors int
		err = r.Scan(&id, &name, &city, &year, &floors)
		if err != nil {
			return nil, c.Error(errorHandling.ErrDB)
		}
		building := entities.Building{
			ID:     id,
			Name:   name,
			City:   city,
			Year:   year,
			Floors: floors,
		}
		buildings = append(buildings, building)
	}
	return buildings, nil
}
