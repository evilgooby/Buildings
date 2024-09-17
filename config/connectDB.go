package config

import (
	"Building/internal/middleware/errorHandling"
	"database/sql"
	"github.com/spf13/viper"
	"log"
)

const (
	createTable = `CREATE TABLE buildings (id SERIAL PRIMARY KEY, name TEXT NOT NULL, city TEXT NOT NULL, year INT NOT NULL, floors INT NOT NULL);`
	verifyTable = `SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = $1)`
)

type StoragePostgres struct {
	PostgresDB *sql.DB
}

var (
	StorePostgres = &StoragePostgres{}
)

// Подключение к базе данных
func InitializeStorePostgres() *StoragePostgres {
	connStr := viper.GetString("db.url")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(errorHandling.ErrConnectDb)
	}
	exists, err := TableExists(db, "buildings")
	if err != nil {
		log.Fatal(errorHandling.ErrConnectDb)
	}
	if !exists {
		_, err = db.Exec(createTable)
		if err != nil {
			log.Fatal(errorHandling.ErrConnectDb)
		}
	}
	StorePostgres.PostgresDB = db
	return StorePostgres
}

// Проверка существует ли таблица
func TableExists(db *sql.DB, tableName string) (bool, error) {
	stmt, err := db.Prepare(verifyTable)
	if err != nil {
		return false, nil
	}
	var exists bool
	if err = stmt.QueryRow(tableName).Scan(&exists); err != nil {
		return false, nil
	}
	return exists, nil
}
