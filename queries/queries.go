package queries

import (
	"ISAN-api/isan"
	"database/sql"
	"errors"
	"fmt"
	"strings"

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

	// Create sql.DB connection and later pipeline it to isan.Database struct
	newdb, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", "root", "passw0rd", addr, "api"))

	// Debugging purposes
	fmt.Println(fmt.Sprintf("%s:%s@tcp(%s)/%s", "root", "passw0rd", addr, "api"))

	if err != nil {
		return nil, err
	}

	// Create the isan.Database struct to be used in API calls
	newD := Database{
		db:   newdb,
		Addr: addr,
		Port: port,
	}

	return &newD, nil
}

// Function's purpose is to validate parameter 'field'.
// Used to validate SQL query
func (d *Database) ValidateField(field string) bool {
	switch field {
	case "record_type", "title", "work_type", "release_date", "director", "duration_min", "original_language", "isan_num":
		return true
	default:
		return false
	}
}

// Function takes in argument of type string, which is used to filter out results based on given argument.
// Function returns a string which is used to make SQL query to the database.
// Currently function does additional validation by accepting only specific strings as the value of field, which is a bit redundant, so it will later be removed.
func SelectFieldString(field string) string {
	switch field {
	case "record_type":
		return fmt.Sprintf("SELECT isan_num, record_type, title, work_type, release_date, director, duration_min, original_language FROM isan WHERE %s=?", field)
	case "title":
		return fmt.Sprintf("SELECT isan_num, record_type, title, work_type, release_date, director, duration_min, original_language FROM isan WHERE %s=?", field)
	case "work_type":
		return fmt.Sprintf("SELECT isan_num, record_type, title, work_type, release_date, director, duration_min, original_language FROM isan WHERE %s=?", field)
	case "release_date":
		return fmt.Sprintf("SELECT isan_num, record_type, title, work_type, release_date, director, duration_min, original_language FROM isan WHERE %s=?", field)
	case "director":
		return fmt.Sprintf("SELECT isan_num, record_type, title, work_type, release_date, director, duration_min, original_language FROM isan WHERE %s=?", field)
	case "duration_min":
		return fmt.Sprintf("SELECT isan_num, record_type, title, work_type, release_date, director, duration_min, original_language FROM isan WHERE %s=?", field)
	case "original_language":
		return fmt.Sprintf("SELECT isan_num, record_type, title, work_type, release_date, director, duration_min, original_language FROM isan WHERE %s=?", field)
	case "isan":
		return fmt.Sprintf("SELECT isan_num, record_type, title, work_type, release_date, director, duration_min, original_language FROM isan WHERE %s=?", field)
	default:
		return ""
	}

}

// Function accepts parameters 'field' and 'value', where 'field' is the field in SQL
// database, and 'value' is the value of that field.
// eg. GetRowsByFilter("duration_min", "120") will return all rows, where duration_min value is 120.
func (d *Database) GetRowsByFilter(field, value string) ([]isan.ISAN, error) {
	var rowsISAN []isan.ISAN

	if !d.ValidateField(field) {
		err := errors.New(fmt.Sprintf("Field validation failed. Field %s does not exists.\n", field))
		return nil, err
	}

	fieldString := SelectFieldString(field)
	if strings.Compare(fieldString, "") == 0 {
		err := errors.New(fmt.Sprintf("Field %s not found on allowed field list.\n", field))
		return nil, err
	}

	rows, err := d.db.Query(fieldString, value)
	if err != nil {
		fmt.Println("Something unexpected happened in db query.")
		fmt.Println("Field: " + field + ", Value: " + value)
		return nil, err
	}

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

		rowsISAN = append(rowsISAN, newisan)
	}
	return rowsISAN, nil
}

// WORKING FUNCTION
// Get all fetches all of the rows, and returns an array of isan.ISAN struct.
// This function is here for development/testing/development purposes.
func (d *Database) GetAll() ([]isan.ISAN, error) {
	var allISAN []isan.ISAN

	rows, err := d.db.Query("SELECT isan_num, record_type, title, work_type, release_date, director, duration_min, original_language FROM isan")
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

// WORKING FUNCTION
// Function to insert ISAN record into database.
func (d *Database) InsertISAN(isan isan.ISAN) (int, error) {
	is, rt, ti, wt, rd, di, dm, ol := isan.GetIsanVals()

	query := "INSERT INTO `isan` (`isan_num`, `record_type`, `title`, `work_type`, `release_date`, `director`, `duration_min`, `original_language`) VALUES(?,?,?,?,?,?,?,?)"

	result, err := d.db.Exec(query, is, rt, ti, wt, rd, di, dm, ol)
	if err != nil {
		return 0, err
	}

	resvar, err := result.RowsAffected()
	if err != nil {
		return 0, nil
	}

	return int(resvar), nil
}
