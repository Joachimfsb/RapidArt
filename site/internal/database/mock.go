package database

import (
	"database/sql/driver"
	"log"
	"reflect"

	"github.com/DATA-DOG/go-sqlmock"
)

/////////////// MOCK //////////////////

// Create mock and setup fake db.
//
// Warning, only one mock can be in use at a time.
// Creation of multiple mock will result in only the last one working.
func CreateMock() sqlmock.Sqlmock {

	var err error
	var mock sqlmock.Sqlmock

	db, mock, err = sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	log.Println("Mock db initialized")

	return mock
}

func DeleteMock() {
	db.Close()

	log.Println("Mock db closed")
}

// Generate rows from data
func GenRows(data any) *sqlmock.Rows {

	slice := reflect.ValueOf(data)         // Slice
	sliceElemType := slice.Index(0).Type() // Types of element that slice contains

	// Check that data is slice
	if slice.Kind() != reflect.Slice {
		log.Println("Could not generate rows: data not slice")
		return nil
	}
	// Check that element type is struct
	if sliceElemType.Kind() != reflect.Struct {
		log.Println("Could not generate rows: slice elements not struct")
		return nil
	}
	// Check if length is 0
	if slice.Len() == 0 {
		return sqlmock.NewRows(nil)
	}

	// Get columns of struct fields
	var columns []string
	for i := 0; i < sliceElemType.NumField(); i++ {
		columns = append(columns, sliceElemType.Field(i).Name)
	}

	// Define columns
	rows := sqlmock.NewRows(columns)

	// Loop through each input row
	for i := 0; i < slice.Len(); i++ {
		// Loop through each column value
		var values []driver.Value
		for j := 0; j < sliceElemType.NumField(); j++ {
			values = append(values, slice.Index(i).Field(j).Interface())
		}
		rows.AddRow(values...) // Add single row
	}

	return rows
}
