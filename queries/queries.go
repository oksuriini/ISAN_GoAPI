package queries

import (
	"ISAN-api/isan"
	"database/sql"
	"errors"
	"fmt"

	// mysql driver imported
	_ "github.com/go-sql-driver/mysql"
)

// Database struct is meant to hold sql.DB struct, address of the database Addr, and port number Port.
type Database struct {
	db   *sql.DB
	Addr string
	Port string
}

// Function creates Database struct based on given arguments.
// Function has two parameters 'addr' and 'port', where addr designates the address of sql server
// and port designates the port to be used.
func CreateDBstruct(addr string, port string) (*Database, error) {
	newdb, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", "root", "passw0rd", addr, "api"))
	if err != nil {
		return nil, err
	}

	newD := Database{
		db:   newdb,
		Addr: addr,
		Port: port,
	}

	return &newD, nil
}

// Function's purpose is to validate parameter 'field'.
// Used to validate SQL query
func ValidateField(field string) bool {
	switch field {
	case "record_type", "title", "work_type", "release_date", "director", "duration_min", "original_language", "isan":
		return true
	default:
		return false
	}
}

// Function accepts parameters 'field' and 'value', where 'field' is the field in SQL
// database, and 'value' is the value of that field.
// eg. GetRowsByFilter("duration_min", "120") will return all rows, where duration_min value is 120.
func (d *Database) GetRowsByFilter(field, value string) ([]isan.ISAN, error) {
	var rowsISAN []isan.ISAN

	if !ValidateField(field) {
		err := errors.New(fmt.Sprintf("Field validation failed. Field %s does not exists.\n", field))
		return nil, err
	}

	return rowsISAN, nil
}

// Get all fetches all of the rows, and returns an array of isan.ISAN struct.
// This function is here for development/testing/development purposes.
func (d *Database) GetAll() ([]isan.ISAN, error) {
	var allISAN []isan.ISAN

	rows, err := d.db.Query("SELECT isan, record_type, title, work_type, release_date, director, duration_min, original_language FROM isan")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var isan_num string
		var record_type string
		var title string
		var work_type string
		var release_date string
		var director string
		var duration_min string
		var original_language string

		if err := rows.Scan(&isan_num, &record_type, &title, &work_type, &release_date, &director, &duration_min, &original_language); err != nil {
			return nil, err
		}

		fmt.Println(isan_num)

		newisan := isan.CreateNewISAN(isan_num, record_type, title, work_type, release_date, director, duration_min, original_language)

		allISAN = append(allISAN, newisan)
	}

	return allISAN, nil
}

// Function to insert ISAN record into database.
// PLACEHOLDER TODO LATER
func (d *Database) InsertISAN(isan string) (int, error) {
	return 1, nil
}
